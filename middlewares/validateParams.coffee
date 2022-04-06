_         = require 'lodash'
config    = require 'config'
code      = require "#{process.cwd()}/config/code"
Validator = require "#{process.cwd()}/utils/validation/Validator.coffee"

module.exports = (req, res, next) ->
  try
    validator = new Validator req

    if _.keys(validator).length > 0

      result = validator.validateParams()
      
      if result.valid
        return next()
      else
        response = {
          code: result.code
          message: result.reason
          codeMessage: result.codeMessage
        }

        return res.json response

    next()

  catch error
    console.error "Middlewares::validateParams::error: ", error
    return next error
