import Joi from 'joi';
import { TaskStatusEnum } from './task.enum';

export const createTaskPayloadValidator = Joi.object({
  name: Joi.string().required()
}).label('createTask - Payload');

export const taskReasonValidator = Joi.object({
  errorCode: Joi.string().optional(),
  message: Joi.string().optional()
}).label('Task - Reason');

export const TaskValidator = Joi.object({
  id: Joi.string().required(),
  userId: Joi.string().required(),
  name: Joi.string().required(),
  status: Joi.string()
    .valid(...Object.keys(TaskStatusEnum))
    .required(),
  reason: taskReasonValidator
}).label('Task');
