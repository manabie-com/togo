import { ERROR_CODE } from '../error/error.list';
import { AppError } from '../error/error.service';
import userRepository from './user.repository';
import { ICreateUserPayload, IUser } from './user.type';

const createUser = async (payload: ICreateUserPayload): Promise<IUser> => {
  try {
    const user = await userRepository.createUser(payload);
    return user;
  } catch (err) {
    if ((err as any).code === 11000) {
      throw new AppError(ERROR_CODE.DUPLICATE_USER);
    }
    throw new AppError(ERROR_CODE.CREATE_USER_ERROR);
  }
};

const getById = async (userId: string): Promise<IUser> => {
  const user = await userRepository.getById(userId);
  if (!user) {
    throw new AppError(ERROR_CODE.USER_NOT_FOUND);
  }
  return user;
};

export default {
  createUser,
  getById
};
