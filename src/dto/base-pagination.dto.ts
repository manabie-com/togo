import { Transform } from 'class-transformer';
import { ValidateIf, IsNumber, Min } from 'class-validator';
import { isNil } from 'lodash';

export class PaginationDto {
  @ValidateIf((obj, value) => !isNil(value))
  @Transform(({ value }) => +value)
  @IsNumber()
  @Min(1)
  limit = 20;

  @ValidateIf((obj, value) => !isNil(value))
  @Transform(({ value }) => +value)
  @IsNumber()
  @Min(0)
  page = 0;
}
