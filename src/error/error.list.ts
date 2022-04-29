import { StatusCode } from '../common/enum';

enum ERROR_CODE {
  INVALID_REQUEST = 'INVALID_REQUEST',
  UNKNOWN_ERROR = 'UNKNOWN_ERROR',
  INCORRECT_FIELD = 'INCORRECT_FIELD',
  REQUIRED = 'REQUIRED',
  MAX_LENGTH = 'MAX_LENGTH'
}

// customized error message for joi
const JoiValidationErrors = {
  required: ERROR_CODE.REQUIRED,
  'string.max': ERROR_CODE.MAX_LENGTH
};

const ErrorList = {
  [ERROR_CODE.INVALID_REQUEST]: {
    statusCode: StatusCode.BAD_REQUEST,
    message: 'Invalid request'
  },
  [ERROR_CODE.UNKNOWN_ERROR]: {
    statusCode: StatusCode.NOT_FOUND,
    message: 'Unknown error'
  },
  [ERROR_CODE.INCORRECT_FIELD]: {
    statusCode: StatusCode.BAD_REQUEST,
    message: 'Incorrect data'
  },
  [ERROR_CODE.REQUIRED]: {
    statusCode: StatusCode.BAD_REQUEST,
    message: 'Required'
  },
  [ERROR_CODE.MAX_LENGTH]: {
    statusCode: StatusCode.BAD_REQUEST,
    message: 'Max Length'
  }
};

export { ERROR_CODE, ErrorList, JoiValidationErrors };
