'use strict'

const Joi = require('joi')

const taskParams = Joi.object({
  name: Joi.string().required().min(1).max(100),
})

module.exports = {
  taskParams
}
