import { isMongoObjectId } from '../common/util';
import userModel from './user.model';
import { ICreateUserPayload, IUser } from './user.type';

const createUser = async (payload: ICreateUserPayload): Promise<IUser> => {
  const user: IUser = await userModel.create(payload);
  return user;
};

const getById = async (userId: string): Promise<IUser | null> => {
  if (!isMongoObjectId(userId)) {
    return null;
  }
  const user: IUser = await userModel
    .findOne({
      _id: userId
    })
    .exec();
  return user;
};

export default {
  createUser,
  getById
};
