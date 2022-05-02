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
}).label('CreateUser - Payload');

export const createUserResponseValidator = Joi.object({
  userId: Joi.string().required()
}).label('createUser - Response');

export const userParamValidator = Joi.object({
  userId: Joi.string().required()
}).label('UserId - Params');

export const getTasksByUserQueryValidator = Joi.object({
  status: Joi.string().valid(...Object.keys(TaskStatusEnum), 'ALL')
}).label('getTasksByUser - Query');

export const getTasksByUserResponseValidator = Joi.array()
  .items(TaskValidator)
  .label('getTasksByUser - Response');

export const userConfigurationValidator = Joi.object({
  type: Joi.string()
    .valid(...Object.keys(UserConfigurationEnum))
    .required(),
  limit: Joi.number().integer().greater(0)
}).label('User - Configuration');

export const userValidator = Joi.object({
  id: Joi.string().required(),
  username: Joi.string().required(),
  password: Joi.string().required(),
  configuration: userConfigurationValidator
}).label('User');
