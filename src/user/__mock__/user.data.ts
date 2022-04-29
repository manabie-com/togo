import { IUserConfigurationEnum } from '../user.enum';

export const createUserPayload = {
  username: 'username',
  password: 'password',
  configuration: {
    limit: 100,
    count: 0,
    type: IUserConfigurationEnum.DAILY
  }
};
