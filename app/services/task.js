'use strict'

const models = require('../models/index')
const baseService = require('./base')

const save = (userId ,name, localDate) => {
  const task = {
    userId,
    name,
    localDate
  }
  return baseService.insert(models.task, task)
}

const getTasksInDate = (userId, date) => {
  const options = {
    attributes: ['taskId'],
    where: {
      localDate: date,
      userId
    }
  }
  return baseService.findAllByOptions(models.task, options)
}

module.exports = {
  save,
  getTasksInDate,
}
