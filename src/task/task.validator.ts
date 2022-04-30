import Joi from 'joi';
import { ICreateTaskPayload } from './task.type';

export const createTaskPayloadValidator = Joi.object<ICreateTaskPayload>({
  name: Joi.string().required(),
  userId: Joi.string().required()
});
