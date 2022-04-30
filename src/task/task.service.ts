import { ERROR_CODE } from '../error/error.list';
import { AppError } from '../error/error.service';
import userService from '../user/user.service';
import taskRepository from './task.repository';
import { ICreateTaskPayload, ITask } from './task.type';

const createTask = async (payload: ICreateTaskPayload): Promise<ITask> => {
  try {
    await userService.getById(payload.userId);
    return await taskRepository.createTask(payload);
  } catch (err) {
    throw err;
  }
};

const getById = async (taskId: string): Promise<ITask> => {
  const task = await taskRepository.getById(taskId);
  if (!task) {
    throw new AppError(ERROR_CODE.TASK_NOT_FOUND);
  }
  return task;
};
export default {
  createTask,
  getById
};
