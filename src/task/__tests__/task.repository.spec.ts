import { Types } from 'mongoose';
import taskRepository from '../task.repository';
import taskModel from '../task.model';

import { createTaskPayload } from '../__mock__/task.data';

jest.mock('../task.model', () => ({
  create: jest.fn(),
  findOne: jest.fn()
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
      const userId = new Types.ObjectId();
      (taskModel.findOne as unknown as jest.Mock).mockImplementationOnce(
        () => ({
          exec: jest.fn().mockResolvedValueOnce({
            _id: userId,
            ...createTaskPayload
          })
        })
      );

      const expected = await taskRepository.getById(userId.toString());
      expect(expected).toEqual(expect.objectContaining(createTaskPayload));
    });

    it('Should return null when not found user by id', async () => {
      const userId = new Types.ObjectId();
      (taskModel.findOne as unknown as jest.Mock).mockImplementationOnce(
        () => ({
          exec: jest.fn().mockResolvedValueOnce(null)
        })
      );

      const expected = await taskRepository.getById(userId.toString());
      expect(expected).toBeNull();
    });

    it('Should return null when invalid user id', async () => {
      const expected = await taskRepository.getById('_test');
      expect(expected).toBeNull();
    });
  });
});
