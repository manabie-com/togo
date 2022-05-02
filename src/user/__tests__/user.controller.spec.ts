import hapi from '@hapi/hapi';
import ResponseWrapper from '../../plugins/responseWrapper.plugin';
import { StatusCode } from '../../common/enum';
import taskService from '../../task/task.service';
import userController from '../user.controller';
import { UserConfigurationEnum } from '../user.enum';
import userService from '../user.service';
import { createUserPayload } from '../__mock__/user.data';
import { createTaskPayloadFn } from '../../task/__mock__/task.data';
import { TaskStatusEnum } from '../../task/task.enum';
import { ERROR_CODE } from '../../error/error.list';
import { AppError } from '../../error/error.service';

jest.mock('../user.service');
jest.mock('../../task/task.service');

describe('user.controller', () => {
  let server: hapi.Server;
  beforeAll(async () => {
    server = new hapi.Server();
    server.register([ResponseWrapper]);
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
            type: UserConfigurationEnum.DAILY
          }
        }
      };

      (userService.createUser as jest.Mock).mockResolvedValueOnce(
        createUserPayload
      );

      const response = await server.inject(options);

      expect(response.statusCode).toEqual(StatusCode.CREATED);
      expect(response.result).toEqual({
        data: { userId: createUserPayload.id }
      });
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

  describe('GET /user/{userId}/tasks', () => {
    it('Should return tasks by user id successfully', async () => {
      const userId = '_userId';
      const taskId = '_taskId';
      const tasks = [
        createTaskPayloadFn({ userId, id: taskId, status: TaskStatusEnum.DONE })
      ];
      const options = {
        method: 'GET',
        url: `/user/${userId}/tasks`
      };

      (userService.getById as jest.Mock).mockResolvedValueOnce(null);
      (taskService.getsByUserId as jest.Mock).mockResolvedValueOnce(tasks);

      const response = await server.inject(options);

      expect(response.statusCode).toEqual(StatusCode.OK);
      expect(response.result).toEqual({ data: tasks });
    });

    it(`Should throw ${ERROR_CODE.USER_NOT_FOUND} when not found the user`, async () => {
      const userId = '_userId';
      const options = {
        method: 'GET',
        url: `/user/${userId}/tasks`
      };

      (userService.getById as jest.Mock).mockRejectedValueOnce(
        new AppError(ERROR_CODE.USER_NOT_FOUND)
      );

      const response = await server.inject(options);

      expect(response.statusCode).toEqual(StatusCode.BAD_REQUEST);
      expect(response.result).toEqual(
        expect.objectContaining({
          statusCode: StatusCode.BAD_REQUEST,
          message: ERROR_CODE.USER_NOT_FOUND
        })
      );
    });
  });
});
