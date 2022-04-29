import userRepository from '../user.repository';
import userModel from '../user.model';

import { createUserPayload } from '../__mock__/user.data';

jest.mock('../user.model', () => ({
  create: jest.fn()
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
});
