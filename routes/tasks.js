const express = require('express');
const router = express.Router();
const tasks = require('../controllers/tasks');
const { validateTask, validateTaskCount } = require('../middleware/tasks');

router.route('/')
    .get(tasks.getAllTasks)
    .post(validateTask, validateTaskCount, tasks.createTask);

router.route('/:id')
    .patch(tasks.updateTask)
    .delete(tasks.deleteTask);

router.route('/count').get(tasks.getTaskCount);

module.exports = router;