const express = require('express')
const router = express.Router()
const bodyParser = require('body-parser')
const jsonParser = bodyParser.json()
const { checkAndGenerate } = require('../validate')

const tasks = []
const users = [
  {
    id: 1,
    name: 'Joy',
    taskLimit: 1,
  },
]

router.get('/tasks', function (req, res) {
  res.send(tasks)
})

router.post('/tasks', jsonParser, function (req, res) {
  const inputTask = checkAndGenerate(req.body.name, req.body.userId)
  if (!inputTask) {
    res.status(400).send({
      message:
        'Name and userId are required. Name should be a string and userId should be a number.',
    })
    return
  }

  const user = users.find((o) => o.id === inputTask.userId)

  if (!user) {
    res.status(400).send({ message: 'User does not exist.' })
    return
  }

  const userTasks = tasks.filter((o) => o.userId === inputTask.userId)

  if (userTasks.length >= user.taskLimit) {
    res.status(400).send({
      message: 'This user already reached a maximum limit of tasks per day.',
    })
    return
  }

  const task = {
    id: tasks.length + 1,
    ...inputTask,
  }
  tasks.push(task)
  res.status(201).send(task)
})

module.exports = router
