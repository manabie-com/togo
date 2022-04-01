import { ApiProperty } from '@nestjs/swagger';
import {
  IsDateString,
  IsNotEmpty,
  IsString,
  ValidateIf,
} from 'class-validator';
import { isNil } from 'lodash';

export class CreateTaskDto {
  @ApiProperty()
  @IsNotEmpty()
  @IsString()
  title: string;

  @ApiProperty()
  @IsNotEmpty()
  @IsString()
  desc: string;

  @ApiProperty()
  @ValidateIf((obj, value) => !isNil(value))
  @IsDateString()
  deadlineAt: string;
}
