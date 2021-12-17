const knex = require('./db/knex');

// Validate if the user had created more than 5 tasks per day
module.exports.validateTaskCount = async (req, res, next) => {
    const { reporter_id } = req.body;
    const today = new Date();
    const countRecord = await knex('user_tasks').where({
        'reporter_id': reporter_id,
        'created_date': today.toISOString().split("T")[0],
    }).select('task_count');

    if (countRecord.length > 0 && countRecord[0].task_count >= 5) {
        res.status(403).json({
            success: false,
            message: "Reached task count limit of 5"
        });
    } else if (countRecord.length === 0) {
        next();
    } else {
        next();
    }
}