'use strict'

const models = require('../models/index')
const baseService = require('./base')

const getUserByAccount = (userName, passWord) => {
  const options = {
    attributes: ['userId', 'userName'],
    where: {
      userName: userName,
      passWord: passWord,
      isActive: true
    }
  }
  return baseService.getOneByOptions(models.account, options)
}

const getUserById = (userId) => {
  const options = {
    attributes: ['userId', 'name', 'tasksPerDay', 'timezone'],
    where: {
      userId: userId
    }
  }
  return baseService.getOneByOptions(models.user, options)
}

module.exports = {
  getUserByAccount,
  getUserById,
}
