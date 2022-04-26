const Joi = require('joi');

const createTask = {
  body: Joi.object().keys({
    task_name: Joi.string().required(),
    user_id: Joi.number().required(),
    task_priority: Joi.number().optional(),
  }),
};

const updateTask = {
  params: Joi.object().keys({
    taskId: Joi.number().required(),
  }),
  body: Joi.object().keys({
    task_name: Joi.string().required(),
    task_priority: Joi.number().optional(),
  }),
};

const deleteTask = {
  params: Joi.object().keys({
    taskId: Joi.number().required(),
  }),
};

module.exports = {
  deleteTask,
  updateTask,
  createTask,
};
