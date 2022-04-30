import hapi from '@hapi/hapi';
import { StatusCode } from '../common/enum';
import logger from '../logger';
import taskService from './task.service';
import { ICreateTaskPayload } from './task.type';
import { createTaskPayloadValidator } from './task.validator';

const createTask: hapi.ServerRoute = {
  method: 'POST',
  path: '/task',
  options: {
    description: 'Create new task',
    tags: ['api', 'task'],
    validate: {
      payload: createTaskPayloadValidator
    },
    handler: async (req, res) => {
      logger.info('createTask >>>>');
      await taskService.createTask(req.payload as ICreateTaskPayload);
      return res.response().code(StatusCode.CREATED);
    }
  }
};

const taskController: hapi.ServerRoute[] = [createTask];

export default taskController;
