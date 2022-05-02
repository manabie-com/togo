import Joi from 'joi';
import { TaskStatusEnum } from './task.enum';
import { ICreateTaskPayload } from './task.type';

export const createTaskPayloadValidator = Joi.object<ICreateTaskPayload>({
  name: Joi.string().required(),
  userId: Joi.string().required()
});

export const TaskValidator = Joi.object({
  id: Joi.string().required(),
  userId: Joi.string().required(),
  name: Joi.string().required(),
  status: Joi.string()
    .valid(...Object.keys(TaskStatusEnum))
    .required(),
  reason: {
    errorCode: Joi.string().optional(),
    message: Joi.string().optional()
  }
});
