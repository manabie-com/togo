const express = require('express')
const router = express.Router()
const TaskService = require('../services/task')
const ConfigService = require('../services/config')
const UserService = require('../services/users')
const { Task } = require('../models')
const { User } = require('../models')
const { Config } = require('../models')

const taskService = new TaskService(Task)
const configService = new ConfigService(Config)
const userService = new UserService(User)

router.post('/', async(req, res) => {
    try {
        const { title, description, text } = req.body
            // const user = await userService.getById(req.userId)
            // console.log('user', user)
            // const config = await configService.getLimitByRole(user.role)
            // console.log('config', config)
        await taskService.create({ title, description, text, author: req.userId })
        res.status(201).json({ title, description, text })
    } catch (err) {
        console.log('err', err)
        res.status(400).send(err.message)
    }
})

module.exports = router