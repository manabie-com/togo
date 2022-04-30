import taskService from '../task.service';
import userService from '../../user/user.service';
import taskRepository from '../task.repository';
import { createTaskPayload } from '../__mock__/task.data';
import { ERROR_CODE } from '../../error/error.list';
import { AppError } from '../../error/error.service';

jest.mock('../task.repository');
jest.mock('../../user/user.service');

describe('task.service', () => {
  describe('createTask', () => {
    it('Should create task successfully', async () => {
      (taskRepository.createTask as jest.Mock).mockResolvedValueOnce(
        createTaskPayload
      );
      (userService.getById as jest.Mock).mockResolvedValueOnce({
        _id: createTaskPayload.userId
      });

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
      const userId = '_test';
      (taskRepository.getById as jest.Mock).mockResolvedValueOnce({
        ...createTaskPayload,
        _id: userId
      });

      const expected = await taskService.getById(userId);
      expect(expected).toEqual(expect.objectContaining(createTaskPayload));
    });

    it(`Should throw error ${ERROR_CODE.TASK_NOT_FOUND} when not found task`, async () => {
      const userId = '_test';
      (taskRepository.getById as jest.Mock).mockResolvedValueOnce(null);

      const expected = await taskService.getById(userId).catch((err) => err);
      expect(expected).toBeInstanceOf(AppError);
      expect(expected.errorCode).toEqual(ERROR_CODE.TASK_NOT_FOUND);
    });
  });
});
