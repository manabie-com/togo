import hapi from '@hapi/hapi';
import { StatusCode } from '../../common/enum';
import userController from '../user.controller';
import { IUserConfigurationEnum } from '../user.enum';
import userService from '../user.service';
import { createUserPayload } from '../__mock__/user.data';

jest.mock('../user.service');

describe('user.controller', () => {
  let server: hapi.Server;
  beforeAll(async () => {
    server = new hapi.Server();
    server.route(userController);
  });

  describe('POST /user', () => {
    it(`Should return status ${StatusCode.CREATED} when creating successfully`, async () => {
      const options = {
        method: 'POST',
        url: '/user',
        payload: {
          username: 'username',
          password: 'password',
          configuration: {
            limit: 100,
            type: IUserConfigurationEnum.DAILY
          }
        }
      };

      (userService.createUser as jest.Mock).mockResolvedValueOnce(
        createUserPayload
      );

      const response = await server.inject(options);
      expect(response.statusCode).toEqual(StatusCode.CREATED);
    });

    it(`Should return status ${StatusCode.BAD_REQUEST} when wrong input payload`, async () => {
      const options = {
        method: 'POST',
        url: '/user',
        payload: {}
      };

      (userService.createUser as jest.Mock).mockResolvedValueOnce(
        createUserPayload
      );

      const response = await server.inject(options);
      expect(response.statusCode).toEqual(StatusCode.BAD_REQUEST);
    });
  });
});
