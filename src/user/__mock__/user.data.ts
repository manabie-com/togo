import { UserConfigurationEnum } from '../user.enum';

export const createUserPayload = {
  id: '_id',
  username: 'username',
  password: 'password',
  configuration: {
    limit: 100,
    count: 0,
    type: UserConfigurationEnum.DAILY
  }
};
