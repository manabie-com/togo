import { Types } from 'mongoose';
import taskService from '../task.service';
import userService from '../../user/user.service';
import taskRepository from '../task.repository';
import KafkaService from '../../common/kafka';
import RedisService from '../../common/redis';
import { createTaskPayload } from '../__mock__/task.data';
import { ErrorList, ERROR_CODE } from '../../error/error.list';
import { AppError } from '../../error/error.service';
import { TaskStatusEnum } from '../task.enum';
import { createUserPayload } from '../../user/__mock__/user.data';

jest.mock('../task.repository');
jest.mock('../../user/user.service');
jest.mock('../../common/kafka');
jest.mock('../../common/redis');
jest.mock('redis');

describe('task.service', () => {
  describe('createTask', () => {
    it('Should create task successfully', async () => {
      const taskId = new Types.ObjectId();
      (taskRepository.createTask as jest.Mock).mockResolvedValueOnce({
        ...createTaskPayload,
        _id: taskId
      });
      (userService.getById as jest.Mock).mockResolvedValueOnce({
        _id: createTaskPayload.userId
      });
      (KafkaService.produceMessage as jest.Mock).mockResolvedValueOnce(null);

      const expected = await taskService.createTask(createTaskPayload);
      expect(expected).toEqual(expect.objectContaining(createTaskPayload));
    });

    it(`Should throw ${ERROR_CODE.USER_NOT_FOUND} when not found the user`, async () => {
      (userService.getById as jest.Mock).mockRejectedValueOnce(
        new AppError(ERROR_CODE.USER_NOT_FOUND)
      );

      const expected = await taskService
        .createTask(createTaskPayload)
        .catch((err) => err);

      expect(expected).toBeInstanceOf(AppError);
      expect(expected.errorCode).toEqual(ERROR_CODE.USER_NOT_FOUND);
    });
  });

  describe('getById', () => {
    it('Should get task by id successfully', async () => {
      const taskId = '_taskId';
      (taskRepository.getById as jest.Mock).mockResolvedValueOnce({
        ...createTaskPayload,
        _id: taskId
      });

      const expected = await taskService.getById(taskId);
      expect(expected).toEqual(expect.objectContaining(createTaskPayload));
    });

    it(`Should throw error ${ERROR_CODE.TASK_NOT_FOUND} when not found task`, async () => {
      const taskId = '_test';
      (taskRepository.getById as jest.Mock).mockResolvedValueOnce(null);

      const expected = await taskService.getById(taskId).catch((err) => err);
      expect(expected).toBeInstanceOf(AppError);
      expect(expected.errorCode).toEqual(ERROR_CODE.TASK_NOT_FOUND);
    });
  });

  describe('processTask', () => {
    it(`Should update the status of task ${TaskStatusEnum.DONE} when count is null`, async () => {
      const taskId = '_taskId';
      const userId = '_userId';
      const count = null;
      (taskRepository.getById as jest.Mock).mockResolvedValueOnce({
        ...createTaskPayload,
        _id: taskId,
        userId
      });
      (RedisService.get as jest.Mock).mockResolvedValueOnce(count);
      (RedisService.set as jest.Mock).mockResolvedValueOnce(null);
      (userService.getById as jest.Mock).mockResolvedValueOnce({
        ...createUserPayload,
        _id: userId
      });
      (taskRepository.updateById as jest.Mock).mockResolvedValueOnce(null);

      await taskService.processTask(taskId);

      expect(RedisService.set).toBeCalledWith(userId, '1');
      expect(taskRepository.updateById).toBeCalledWith(taskId, {
        status: TaskStatusEnum.DONE
      });
    });

    it(`Should update the status of task ${TaskStatusEnum.DONE} when count < limit`, async () => {
      const taskId = '_taskId';
      const userId = '_userId';
      const count = 1;
      (taskRepository.getById as jest.Mock).mockResolvedValueOnce({
        ...createTaskPayload,
        _id: taskId,
        userId
      });
      (RedisService.get as jest.Mock).mockResolvedValueOnce(count);
      (RedisService.set as jest.Mock).mockResolvedValueOnce(null);
      (userService.getById as jest.Mock).mockResolvedValueOnce({
        ...createUserPayload,
        _id: userId
      });
      (taskRepository.updateById as jest.Mock).mockResolvedValueOnce(null);

      await taskService.processTask(taskId);

      expect(RedisService.set).toBeCalledWith(userId, (count + 1).toString());
      expect(taskRepository.updateById).toBeCalledWith(taskId, {
        status: TaskStatusEnum.DONE
      });
    });

    it(`Should update the status of task ${TaskStatusEnum.FAILED} when count > limit`, async () => {
      const taskId = '_taskId';
      const userId = '_userId';
      const count = 1;
      const limit = 1;
      (taskRepository.getById as jest.Mock).mockResolvedValueOnce({
        ...createTaskPayload,
        _id: taskId,
        userId
      });
      (RedisService.get as jest.Mock).mockResolvedValueOnce(count);
      (RedisService.set as jest.Mock).mockResolvedValueOnce(null);
      (userService.getById as jest.Mock).mockResolvedValueOnce({
        ...createUserPayload,
        _id: userId,
        configuration: {
          ...createUserPayload.configuration,
          limit
        }
      });
      (taskRepository.updateById as jest.Mock).mockResolvedValueOnce(null);

      await taskService.processTask(taskId);

      expect(RedisService.set).not.toBeCalled();
      expect(taskRepository.updateById).toBeCalledWith(taskId, {
        status: TaskStatusEnum.FAILED,
        reason: {
          errorCode: ERROR_CODE.TASK_MAXIMUM_LIMIT,
          message: ErrorList[ERROR_CODE.TASK_MAXIMUM_LIMIT].message
        }
      });
    });
  });
});
