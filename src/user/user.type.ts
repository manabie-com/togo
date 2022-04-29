import { IUserConfigurationEnum } from './user.enum';

export interface IUserConfiguration {
  type: IUserConfigurationEnum;
  limit: number;
  count: number;
}

export interface IUser {
  username: string;
  password: string;
  configuration: IUserConfiguration;
}

export interface ICreateUserPayload {
  username: string;
  password: string;
  configuration: IUserConfiguration;
}
