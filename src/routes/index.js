const express = require('express')
const router = express.Router()
const authRoutes = require('./auth')
const usersRoutes = require('./users')
const taskRoutes = require('./task')

router.use('/auth', authRoutes)
router.use('/users', usersRoutes)
router.use('/task', taskRoutes)

module.exports = router