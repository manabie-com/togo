import hapi from '@hapi/hapi';
import { StatusCode } from '../common/enum';
import logger from '../logger';
import taskService from '../task/task.service';
import { ICreateTaskPayload } from '../task/task.type';
import { createTaskPayloadValidator } from '../task/task.validator';
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

const createTask: hapi.ServerRoute = {
  method: 'POST',
  path: '/user/{userId}/task',
  options: {
    description: 'Create new task',
    tags: ['api', 'user'],
    validate: {
      payload: createTaskPayloadValidator
    },
    handler: async (req, res) => {
      logger.info('createTask >>>>');
      const { payload, params } = req;
      const createTaskPayload = {
        ...(payload as object),
        userId: params.userId
      } as ICreateTaskPayload;
      await taskService.createTask(createTaskPayload);
      return res.response().code(StatusCode.CREATED);
    }
  }
};

const userController: hapi.ServerRoute[] = [
  createUser,
  getTasksByUser,
  createTask
];

export default userController;
