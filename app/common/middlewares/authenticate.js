const jwt = require('jsonwebtoken')
const { StatusCodes } = require('http-status-codes')

const { getUserById } = require('../../services/user')

const config = require('config')
const accessToken = config.get('accessToken')
const accessTokenSecret = accessToken.accessTokenSecret

const authenticateJWT = (req, res, next) => {
  const authHeader = req.headers.authorization

  if (authHeader) {
    const token = authHeader.split(' ')[1]

    jwt.verify(token, accessTokenSecret, (err, user) => {
      if (err) {
        return res.sendStatus(StatusCodes.FORBIDDEN)
      }

      getUserById(user.userId)
        .then(user => {
          if (user) {
            req.requestUser = user
            next()
          } else {
            res.sendStatus(StatusCodes.NON_AUTHORITATIVE_INFORMATION)
          }
        })
        .catch(() => {
          res.sendStatus(StatusCodes.INTERNAL_SERVER_ERROR)
        })
    })
  } else {
    res.sendStatus(StatusCodes.UNAUTHORIZED)
  }
}

module.exports = {
  authenticateJWT
}
