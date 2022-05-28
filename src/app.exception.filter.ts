import { ExceptionFilter, Catch, ArgumentsHost, HttpException, HttpStatus, Logger } from '@nestjs/common';
import { Response } from 'express';
import { ValidationError } from 'class-validator';
import { InvalidTokenException, TokenExpiredException } from '@modules/common/exceptions';

@Catch()
export class AppExceptionsFilter implements ExceptionFilter {
  private logger = new Logger('AppExceptionsFilter', true);

  catch(exception: any, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const response = ctx.getResponse<Response>();
    const request = ctx.getRequest<Request>();
    const defaultErrorResponse = {
      statusCode: HttpStatus.INTERNAL_SERVER_ERROR,
      error: 'unknown_error',
      message: [exception.message || 'Unknown error occurred'],
    };

    this.logger.log(`${'='.repeat(20)}ERROR${'='.repeat(20)}`);
    this.logger.log(exception);
    this.logger.error(request.url);
    this.logger.error(exception.message);
    let resJson: any = { ...defaultErrorResponse };

    if (this.isAuthenticationError(exception)) {
      const status = exception.getStatus();
      const errorResponse = exception.getResponse() as {
        error: string;
        message: string;
      };

      resJson = { ...defaultErrorResponse, ...errorResponse };
      response
        .status(status)
        .header(
          'WWW-Authenticate',
          `Bearer realm="service" error="${errorResponse.error}", error_description="${errorResponse.message}"`,
        )
        .json(resJson);
    } else if (exception instanceof HttpException) {
      const status = exception.getStatus();
      const errorResponse = exception.getResponse();

      if (this.isValidationErrorResponse(errorResponse) && process.env.NODE_ENV == 'production') {
        resJson = { ...defaultErrorResponse, ...errorResponse, ...{ message: 'Validation error' } };
        response.status(status).json(resJson);
      } else if (errorResponse instanceof Object) {
        resJson = { ...defaultErrorResponse, ...errorResponse };
        response.status(status).json(resJson);
      } else {
        resJson = { ...defaultErrorResponse, ...{ message: errorResponse } };
        response.status(status).json(resJson);
      }
    } else {
      response.status(defaultErrorResponse.statusCode).json(resJson);
    }
  }

  private isAuthenticationError(exception: any): exception is HttpException {
    return exception instanceof InvalidTokenException || exception instanceof TokenExpiredException;
  }

  private isValidationErrorResponse(response: any): response is ValidationError {
    return (
      response.message instanceof Array && response.message.length > 0 && response.message[0] instanceof ValidationError
    );
  }
}
