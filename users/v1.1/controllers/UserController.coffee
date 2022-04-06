_          = require 'lodash'
fibrous    = require 'fibrous'
md5        = require 'md5'
code       = require "#{process.cwd()}/config/code"
config     = require "#{process.cwd()}/config/development"
Models     = require "#{process.cwd()}/models"
UsersModel = Models.UsersModel

class UserController
  validateUserSignUpInfo: (userName, email, password, callback ) ->
    try
      fibrous.run () ->
        userInfo = UsersModel.Collection.sync.find {
          $or: [ { userName }, { email }]
        }
        if userInfo?.length > 0 then isValid = false
        else isValid = true
      , callback
      
    catch error
      console.error "validateUserSignUpInfo:error", error

  signUp: (req, res, callback) =>
    try
      { userName, email, password, dailyTaskLimit } = req.body

      fibrous.run () =>
        validateResult = @validateUserSignUpInfo.sync userName, email, password
        if not validateResult then return {
          code: code.CODE_ERROR
          message: "USER_NAME OR EMAIL HAS BEEN USED"
          data: {}
        }

        newUser = UsersModel.sync.create { userName, email, dailyTaskLimit, password: md5 password }
        return {
          code: code.CODE_SUCCESS
          entries: newUser || []
          meta: {
            total: 1
          }
        }
      , (err, rs) ->
        if err? then return callback err
        return res.json rs
    catch error
      console.error "UserController::signUp::error", error
      return callback error

  getAllUsers: (req, res, callback) ->
    try

      fibrous.run () ->
        users = UsersModel.Collection.sync.find()

        users = _.map users, (user) ->
          user.userId = user._id
          user = _.pick user, ['userId', 'userName', 'email', 'dailyTaskLimit']
          user

        return {
          code: code.CODE_SUCCESS
          entries: users || []
          meta: {
            total: users?.length || 0
          }
        }
      , (err, rs) ->
        if err? then return callback err
        return res.json rs
    catch error
      console.error "UserController::getAllUsers::error", error
      return callback error

module.exports = UserController
