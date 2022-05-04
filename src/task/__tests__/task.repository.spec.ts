import { Types } from 'mongoose';
import taskRepository from '../task.repository';
import taskModel from '../task.model';

import { createTaskPayload } from '../__mock__/task.data';
import { TaskStatusEnum } from '../task.enum';
import { ERROR_CODE } from '../../error/error.list';
import { AppError } from '../../error/error.service';
import { documentToObject } from '../../common/util';

jest.mock('../task.model', () => ({
  create: jest.fn(),
  findOne: jest.fn(),
  findByIdAndUpdate: jest.fn(),
  find: jest.fn()
}));

describe('task.repository', () => {
  describe('createTask', () => {
    it('Should create task successfully', async () => {
      (taskModel.create as unknown as jest.Mock).mockResolvedValueOnce({
        _id: 1,
        ...createTaskPayload
      });

      const expected = await taskRepository.createTask(createTaskPayload);
      expect(expected).toEqual(expect.objectContaining(createTaskPayload));
    });
  });

  describe('getById', () => {
    it('Should get task by id successfully', async () => {
      const taskId = new Types.ObjectId();
      (taskModel.findOne as unknown as jest.Mock).mockImplementationOnce(
        () => ({
          exec: jest.fn().mockResolvedValueOnce({
            _id: taskId,
            ...createTaskPayload
          })
        })
      );

      const expected = await taskRepository.getById(taskId.toString());
      expect(expected).toEqual(expect.objectContaining(createTaskPayload));
    });

    it('Should return null when not found user by id', async () => {
      const taskId = new Types.ObjectId();
      (taskModel.findOne as unknown as jest.Mock).mockImplementationOnce(
        () => ({
          exec: jest.fn().mockResolvedValueOnce(null)
        })
      );

      const expected = await taskRepository.getById(taskId.toString());
      expect(expected).toBeNull();
    });

    it('Should return null when invalid user id', async () => {
      const expected = await taskRepository.getById('_test');
      expect(expected).toBeNull();
    });
  });

  describe('updateById', () => {
    it('Should update the task successfully', async () => {
      const taskId = new Types.ObjectId();
      const status = TaskStatusEnum.DONE;
      const updatedObj = {
        _id: taskId,
        ...createTaskPayload,
        status
      };
      (
        taskModel.findByIdAndUpdate as unknown as jest.Mock
      ).mockImplementationOnce(() => ({
        exec: jest.fn().mockResolvedValueOnce(updatedObj)
      }));

      const expected = await taskRepository.updateById(taskId.toString(), {
        status
      });

      expect(taskModel.findByIdAndUpdate).toBeCalledWith(
        taskId.toString(),
        { status },
        { new: true }
      );
      expect(expected).toEqual(updatedObj);
    });

    it(`Should throw error ${ERROR_CODE.TASK_NOT_FOUND} when taskId is null`, async () => {
      const expected = await taskRepository
        .updateById(null, {
          status: TaskStatusEnum.DONE
        })
        .catch((err) => err);

      expect(expected).toBeInstanceOf(AppError);
      expect(expected.errorCode).toBe(ERROR_CODE.TASK_NOT_FOUND);
    });

    it(`Should throw error ${ERROR_CODE.TASK_NOT_FOUND} when not found the task`, async () => {
      const taskId = new Types.ObjectId();
      const status = TaskStatusEnum.DONE;
      const updatedObj = {
        _id: taskId,
        ...createTaskPayload,
        status
      };
      (
        taskModel.findByIdAndUpdate as unknown as jest.Mock
      ).mockImplementationOnce(() => ({
        exec: jest.fn().mockResolvedValueOnce(null)
      }));

      const expected = await taskRepository
        .updateById(taskId.toString(), {
          status
        })
        .catch((err) => err);

      expect(taskModel.findByIdAndUpdate).toBeCalledWith(
        taskId.toString(),
        { status },
        { new: true }
      );
      expect(expected).toBeInstanceOf(AppError);
      expect(expected.errorCode).toBe(ERROR_CODE.TASK_NOT_FOUND);
    });
  });

  describe('getsByUserId', () => {
    it('Should get tasks by user id successfully', async () => {
      const taskId = new Types.ObjectId();
      const tasks = [
        {
          _id: taskId,
          ...createTaskPayload
        }
      ];
      (taskModel.find as unknown as jest.Mock).mockImplementationOnce(() => ({
        lean: jest.fn().mockImplementationOnce(() => ({
          exec: jest.fn().mockResolvedValueOnce(tasks)
        }))
      }));

      const expected = await taskRepository.getsByUserId({
        userId: createTaskPayload.userId
      });

      expect(expected).toEqual(tasks.map((task) => documentToObject(task)));
    });

    it('Should omit the status property on query when the status is ALL', async () => {
      const taskId = new Types.ObjectId();
      const tasks = [
        {
          _id: taskId,
          ...createTaskPayload
        }
      ];
      (taskModel.find as unknown as jest.Mock).mockImplementationOnce(() => ({
        lean: jest.fn().mockImplementationOnce(() => ({
          exec: jest.fn().mockResolvedValueOnce(tasks)
        }))
      }));

      const expected = await taskRepository.getsByUserId({
        userId: createTaskPayload.userId,
        status: 'ALL'
      });

      expect(taskModel.find).toBeCalledWith({
        userId: createTaskPayload.userId
      });
      expect(expected).toEqual(tasks.map((task) => documentToObject(task)));
    });
  });
});
