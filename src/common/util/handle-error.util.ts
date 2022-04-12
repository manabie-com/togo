import { HttpException, HttpStatus } from '@nestjs/common';

export enum ErrorCode {
  OK = 'OK',
  BAD_REQUEST = 'BAD_REQUEST',
  NOT_FOUND = 'NOT_FOUND',
  PERMISSION_DENIED = 'PERMISSION_DENIED',
}

export const handleError = (
  message: string,
  errorCode: ErrorCode,
  httpStatusCode: HttpStatus,
): void => {
  throw new HttpException(
    {
      message: message,
      errorCode: errorCode,
      status: httpStatusCode,
    },
    httpStatusCode,
  );
};
