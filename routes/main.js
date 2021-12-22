`use strict`;

const routers = require('express').Router();
const taskController = require('../controllers/task');

//create new task
routers.route('/v1/create', async (req, res) => {
    try {
        let userID = req.body.userID;
        let taskData = req.body.data;
        let date = req.body.date;

        let result = await taskController.createTask(userID, taskData, date);
        return res.send({data: result, error: 0});

    } catch (err) {
        return res.send({error: 1, error_message: err.message});
    }
});

module.exports = routers;