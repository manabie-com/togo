import { Transform } from 'class-transformer';
import { IsInt, Min } from 'class-validator';

export class PaginationQueryDto {
  @IsInt()
  @Min(1)
  @Transform(({ value }) => parseInt(value))
  page: number;

  @IsInt()
  @Min(1)
  @Transform(({ value }) => parseInt(value))
  size: number;
}
