import { Types } from 'mongoose';
import userRepository from '../user.repository';
import userModel from '../user.model';

import { createUserPayload } from '../__mock__/user.data';

jest.mock('../user.model', () => ({
  create: jest.fn(),
  findOne: jest.fn()
}));

describe('user.repository', () => {
  describe('createUser', () => {
    it('Should create user successfully', async () => {
      (userModel.create as unknown as jest.Mock).mockResolvedValueOnce({
        _id: 1,
        ...createUserPayload
      });

      const expected = await userRepository.createUser(createUserPayload);
      expect(expected).toEqual(expect.objectContaining(createUserPayload));
    });
  });

  describe('getById', () => {
    it('Should get user by id successfully', async () => {
      const userId = new Types.ObjectId();
      (userModel.findOne as unknown as jest.Mock).mockImplementationOnce(
        () => ({
          exec: jest.fn().mockResolvedValueOnce({
            _id: userId,
            ...createUserPayload
          })
        })
      );

      const expected = await userRepository.getById(userId.toString());
      expect(expected).toEqual(expect.objectContaining(createUserPayload));
    });

    it('Should return null when not found user by id', async () => {
      const userId = new Types.ObjectId();
      (userModel.findOne as unknown as jest.Mock).mockImplementationOnce(
        () => ({
          exec: jest.fn().mockResolvedValueOnce(null)
        })
      );

      const expected = await userRepository.getById(userId.toString());
      expect(expected).toBeNull();
    });

    it('Should return null when invalid user id', async () => {
      const expected = await userRepository.getById('_test');
      expect(expected).toBeNull();
    });
  });
});
