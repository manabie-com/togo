import hapi from '@hapi/hapi';
import Joi from 'joi';

import responseWrapper from '../responseWrapper.plugin';
import { AppError } from '../../error/error.service';
import { ERROR_CODE, ErrorList } from '../../error/error.list';

describe('Plugin - responseWrapper', () => {
  const testServer = new hapi.Server();
  const testHandler = jest.fn();
  const testHandler2 = jest.fn();
  testServer.route([
    {
      method: 'POST',
      path: '/test',
      options: {
        handler: testHandler,
        validate: {
          payload: Joi.object({
            name: Joi.string()
          })
        }
      }
    },
    {
      method: 'POST',
      path: '/test2',
      options: {
        handler: testHandler2,
        validate: {
          payload: Joi.object({
            name: Joi.string()
          })
        }
      }
    },
    {
      method: 'GET',
      path: '/documentation',
      options: {
        handler: () => 'documentation'
      }
    }
  ]);

  testServer.register(responseWrapper);

  beforeEach(() => {
    jest.resetAllMocks();
  });

  describe('handleHapiResponse', () => {
    it('should wrap success response in data object', async () => {
      testHandler.mockResolvedValueOnce('test result');

      const response: hapi.ServerInjectResponse = await testServer.inject({
        method: 'POST',
        url: '/test',
        payload: {
          name: 'test'
        }
      });

      expect(response.result).toEqual({
        data: 'test result'
      });
    });

    it('should wrap AppError response in error object', async () => {
      testHandler.mockRejectedValueOnce(
        new AppError(ERROR_CODE.INVALID_REQUEST)
      );

      const response: hapi.ServerInjectResponse = await testServer.inject({
        method: 'POST',
        url: '/test',
        payload: {
          name: 'test'
        }
      });

      expect(response.result).toEqual({
        errors: undefined,
        statusCode: ErrorList[ERROR_CODE.INVALID_REQUEST].statusCode,
        message: 'INVALID_REQUEST'
      });
      expect(response.statusCode).toEqual(
        ErrorList[ERROR_CODE.INVALID_REQUEST].statusCode
      );
    });

    it('should wrap 500 error response in error object', async () => {
      testHandler.mockRejectedValueOnce('unexpected error');

      const response: hapi.ServerInjectResponse = await testServer.inject({
        method: 'POST',
        url: '/test',
        payload: {
          name: 'test'
        }
      });
      expect(response.result).toEqual({
        statusCode: 500,
        payload: {
          statusCode: 500,
          error: 'Internal Server Error',
          message: 'An internal server error occurred'
        },
        headers: {}
      });
      expect(response.statusCode).toEqual(500);
    });

    it('should ignore document path', async () => {
      const response: hapi.ServerInjectResponse = await testServer.inject({
        method: 'GET',
        url: '/documentation'
      });

      expect(response.result).not.toHaveProperty('data');
      expect(response.result).not.toHaveProperty('error');
    });
  });
});
