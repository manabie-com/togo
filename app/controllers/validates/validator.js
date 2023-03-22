'use strict'

const Joi = require('joi')
const Promise = require('bluebird')
const log = require('log4js').getLogger()

const validate = (body, schema) => {
  return new Promise((resolve, reject) => {
    return Joi.validate(body, schema, err => {
      if (!err) {
        return resolve()
      }

      log.error('Validate request information fail!. Error: ', err.message)
      reject(err.message)
    })
  })
}

module.exports = {
  validate
}
