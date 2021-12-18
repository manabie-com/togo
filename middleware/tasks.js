const knex = require('../db/knex').knex;
const { taskSchema } = require('../schema');
const { TASK_LIMIT } = require('../config');

// Convert datetime string to unix timestamp
function toUnixTimestamp(datetimeString) {
    return Math.round(new Date(datetimeString).getTime() / 1000);
}

// Validate the task body
const validateTask = async (req, res, next) => {
    if (req.body.due_at) {
        req.body.due_at = toUnixTimestamp(req.body.due_at);
    }

    const { error } = taskSchema.validate(req.body);
    if (error) {
        res.status(400).json(error);
    } else {
        next();
    }
}

// Validate if the user had created more than TASK_LIMIT tasks per day
const validateTaskCount = async (req, res, next) => {
    const { reporter_id } = req.body;

    const today = new Date();
    const countRecord = await knex('user_tasks').where({
        'reporter_id': reporter_id,
        'created_date': today.toISOString().split("T")[0],
    }).select('task_count');

    if (countRecord.length > 0 && countRecord[0].task_count >= TASK_LIMIT) {
        res.status(403).json({
            success: false,
            message: `Reached task count limit of ${TASK_LIMIT}`
        });
    } else if (countRecord.length === 0) {
        next();
    } else {
        next();
    }
}

module.exports = {
    validateTask,
    validateTaskCount
}