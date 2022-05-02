import Joi from 'joi';
import { TaskStatusEnum } from '../task/task.enum';
import { TaskValidator } from '../task/task.validator';
import { UserConfigurationEnum } from './user.enum';
import { ICreateUserPayload } from './user.type';

export const createUserPayloadValidator = Joi.object<ICreateUserPayload>({
  username: Joi.string().required(),
  password: Joi.string().required(),
  configuration: Joi.object({
    type: Joi.string().valid(...Object.keys(UserConfigurationEnum)),
    limit: Joi.number().required().integer().greater(0)
  })
});

export const createUserResponseValidator = Joi.object({
  userId: Joi.string().required()
});

export const getTasksByUserQueryValidator = Joi.object({
  status: Joi.string().valid(...Object.keys(TaskStatusEnum), 'ALL')
});

export const getTasksByUserResponseValidator = Joi.array().items(TaskValidator);
