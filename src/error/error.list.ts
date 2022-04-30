import { StatusCode } from '../common/enum';
import { IErrorList } from './error.type';

enum ERROR_CODE {
  INVALID_REQUEST = 'INVALID_REQUEST',
  UNKNOWN_ERROR = 'UNKNOWN_ERROR',
  INCORRECT_FIELD = 'INCORRECT_FIELD',
  REQUIRED = 'REQUIRED',
  MAX_LENGTH = 'MAX_LENGTH',
  DUPLICATE_USER = 'DUPLICATE_USER',
  CREATE_USER_ERROR = 'CREATE_USER_ERROR',
  USER_NOT_FOUND = 'USER_NOT_FOUND',
  TASK_NOT_FOUND = 'TASK_NOT_FOUND'
}

// customized error message for joi
const JoiValidationErrors = {
  required: ERROR_CODE.REQUIRED,
  'string.max': ERROR_CODE.MAX_LENGTH
};

const ErrorList: IErrorList = {
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
  },
  [ERROR_CODE.DUPLICATE_USER]: {
    statusCode: StatusCode.BAD_REQUEST,
    message: 'Duplicate username'
  },
  [ERROR_CODE.CREATE_USER_ERROR]: {
    statusCode: StatusCode.BAD_REQUEST,
    message: 'Create user error'
  },
  [ERROR_CODE.USER_NOT_FOUND]: {
    statusCode: StatusCode.BAD_REQUEST,
    message: 'User not found'
  },
  [ERROR_CODE.TASK_NOT_FOUND]: {
    statusCode: StatusCode.BAD_REQUEST,
    message: 'Task not found'
  }
};

export { ERROR_CODE, ErrorList, JoiValidationErrors };
