import { StatusCode } from '../common/enum';
import { ERROR_CODE, ErrorList } from './error.list';
import { IErrorDetails } from './error.type';

export class AppError extends Error {
  public errorCode: ERROR_CODE;
  errors?: IErrorDetails[];
  code: number = StatusCode.INTERNAL_SERVER_ERROR;
  constructor(errorCode: ERROR_CODE, errors?: IErrorDetails[]) {
    super(errorCode);
    this.errorCode = errorCode;
    this.name = AppError.name;
    this.errors = errors;
    this.code = ErrorList[errorCode].statusCode;
  }

  getErrors() {
    const error = ErrorList[this.errorCode];
    return {
      errors: this.errors,
      statusCode: error.statusCode,
      message: this.errorCode
    };
  }
}
