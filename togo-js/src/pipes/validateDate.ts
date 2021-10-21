import {
  PipeTransform,
  Injectable,
  HttpException,
  HttpStatus,
} from '@nestjs/common';
import * as moment from 'moment';

export function checkDate(dateString: string) {
  return moment(dateString, 'YYYY-MM-DD', true).isValid();
}

export function checkDateTime(dateString: string) {
  return moment(dateString).isValid();
}

@Injectable()
export class ValidateDatePipe implements PipeTransform<any> {
  async transform(value: string) {
    const isValidDate = checkDate(value);
    if (!isValidDate)
      throw new HttpException(
        {
          statusCode: HttpStatus.BAD_REQUEST,
          message: 'INVALID_DATE',
        },
        HttpStatus.BAD_REQUEST,
      );
    return value;
  }
}
