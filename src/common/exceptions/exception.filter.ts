import {
  ArgumentsHost,
  Catch,
  ExceptionFilter,
  ForbiddenException,
  HttpException,
  HttpStatus,
  UnauthorizedException,
} from '@nestjs/common';
import { HttpArgumentsHost } from '@nestjs/common/interfaces/features/arguments-host.interface';
import { Response } from 'express';
import { QueryFailedError } from 'typeorm';

import { LoggerService } from 'src/common/logger/logger.service';
import responseHelper from 'src/common/helpers/response.helper';

@Catch()
export class AllExceptionFilter implements ExceptionFilter {
  constructor(private logger: LoggerService) {}

  private static handleResponse(
    response: Response,
    exception: HttpException | QueryFailedError | Error,
  ): void {
    let statusCode = HttpStatus.OK;
    let message = exception.message;
    if (
      exception instanceof UnauthorizedException ||
      exception instanceof ForbiddenException
    ) {
      statusCode = exception.getStatus();
    }
    if (exception instanceof QueryFailedError) message = 'Đã có lỗi xảy ra';
    response.status(statusCode).json(responseHelper.failed(message));
  }

  catch(exception: HttpException | Error, host: ArgumentsHost): void {
    const ctx: HttpArgumentsHost = host.switchToHttp();
    const response: Response = ctx.getResponse();

    // Response to client
    AllExceptionFilter.handleResponse(response, exception);
    console.log(exception);

    // Handling error message and logging
    if (
      exception instanceof HttpException &&
      exception.getStatus() !== HttpStatus.OK &&
      exception.getStatus() !== HttpStatus.CREATED
    ) {
      this.handleMessage(exception);
    }
  }

  private handleMessage(
    exception: HttpException | QueryFailedError | Error,
  ): void {
    let message = 'Đã có lỗi xảy ra';

    if (exception instanceof HttpException) {
      message = JSON.stringify(exception.getResponse());
    } else if (exception instanceof QueryFailedError) {
      message = exception.stack.toString();
    } else if (exception instanceof Error) {
      message = exception.stack.toString();
    }

    this.logger.error(message);
  }
}
