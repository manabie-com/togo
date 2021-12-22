const express = require('express')
const auth = require('../middleware/auth')
const {insertTask, getTask, deleteTask, updateTask, getTaskById} = require('../controller/task')
const router = express.Router()
const taskValidate = require('./validators/task.validator')

router.post('/', auth, taskValidate, insertTask)
router.get('/', auth, getTask)
router.get('/:id', auth, getTaskById)
router.delete('/:id', auth, deleteTask)
router.put('/:id', auth, taskValidate, updateTask)

module.exports = router