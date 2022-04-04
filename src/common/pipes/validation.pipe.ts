import {
  PipeTransform,
  Injectable,
  ArgumentMetadata,
  BadRequestException,
} from '@nestjs/common';
import { validate } from 'class-validator';
import { plainToClass } from 'class-transformer';
import * as _ from 'lodash';

@Injectable()
export class ValidationPipe implements PipeTransform<any> {
  async transform(value: any, { metatype }: ArgumentMetadata) {
    if (!metatype || !this.toValidate(metatype)) {
      return value;
    }
    const object = plainToClass(metatype, value);
    console.log(object);
    const errors = await validate(object);
    if (errors.length > 0) {
      console.log('bad request error : ', errors);
      let errorMsg = 'Validation failed';

      if (_.isObject(errors[0].constraints)) {
        errorMsg = _.values(errors[0].constraints)[0];
      } else if (errors[0].children && _.isObject(errors[0].children[0])) {
      }
      throw new BadRequestException(errorMsg);
    }

    return object;
  }

  private toValidate(metatype: any): boolean {
    const types: any[] = [String, Boolean, Number, Array, Object];
    return !types.includes(metatype);
  }
}
