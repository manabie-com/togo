const express = require('express')
const router = new express.Router()
const userController = require('../controller/user')
const taskController = require('../controller/task')
const { authen } = require('../util/middleware')

// user
router.get('/user/:id', authen, userController.getUser)
router.post('/user/register', userController.register)
router.post('/user/login', userController.login)
// task
router.post('/task', authen, taskController.createTask)

module.exports = router