import Joi from 'joi';
import { IUserConfigurationEnum } from './user.enum';
import { ICreateUserPayload } from './user.type';

export const createUserPayloadValidator = Joi.object<ICreateUserPayload>({
  username: Joi.string().required(),
  password: Joi.string().required(),
  configuration: Joi.object({
    type: Joi.string().valid(...Object.keys(IUserConfigurationEnum)),
    limit: Joi.number().required()
  })
});
