import errorHandler, { buildMappedErrorDetails } from '../errorHandler.plugin';
import { ERROR_CODE } from '../../error/error.list';
import { AppError } from '../../error/error.service';
import hapi from '@hapi/hapi';

describe('errorHandler.plugin', () => {
  describe('errorHandler', () => {
    it('Should return payload if there is no error', () => {
      const expected = 'expected';
      const _req = {} as hapi.Request;
      const _res = { continue: expected } as any as hapi.ResponseToolkit;
      const result = errorHandler(_req, _res, null);
      expect(result).toEqual(expected);
    });

    it('Should throw error when the error is not AppError', () => {
      const expected = 'expected';
      const _req = {} as hapi.Request;
      const _res = { continue: expected } as any as hapi.ResponseToolkit;
      const error = new Error('_test');
      try {
        errorHandler(_req, _res, error);
      } catch (err) {
        expect(err).toEqual(error);
      }
    });

    it('Should throw error when the error is AppError', () => {
      const expected = 'expected';
      const _req = {} as hapi.Request;
      const _res = { continue: expected } as any as hapi.ResponseToolkit;
      const error = new AppError(ERROR_CODE.INVALID_REQUEST);
      try {
        errorHandler(_req, _res, error);
      } catch (err) {
        expect(err).toEqual(error);
      }
    });
  });

  describe('buildMappedErrorDetails', () => {
    const detail = {
      message: '_message',
      path: ['id'],
      type: 'string.max',
      context: { limit: 3, value: '_test', label: 'id', key: 'id' }
    };
    it('Should return the error mapping', () => {
      const details = [detail];
      const mappedDetails = [
        { code: ERROR_CODE.MAX_LENGTH, key: 'id', type: detail.type }
      ];

      const expected = buildMappedErrorDetails(details);
      expect(expected).toEqual(mappedDetails);
    });

    it('Should return the error mapping with the constraint not defined', () => {
      const details = [
        {
          ...detail,
          type: 'string.unknown'
        }
      ];
      const mappedDetails = [
        { code: ERROR_CODE.INCORRECT_FIELD, key: 'id', type: 'string.unknown' }
      ];

      const expected = buildMappedErrorDetails(details);
      expect(expected).toEqual(mappedDetails);
    });

    it('Should return the error mapping with no duplicate', () => {
      const details = [detail, detail];
      const mappedDetails = [
        { code: ERROR_CODE.MAX_LENGTH, key: 'id', type: detail.type }
      ];

      const expected = buildMappedErrorDetails(details);

      expect(expected.length).toEqual(1);
      expect(expected).toEqual(mappedDetails);
    });
  });
});
