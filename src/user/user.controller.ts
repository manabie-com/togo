import hapi from '@hapi/hapi';
import { StatusCode } from '../common/enum';
import logger from '../logger';
import taskService from '../task/task.service';
import userService from './user.service';
import { ICreateUserPayload } from './user.type';
import {
  createUserPayloadValidator,
  createUserResponseValidator,
  getTasksByUserQueryValidator,
  getTasksByUserResponseValidator
} from './user.validator';

const createUser: hapi.ServerRoute = {
  method: 'POST',
  path: '/user',
  options: {
    description: 'Create new user',
    tags: ['api', 'user'],
    validate: {
      payload: createUserPayloadValidator
    },
    response: {
      schema: createUserResponseValidator
    },
    handler: async (req, res) => {
      logger.info('createUser >>>>');
      const user = await userService.createUser(
        req.payload as ICreateUserPayload
      );

      return res.response({ userId: user.id }).code(StatusCode.CREATED);
    }
  }
};

const getTasksByUser: hapi.ServerRoute = {
  method: 'GET',
  path: '/user/{userId}/tasks',
  options: {
    description: 'Get list tasks by user',
    tags: ['api', 'user'],
    validate: {
      query: getTasksByUserQueryValidator
    },
    response: {
      schema: getTasksByUserResponseValidator
    },
    handler: async (req, res) => {
      logger.info('getTasksByUser >>>>');
      const {
        params: { userId },
        query
      } = req;

      await userService.getById(userId);
      const tasks = await taskService.getsByUserId({ ...query, userId });

      return res.response(tasks).code(StatusCode.OK);
    }
  }
};

const userController: hapi.ServerRoute[] = [createUser, getTasksByUser];

export default userController;
