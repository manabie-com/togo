import hapi from '@hapi/hapi';
import { ValidationError, ValidationErrorItem } from 'joi';

import { AppError } from '../error/error.service';
import { ERROR_CODE, JoiValidationErrors } from '../error/error.list';
import { IErrorDetails } from '../error/error.type';

interface JoiValidationErrors {
  [index: string]: ERROR_CODE;
}

const buildKey = (paths: (string | number)[]): string => {
  return paths.join('.');
};

const buildMappedErrorDetails = (
  details: ValidationErrorItem[]
): IErrorDetails[] => {
  return details.reduce<IErrorDetails[]>((acc, detail, index) => {
    if (
      index !== 0 &&
      buildKey(detail.path) === buildKey(details[index - 1].path)
    ) {
      return acc;
    }
    const constraint = detail.type;
    const errorCode =
      (JoiValidationErrors as JoiValidationErrors)[constraint] ||
      ERROR_CODE.INCORRECT_FIELD;
    acc.push({
      code: errorCode,
      key: buildKey(detail.path),
      type: detail.type
    });

    return acc;
  }, []);
};

const errorHandler: hapi.Lifecycle.Method = (
  _req: hapi.Request,
  res: hapi.ResponseToolkit,
  err?: Error
) => {
  if (!err) {
    return res.continue;
  }
  if ((err as ValidationError).isJoi) {
    const { details } = err as ValidationError;
    const mappedDetails = buildMappedErrorDetails(details);
    throw new AppError(ERROR_CODE.INVALID_REQUEST, mappedDetails);
  }
  throw err;
};

export { buildKey, buildMappedErrorDetails };
export default errorHandler;
