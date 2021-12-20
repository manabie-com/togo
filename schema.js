const Joi = require('joi');

const taskSchema = Joi.object({
    title: Joi.string().required(),
    reporter_id: Joi.number().required(),
    detail: Joi.string(),
    due_at: Joi.date().timestamp('unix'),
    assignee_id: Joi.number()
})

const userSchema = Joi.object({
    name: Joi.string(),
    email: Joi.string().required()
})

module.exports = {
    taskSchema,
    userSchema
}