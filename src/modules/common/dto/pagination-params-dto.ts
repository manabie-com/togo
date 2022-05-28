import { ApiProperty } from '@nestjs/swagger';
import { IsOptional } from 'class-validator';
import { DEFAULT_PER_PAGE } from '../constant';

export class PaginationParamsDto {
  @ApiProperty({ required: false })
  @IsOptional()
  readonly page?: number = 1;

  @ApiProperty({ required: false })
  @IsOptional()
  readonly perPage?: number = DEFAULT_PER_PAGE;

  @ApiProperty({ required: false, description: '{sortKeyWord}|{direction}. Direction: ASC,DESC' })
  @IsOptional()
  readonly orderBy?: string;

  constructor(partial: Partial<PaginationParamsDto>) {
    Object.assign(this, partial);
  }
}
