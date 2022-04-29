'user strictt'

process.env.NODE_ENV = 'testing'

const chai = require('chai')
const chaiHttp = require('chai-http')
const expect = require('chai').expect

const app = require('../server')
const db = require('../app/models')
const baseServive = require('../app/services/base')
const constant = require('../app/common/constant')

chai.use(chaiHttp)

module.exports = {
  app,
  baseServive,
  chai,
  expect,
  db,
  constant
}
