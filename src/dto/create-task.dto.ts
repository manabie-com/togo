import {
  IsDateString,
  IsNotEmpty,
  IsNumber,
  IsString,
  Min,
  ValidateIf,
} from 'class-validator';
import { isNil } from 'lodash';

export class CreateTaskDto {
  @IsNotEmpty()
  @IsString()
  title: string;

  @IsNotEmpty()
  @IsString()
  desc: string;

  @ValidateIf((obj, value) => !isNil(value))
  @IsDateString()
  deadlineAt: string;
}
