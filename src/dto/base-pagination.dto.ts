import { ApiProperty } from '@nestjs/swagger';
import { Transform } from 'class-transformer';
import { ValidateIf, IsNumber, Min } from 'class-validator';
import { isNil } from 'lodash';

export class PaginationDto {
  @ApiProperty({ type: Number, required: false, default: 20, minimum: 1 })
  @ValidateIf((obj, value) => !isNil(value))
  @Transform(({ value }) => +value)
  @IsNumber()
  @Min(1)
  limit = 20;

  @ApiProperty({ type: Number, required: false, default: 0, minimum: 0 })
  @ValidateIf((obj, value) => !isNil(value))
  @Transform(({ value }) => +value)
  @IsNumber()
  @Min(0)
  page = 0;
}
