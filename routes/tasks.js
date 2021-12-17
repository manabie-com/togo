const express = require('express');
const router = express.Router();
const tasks = require('../controllers/tasks');
const { validateTaskCount } = require('../middleware');

router.route('/')
    .get(tasks.getAllTasks)
    .post(validateTaskCount, tasks.createTask);

router.route('/:id')
    .patch(tasks.updateTask)
    .delete(tasks.deleteTask);

router.route('/count').get(tasks.getTaskCount);

module.exports = router;