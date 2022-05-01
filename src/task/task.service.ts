import { ErrorList, ERROR_CODE } from '../error/error.list';
import { AppError } from '../error/error.service';
import userService from '../user/user.service';
import taskRepository from './task.repository';
import { ICreateTaskPayload, ITask, IUpdateTaskByIdPayload } from './task.type';
import KafkaService from '../common/kafka';
import logger from '../logger';
import RedisService from '../common/redis';
import { TaskStatusEnum } from './task.enum';
import { TASK_CONSUMER_TOPIC } from './task.topic';

const createTask = async (payload: ICreateTaskPayload): Promise<ITask> => {
  try {
    await userService.getById(payload.userId);
    const task = await taskRepository.createTask(payload);
    await KafkaService.produceMessage(TASK_CONSUMER_TOPIC, [
      { value: task._id.toString() }
    ]);
    return task;
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

const processTask = async (taskId: string): Promise<void> => {
  logger.info(`processTask >>>> taskId: ${taskId}`);
  const task = await getById(taskId);
  let count = (await RedisService.get(task.userId)) as unknown as number;
  const { configuration } = await userService.getById(task.userId);
  const updateTaskObj: IUpdateTaskByIdPayload = { status: TaskStatusEnum.DONE };

  if (count >= configuration.limit) {
    logger.info(
      `processTask >>>> count(${count}) >= limit(${configuration.limit})`
    );
    updateTaskObj.status = TaskStatusEnum.FAILED;
    updateTaskObj.reason = {
      errorCode: ERROR_CODE.TASK_MAXIMUM_LIMIT,
      message: ErrorList[ERROR_CODE.TASK_MAXIMUM_LIMIT].message
    };
  }
  count++;
  updateTaskObj.status === TaskStatusEnum.DONE &&
    (await RedisService.set(task.userId, count.toString()));
  logger.info(
    `processTask >>>> update payload: ${JSON.stringify({
      taskId,
      updateTaskObj
    })}`
  );
  await taskRepository.updateById(taskId, updateTaskObj);
};

export default {
  createTask,
  getById,
  processTask
};
