const { chai, app, db, expect, baseServive, constant } = require('./common')

let userData = {
  name: 'test',
  email: 'test',
  tasksPerDay: 1,
  timezone: 'Asia/Ho_Chi_Minh',
}

const accountData = {
  userName: 'test',
  passWord: 'test'
}

let userData1 = {
  name: 'test',
  email: 'test',
  tasksPerDay: 1,
}

const accountData1 = {
  userName: 'test1',
  passWord: 'test1'
}

const BASE_URL_USER = '/api/accounts'
const BASE_URL_TASK = '/api/tasks'

describe('Test all APIs task', () => {
  before('Clean data for test', done => {
    db.sequelize.sync({ force: true })
      .then(() => {
        return baseServive.insert(db.user, userData)
      })
      .then((data) => {
        userData = data
        accountData.userId = userData.userId
        return baseServive.insert(db.account, accountData)
      })
      .then(() => {
        return baseServive.insert(db.user, userData1)
      })
      .then((data) => {
        userData1 = data
        accountData1.userId = userData1.userId
        return baseServive.insert(db.account, accountData1)
      })
      .then(() => {
        done()
      })
  })
  describe('Request add task', () => {
    describe('Without logged in account', () => {
      it('Should return error when do not have access token', done => {
        chai.request(app).post(BASE_URL_TASK).send({ name: 'test' })
          .end(function (req, res) {
            expect(res.statusCode).to.equal(401)
            done()
          })
      })
      it('Should return error when invalid token', done => {
        chai.request(app).post(BASE_URL_TASK).set('Authorization', 'bearer abc').send({ name: 'test' })
          .end(function (req, res) {
            expect(res.statusCode).to.equal(403)
            done()
          })
      })
    })
    describe('With logged in account', () => {
      let accessToken;
      let accessToken1;
      before('Login new account to send request', done => {
        chai.request(app).post(`${BASE_URL_USER}/login`).send({ username: 'test', password: 'test' })
          .end(function (req, res) {
            expect(res.statusCode).to.equal(200)
            expect(res.body).to.have.property('accessToken')
            accessToken = res.body.accessToken
            done()
          })
      })
      describe('And without parameters', () => {
        it('Should return error', done => {
          chai.request(app).post(BASE_URL_TASK).set('Authorization', 'bearer ' + accessToken)
            .end(function (req, res) {
              expect(res.statusCode).to.equal(400)
              expect(res.body.message).to.equal('child "name" fails because ["name" is required]')
              done()
            })
        })
      })
      describe('And without invalid parameters', () => {
        it('Should return error when name is not a string', done => {
          chai.request(app).post(BASE_URL_TASK).set('Authorization', 'bearer ' + accessToken).send({ name: 12 })
            .end(function (req, res) {
              expect(res.statusCode).to.equal(400)
              expect(res.body.message).to.equal('child "name" fails because ["name" must be a string]')
              done()
            })
        })
        it('Should return error when name is empty string', done => {
          chai.request(app).post(BASE_URL_TASK).set('Authorization', 'bearer ' + accessToken).send({ name: '' })
            .end(function (req, res) {
              expect(res.statusCode).to.equal(400)
              expect(res.body.message).to.equal('child "name" fails because ["name" is not allowed to be empty]')
              done()
            })
        })
        it('Should return error when name have more than 100 charactors', done => {
          chai.request(app).post(BASE_URL_TASK).set('Authorization', 'bearer ' + accessToken).send({ name: 'adasdgashjdgasdashdasdt6adas6dasd6asd7asdas7d6a7sd6asdas7d6as76das7d6asda6s7dasd6as7das5dd5asd6as5d6as5das6d5as6d5as6d5a6sd5a6sd5a6d5a6sd5a6sd' })
            .end(function (req, res) {
              expect(res.statusCode).to.equal(400)
              expect(res.body.message).to.equal('child "name" fails because ["name" length must be less than or equal to 100 characters long]')
              done()
            })
        })
      })
      describe('And with valid parameters', () => {
        it('Should save task success', done => {
          chai.request(app).post(BASE_URL_TASK).set('Authorization', 'bearer ' + accessToken).send({ name: 'test' })
            .end(function (req, res) {
              expect(res.statusCode).to.equal(200)
              expect(res.body.name).to.equal('test')
              expect(res.body.status).to.equal('NEW')
              expect(res.body).to.have.property('taskId')
              expect(res.body).to.have.property('userId')
              done()
            })
        })
        it('Should return error when task is limited', done => {
          chai.request(app).post(BASE_URL_TASK).set('Authorization', 'bearer ' + accessToken).send({ name: 'test' })
            .end(function (req, res) {
              expect(res.statusCode).to.equal(400)
              expect(res.body.message).to.equal("User's tasks are limited")
              done()
            })
        })
      })
      describe('And another account', ()=>{
        before('Login another account', done=>{
          chai.request(app).post(`${BASE_URL_USER}/login`).send({ username: 'test1', password: 'test1' })
          .end(function (req, res) {
            expect(res.statusCode).to.equal(200)
            expect(res.body).to.have.property('accessToken')
            accessToken1 = res.body.accessToken
            done()
          })
        })
        it('Should save task success', done => {
          chai.request(app).post(BASE_URL_TASK).set('Authorization', 'bearer ' + accessToken1).send({ name: 'test' })
            .end(function (req, res) {
              expect(res.statusCode).to.equal(200)
              expect(res.body.name).to.equal('test')
              expect(res.body.status).to.equal('NEW')
              expect(res.body).to.have.property('taskId')
              expect(res.body).to.have.property('userId')
              done()
            })
        })
      })
    })
  })
})
