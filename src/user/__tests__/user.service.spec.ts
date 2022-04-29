import userService from '../user.service';
import userRepository from '../user.repository';
import { createUserPayload } from '../__mock__/user.data';
import { ERROR_CODE } from '../../error/error.list';
import { AppError } from '../../error/error.service';

jest.mock('../user.repository');

describe('user.service', () => {
  describe('createUser', () => {
    it('Should create user successfully', async () => {
      (userRepository.createUser as jest.Mock).mockResolvedValueOnce(
        createUserPayload
      );

      const expected = await userService.createUser(createUserPayload);
      expect(expected).toEqual(expect.objectContaining(createUserPayload));
    });

    it(`Should throw ${ERROR_CODE.DUPLICATE_USER} when input duplicate username`, async () => {
      (userRepository.createUser as jest.Mock).mockRejectedValueOnce({
        code: 11000
      });

      const expected = await userService
        .createUser(createUserPayload)
        .catch((err) => err);

      expect(expected).toBeInstanceOf(AppError);
      expect(expected.errorCode).toEqual(ERROR_CODE.DUPLICATE_USER);
    });

    it(`Should throw ${ERROR_CODE.CREATE_USER_ERROR} when input wrong payload`, async () => {
      (userRepository.createUser as jest.Mock).mockRejectedValueOnce({});

      const expected = await userService
        .createUser(createUserPayload)
        .catch((err) => err);

      expect(expected).toBeInstanceOf(AppError);
      expect(expected.errorCode).toEqual(ERROR_CODE.CREATE_USER_ERROR);
    });
  });
});
