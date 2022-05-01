import { IBaseModel } from '../common/type';
import { TaskStatusEnum } from './task.enum';

export interface ITaskReason {
  errorCode: string;
  message: string;
}

export interface ITask extends IBaseModel {
  name: string;
  userId: string;
  status: TaskStatusEnum;
  reason?: ITaskReason;
}

export interface ICreateTaskPayload {
  name: string;
  userId: string;
}

export interface IUpdateTaskByIdPayload {
  status?: TaskStatusEnum;
  reason?: ITaskReason;
  name?: string;
}
