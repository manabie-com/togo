import hapi from '@hapi/hapi';
import { StatusCode } from '../common/enum';
import logger from '../logger';
import userService from './user.service';
import { ICreateUserPayload } from './user.type';
import { createUserPayloadValidator } from './user.validator';

const createUser: hapi.ServerRoute = {
  method: 'POST',
  path: '/user',
  options: {
    description: 'Create new user',
    tags: ['api', 'user'],
    validate: {
      payload: createUserPayloadValidator
    },
    handler: async (req, res) => {
      logger.info('createUser >>>>');
      await userService.createUser(req.payload as ICreateUserPayload);
      return res.response().code(StatusCode.CREATED);
    }
  }
};

const userController: hapi.ServerRoute[] = [createUser];

export default userController;
