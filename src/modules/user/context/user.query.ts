import { UserFilter } from './user.filter';
import { UserOrderBy } from './user.orderBy';
import { UserPagination } from './user.pagination';

export class UserQuery {
  filter?: UserFilter;
  orderBy?: UserOrderBy;
  pagination?: UserPagination;
}
