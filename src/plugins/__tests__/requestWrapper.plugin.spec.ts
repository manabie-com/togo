import hapi from '@hapi/hapi';
import Joi from 'joi';

import requestWrapper from '../requestWrapper.plugin';
import responseWrapper from '../responseWrapper.plugin';

describe('Plugin - requestWrapper', () => {
  const testServer = new hapi.Server();
  const testHandler = jest.fn();
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
    }
  ]);

  testServer.register([requestWrapper, responseWrapper]);

  beforeEach(() => {
    jest.resetAllMocks();
    process.env.SERVICE_NAME = '';
  });

  describe('handleHapiRequest', () => {
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
  });
});
