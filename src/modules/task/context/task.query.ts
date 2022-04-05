import { TaskFilter } from './task.filter';
import { TaskOrderBy } from './task.orderBy';
import { TaskPagination } from './task.pagination';

export class TaskQuery {
  filter?: TaskFilter;
  orderBy?: TaskOrderBy;
  pagination?: TaskPagination;
}
