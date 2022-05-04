import { omit } from 'lodash';
import { documentToObject, isMongoObjectId } from '../common/util';
import { ERROR_CODE } from '../error/error.list';
import { AppError } from '../error/error.service';
import taskModel from './task.model';
import {
  ICreateTaskPayload,
  IGetTasksQuery,
  ITask,
  IUpdateTaskByIdPayload
} from './task.type';

const createTask = async (payload: ICreateTaskPayload): Promise<ITask> => {
  const task: ITask = await taskModel.create(payload);
  return task;
};

const getById = async (taskId: string): Promise<ITask | null> => {
  if (!isMongoObjectId(taskId)) {
    return null;
  }
  const task: ITask = await taskModel
    .findOne({
      _id: taskId
    })
    .exec();
  return task;
};

const updateById = async (
  taskId: string,
  updateObj: IUpdateTaskByIdPayload
): Promise<ITask> => {
  if (!isMongoObjectId(taskId)) {
    throw new AppError(ERROR_CODE.TASK_NOT_FOUND);
  }

  const task = await taskModel
    .findByIdAndUpdate(taskId, updateObj, {
      new: true
    })
    .exec();

  if (!task) {
    throw new AppError(ERROR_CODE.TASK_NOT_FOUND);
  }

  return task as ITask;
};

const getsByUserId = async (query: IGetTasksQuery): Promise<ITask[]> => {
  if (query.status === 'ALL') {
    query = omit(query, 'status');
  }
  const tasks = await taskModel.find(query).lean().exec();

  return tasks.map((task) => documentToObject(task) as ITask);
};

export default {
  createTask,
  getById,
  updateById,
  getsByUserId
};
