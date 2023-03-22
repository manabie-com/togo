const { StatusCodes } = require('http-status-codes')
const validator = require('./validates/validator')
const { taskParams } = require('./validates/schemas/task')
const taskService = require('../services/task')
const moment = require('moment-timezone');
const messageCode = require('../common/message-code')
const { dtoAddTask } = require('../dto/task')

const add = (req, res) => {
  const { name } = req.body
  const localDate = req.requestUser.timezone ? moment().tz(req.requestUser.timezone).format('YYYYMMDD') : moment().format('YYYYMMDD')
  validator.validate(req.body, taskParams)
    .then(() => {
      return taskService.getTasksInDate(req.requestUser.userId, localDate)
    })
    .then((tasksInDate) => {
      if (tasksInDate.length < req.requestUser.tasksPerDay) {
        return taskService.save(req.requestUser.userId ,name, localDate)
      } else {
        res.status(StatusCodes.BAD_REQUEST).send(messageCode.responseMessage(messageCode.E002))
      }
    })
    .then((task) => {
      res.status(StatusCodes.OK).send(dtoAddTask(task))
    })
    .catch((error) => {
      res.status(StatusCodes.BAD_REQUEST).send({ message: error })
    })
}

module.exports = {
  add
}
