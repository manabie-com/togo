'use strict'

const models = require('../models/index')

const insert = (model, data, transaction = null) => {
  return model.create(data, {
    transaction
  })
    .then(insertResult => {
      return insertResult
    })
}

const bulkInsert = (model, arrayData) => {
  return model.bulkCreate(arrayData)
    .then(insertResult => {
      return insertResult
    })
    .catch(error => {
      throw error
    })
}

const findAllByOptions = (model, options) => {
  return model
    .findAll(options).then(records => {
      return records
    })
}


const getOneByOptions = (model, options) => {
  return model
    .findOne(options).then(records => {
      return records
    })
}


module.exports = {
  insert,
  bulkInsert,
  findAllByOptions,
  getOneByOptions
}
