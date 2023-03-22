const express = require('express')
const bodyParser = require('body-parser')
const cors = require('cors')
const constant = require('./app/common/constant')
const baseServive = require('./app/services/base')

const log4js = require('log4js')
log4js.configure({
  appenders: { server: { type: 'file', filename: 'log/file.log' } },
  categories: { default: { appenders: ['server'], level: 'ALL' } }
})

const logger = log4js.getLogger('server')
logger.info('Start server...')

const app = express()

const corsOptions = {
  origin: 'http://localhost:8080'
}

app.use(cors(corsOptions))

// parse requests of content-type - application/json
app.use(bodyParser.json())

// parse requests of content-type - application/x-www-form-urlencoded
app.use(bodyParser.urlencoded({ extended: true }))

// log all request to log file
app.use(function (req, res, next) {
  logger.info(`Recived a request with method ${req.method} to ${req.originalUrl}`)
  next()
})

// db configuration
const db = require('./app/models')

// confige public forder
app.use(express.static('public'))

// add router
require('./app/routes/account')(app)
require('./app/routes/task')(app)

// ignore certicate
process.env.NODE_TLS_REJECT_UNAUTHORIZED = 0

// set port, listen for requests
const PORT = process.env.PORT || 8080
app.listen(PORT, () => {
  console.log(`Server is running on port ${PORT}.`)
})

let userData = {
  name: 'test',
  email: 'test',
  tasksPerDay: 1,
  timezone: 'Asia/Ho_Chi_Minh',
}

let userData1 = {
  name: 'test',
  email: 'test',
  tasksPerDay: 2,
  timezone: 'Asia/Taipei',
}

let userData2 = {
  name: 'test',
  email: 'test',
  tasksPerDay: 2,
  timezone: 'Asia/Taipei',
}

const accountData = {
  userName: 'free',
  passWord: 'free'
}

const accountData1 = {
  userName: 'vip',
  passWord: 'vip'
}

const accountData2 = {
  userName: 'custom',
  passWord: 'custom'
}

db.sequelize.sync({ force: true })
  .then(() => {
    return baseServive.insert(db.user, userData)
  })
  .then((data) => {
    userData = data
    return baseServive.insert(db.user, userData1)
  })
  .then((data) => {
    userData1 = data
    return baseServive.insert(db.user, userData2)
  })
  .then((data) => {
    userData2 = data
    accountData.userId = userData.userId
    accountData1.userId = userData1.userId
    accountData2.userId = userData2.userId
  })
  .then(() => {
    return baseServive.bulkInsert(db.account, [accountData, accountData1, accountData2])
  })

module.exports = app
