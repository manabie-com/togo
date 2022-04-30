import hapi from '@hapi/hapi';
import { StatusCode } from '../../common/enum';
import taskController from '../task.controller';
import taskService from '../task.service';
import { createTaskPayload } from '../__mock__/task.data';

jest.mock('../task.service');

describe('task.controller', () => {
  let server: hapi.Server;
  beforeAll(async () => {
    server = new hapi.Server();
    server.route(taskController);
  });

  describe('POST /task', () => {
    it(`Should return status ${StatusCode.CREATED} when creating successfully`, async () => {
      const options = {
        method: 'POST',
        url: '/task',
        payload: {
          userId: '_userId',
          name: '_name'
        }
      };

      (taskService.createTask as jest.Mock).mockResolvedValueOnce(
        createTaskPayload
      );

      const response = await server.inject(options);
      expect(response.statusCode).toEqual(StatusCode.CREATED);
    });

    it(`Should return status ${StatusCode.BAD_REQUEST} when wrong input payload`, async () => {
      const options = {
        method: 'POST',
        url: '/task',
        payload: {}
      };

      (taskService.createTask as jest.Mock).mockResolvedValueOnce(
        createTaskPayload
      );

      const response = await server.inject(options);
      expect(response.statusCode).toEqual(StatusCode.BAD_REQUEST);
    });
  });
});
