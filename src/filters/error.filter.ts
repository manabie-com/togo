import {
  ArgumentsHost,
  Catch,
  ExceptionFilter,
  HttpException,
  HttpStatus
} from '@nestjs/common';
import { Request, Response } from 'express';

@Catch()
export class ErrorExceptionFilter implements ExceptionFilter {
  catch(exception: HttpException, host: ArgumentsHost): unknown {
    const ctx = host.switchToHttp();
    const response = ctx.getResponse<Response>();
    const request = ctx.getRequest<Request>();

    console.log(exception);

    if (exception instanceof HttpException) {
      const data: any = exception.getResponse();
      const err = {
        message: data?.message || data,
        errorCode: data?.errorCode || 1,
        statusCode: data?.statusCode || exception.getStatus(),
        timestamp: new Date().toISOString(),
        path: request.url
      };

      return response.status(err.statusCode).send(err);
    }

    const error: any = new Error(exception);

    response.status(HttpStatus.BAD_REQUEST).send({
      statusCode: HttpStatus.BAD_REQUEST,
      errorCode: 1,
      timestamp: new Date().toISOString(),
      path: request.url,
      message: error?.message,
      stack: error.stack
    });
  }
}
