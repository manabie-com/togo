const jwt = require('jsonwebtoken')
const { StatusCodes } = require('http-status-codes')

const userService = require('../services/user')
const messageCode = require('../common/message-code')
const validator = require('./validates/validator')
const { loginParams, refreshTokenParams } = require('./validates/schemas/account')

const config = require('config')
const accessToken = config.get('accessToken')
const accessTokenSecret = accessToken.accessTokenSecret
const refreshTokenSecret = accessToken.refreshTokenSecret
let refreshTokens = []

const login = (req, res) => {
  const { username, password } = req.body
  validator.validate(req.body, loginParams)
    .then(() => userService.getUserByAccount(username, password))
    .then((user) => {
      const response = {}
      if (user) {
        response.accessToken = jwt.sign({ userName: user.userName, userId: user.userId }, accessTokenSecret, { expiresIn: '30m' })
        response.refreshToken = jwt.sign({ userName: user.userName, userId: user.userId }, refreshTokenSecret)
        response.userName = user.userName
        refreshTokens.push(response.refreshToken)

        res.send(response)
      } else {
        res.status(StatusCodes.UNAUTHORIZED).send(messageCode.responseMessage(messageCode.E001))
      }
    })
    .catch((error) => {
      res.status(StatusCodes.BAD_REQUEST).send({ message: error })
    })
}

const refreshToken = (req, res) => {
  const { token } = req.body
  validator.validate(req.body, refreshTokenParams).then(() => {
    if (!refreshTokens.includes(token)) {
      res.status(StatusCodes.FORBIDDEN).send(messageCode.responseMessage(messageCode.E003))
    }

    jwt.verify(token, refreshTokenSecret, (err, user) => {
      if (err) {
        res.status(StatusCodes.FORBIDDEN).send(messageCode.responseMessage(messageCode.E003))
      }
      const accessToken = jwt.sign({ userName: user.userName, userId: user.userId }, accessTokenSecret, { expiresIn: '20m' })
      res.status(StatusCodes.OK).send({
        accessToken
      })
    })
  }).catch((error) => {
    res.status(StatusCodes.BAD_REQUEST).send({ message: error })
  })
}

module.exports = {
  login,
  refreshToken
}
