const moment = require('moment')
const {
    insertTaskService,
    getTaskByUserIdService,
    updateTaskService,
    deleteTaskService,
    getTaskByIdService
} = require('../services/task')
const {
    validationResult
} = require("express-validator");

const insertTask = async (req, res) => {
    // Create a new task
    try {
        let result = validationResult(req);
        if (result.errors.length !== 0) { //validator   
            let messages = result.mapped();
            let message = "";

            for (m in messages) {
                message = messages[m].msg;
                break;
            }
            throw new Error(message);
        }

        if (await req.user.checkTaskPerDayLimit(req.user._id)) {
            throw new Error("The quantity task of the user is over the limit")
        } else {
            req.body.createdById = req.user._id;
            req.body.createdDate = moment().toDate();
            const task = await insertTaskService(req.body);
            res.status(201).send({
                task
            })
        }
    } catch (error) {
        res.status(400).send({
            error: error.message
        });
    }
}
const getTask = async (req, res) => {
    try {
        let dayQuery
        if (req.query?.day) {
            dayQuery = moment(req.query?.day)
        } else {
            dayQuery = moment()
        }

        const tasks = await getTaskByUserIdService(req.user._id, dayQuery);
        res.status(200).send({
            tasks,
            dayQuery
        })
    } catch (error) {
        res.status(400).send({
            error: error.message
        });
    }
}
const deleteTask = async (req, res) => {
    try {
        const id = req.params?.id
        if (id) {
            const result = await deleteTaskService(id);
            if (result)
                res.status(200).send();
        } else
            res.status(422).send({
                error: "Missing id param"
            });
    } catch (error) {
        res.status(400).send({
            error: error.message
        });
    }
}
const updateTask = async (req, res) => {
    //update task
    try {

        let result = validationResult(req);
        if (result.errors.length !== 0) { //validator   
            let messages = result.mapped();
            let message = "";

            for (m in messages) {
                message = messages[m].msg;
                break;
            }
            throw new Error(message);
        }

        const id = req.params?.id
        if (id) {
            const task = await updateTaskService(id, req.body)
            res.status(200).send(task);
        } else {
            res.status(422).send({
                error: "Missing id param"
            });
        }
    } catch (error) {
        res.status(400).send({
            error: error.message
        });
    }
}
const getTaskById = async (req, res) => {
    //update task
    try {
        const id = req.params?.id
        if (id) {
            let task = await getTaskByIdService(id);
            if (task) {
                res.status(200).send({
                    task
                });
            }
        } else {
            res.status(422).send({
                error: "Missing id param"
            });
        }

    } catch (error) {
        res.status(400).send({
            error: error.message
        });
    }
}
module.exports = {
    insertTask,
    getTask,
    deleteTask,
    updateTask,
    getTaskById
}