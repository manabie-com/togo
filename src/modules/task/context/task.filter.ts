import { PriorityEnum, StatusEnum } from './../../../common/index';

export class TaskFilter {
  id?: number;

  // code?: string;

  title?: string;

  assignee_id?: number;

  description?: string;

  priority?: PriorityEnum;

  status?: StatusEnum;

  updated_time?: Date;

  created_time?: Date;
}
