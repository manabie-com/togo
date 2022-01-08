const taskModel = require('../models/model.task');
const joi = require('joi');

const taskSchema = joi.object({
    userId: joi.string().min(16).max(128).required(),
    title: joi.string().required(),
    description: joi.string().required(),
    status: joi.number().required().min(0)
});