const knex = require('../db/knex');

module.exports.getAllTasks = async (req, res) => {
    const tasks = await knex('tasks').select('*');
    res.status(200).json(tasks);
}

module.exports.createTask = async (req, res, next) => {
    // Insert new task
    const results = await knex('tasks').insert(req.body);

    // ... and then update the task count of that user
    const { reporter_id } = req.body;
    const created_at = await knex('tasks').where('id', results[0]).select('created_at');
    const created_date = created_at[0].created_at.split(" ")[0];
    const countRecord = await knex('user_tasks').where({
        'reporter_id': reporter_id,
        'created_date': created_date
    }).select('task_count');

    if (countRecord.length > 0) {
        // add 1 to the current task count
        await knex('user_tasks').where({
            'reporter_id': reporter_id,
            'created_date': created_date
        }).update({ 'task_count': countRecord[0].task_count + 1 });
    } else {
        // initialize the task count with 1
        await knex('user_tasks').insert({
            'created_date': created_date,
            'reporter_id': reporter_id,
            'task_count': 1,
        });
    }
    const newInsert = await knex('tasks').where('id', results[0]).select('*');
    res.status(201).json(newInsert[0]);
}

module.exports.updateTask = async (req, res) => {
    const id = await knex('tasks').where('id', req.params.id).update(req.body);
    const newUpdate = await knex('tasks').where('id', id).select('*');
    res.status(200).json(newUpdate);
}

module.exports.deleteTask = async (req, res) => {
    await knex('tasks').where('id', req.params.id).del();
    res.status(200).json({ success: true });
}

module.exports.getTaskCount = async (req, res) => {
    const results = await knex('user_tasks').select('*');
    res.status(200).json(results);
}