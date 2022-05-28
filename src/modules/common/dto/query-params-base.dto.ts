import { ApiProperty, PartialType } from '@nestjs/swagger';
import { IsOptional, IsString } from 'class-validator';
import { PaginationParamsDto } from './pagination-params-dto';

export class QueryParamsBaseDto extends PartialType(PaginationParamsDto) {
  @ApiProperty({ required: false, description: 'Search key word' })
  @IsOptional()
  @IsString()
  readonly q?: string;

  offset?: number;

  rows?: number;

  sortKey?: string;

  keyword?: 'ASC' | 'DESC';

  constructor() {
    super();
    Object.defineProperty(this, 'offset', {
      get() {
        const temp = (this.page - 1) * this.perPage;

        return temp > 0 ? temp : 0;
      },
    });
    Object.defineProperty(this, 'rows', {
      get() {
        return this.perPage > 0 ? this.perPage : 0;
      },
    });
    Object.defineProperty(this, 'sortKey', {
      get() {
        const split = this.orderBy.split('|');

        return split[0];
      },
    });
    Object.defineProperty(this, 'keyword', {
      get() {
        const split = this.orderBy.split('|');

        return split[1] === 'ASC' ? 'ASC' : 'DESC';
      },
    });
  }
}
