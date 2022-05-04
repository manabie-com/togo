import { IBaseModel } from '../common/type';
import { UserConfigurationEnum } from './user.enum';

export interface IUserConfiguration {
  type: UserConfigurationEnum;
  limit: number;
}

export interface IUser extends IBaseModel {
  username: string;
  password: string;
  configuration: IUserConfiguration;
}

export interface ICreateUserPayload {
  username: string;
  password: string;
  configuration: IUserConfiguration;
}
