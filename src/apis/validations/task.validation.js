const Joi = require("joi");

const { priorityTypes } = require("../../configs/priority");
const { objectId } = require("./customize.validation");

const taskSchema = {
  body: Joi.object().keys({
    title: Joi.string().required(),
    description: Joi.string(),
    priority: Joi.string().valid(...Object.values(priorityTypes)),
    completed: Joi.boolean().default(false),
  }),
};

const objectIdSchema = {
  params: Joi.object().keys({
    id: Joi.string().required().custom(objectId),
  }),
};

const updateTaskSchema = {
  ...objectIdSchema,
  body: Joi.object().keys({
    title: Joi.string(),
    description: Joi.string(),
    priority: Joi.string().valid(...Object.values(priorityTypes)),
    completed: Joi.boolean().default(false),
  }),
};

module.exports = {
  taskSchema,
  objectIdSchema,
  updateTaskSchema,
};
