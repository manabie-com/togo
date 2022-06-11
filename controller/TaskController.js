const express       = require('express');
const app           = express();
const TaskClass = require('../classes/TaskClass');
const jwt = require("jsonwebtoken")

const task = new TaskClass();

module.exports =
{
    async addTask(req, res)
    {
        let token = req.headers['x-auth-token'];
        if (!token) return res.status(401).json({
            status: 401,
            message: "You are not logged in"
        })

        let decoded = jwt.verify(token, 'akaru-todo');
        req.body.userName = decoded.username;
        let data = req.body;
        
        let task_class = new TaskClass(data);

        let validation   = await task_class.createTask();
        res.send(validation);

    },

    async getTask(req, res)
    {
        
        let data = req.body;
        let task_class = new TaskClass(data);

        let validation = await task_class.getTask();
        console.log(validation.data);
        res.send(validation);
    }
}