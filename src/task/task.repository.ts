import { isMongoObjectId } from '../common/util';
import taskModel from './task.model';
import { ICreateTaskPayload, ITask } from './task.type';

const createTask = async (payload: ICreateTaskPayload): Promise<ITask> => {
  const user: ITask = await taskModel.create(payload);
  return user;
};

const getById = async (userId: string): Promise<ITask | null> => {
  if (!isMongoObjectId(userId)) {
    return null;
  }
  const user: ITask = await taskModel
    .findOne({
      _id: userId
    })
    .exec();
  return user;
};

export default {
  createTask,
  getById
};
