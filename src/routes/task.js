const express = require('express')
const router = express.Router()
const TaskService = require('../services/task')
const { Task } = require('../models')

const taskService = new TaskService(Task)
router.post('/', async(req, res) => {
    try {
        const { title, description, text } = req.body
        await taskService.create({ title, description, text, author: req.userId })
        res.status(201).json({ title, description, text })
    } catch (err) {
        console.log('err', err)
        res.status(400).send(err.message)
    }
})

module.exports = router