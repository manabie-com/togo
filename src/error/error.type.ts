import { StatusCode } from '../common/enum';

export interface IErrorDetails {
  key: string;
  code: string;
}
export interface IErrorList {
  [key: string]: {
    statusCode: StatusCode;
    message: string;
  };
}
