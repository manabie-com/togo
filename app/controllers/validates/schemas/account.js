'use strict'

const Joi = require('joi')

const loginParams = Joi.object({
  username: Joi.string().required().min(3).max(50),
  password: Joi.string().required(),
})

const refreshTokenParams = Joi.object({
  token: Joi.string().required().min(10),
})

module.exports = {
  loginParams,
  refreshTokenParams
}
