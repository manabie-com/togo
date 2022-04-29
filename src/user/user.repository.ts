import userModel from './user.model';
import { ICreateUserPayload, IUser } from './user.type';

const createUser = async (payload: ICreateUserPayload): Promise<IUser> => {
  const user: IUser = await userModel.create(payload);
  return user;
};

export default {
  createUser
};
