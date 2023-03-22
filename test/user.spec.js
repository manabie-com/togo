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

const BASE_URL_USER = '/api/accounts'
const BASE_URL_TASK = '/api/tasks'

describe('Test all APIs account', () => {
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
        done()
      })
  })
  describe('Request login', () => {
    describe('Without parameter', () => {
      it('should be return error params', done => {
        chai.request(app).post(`${BASE_URL_USER}/login`)
          .end(function (req, res) {
            expect(res.statusCode).to.equal(400)
            expect(res.body.message).to.equal('child "username" fails because ["username" is required]')
            done()
          })
      })
    })
    describe('With invalid parameters', () => {
      it('should be return error params when missing username', done => {
        chai.request(app).post(`${BASE_URL_USER}/login`).send({ password: 'abc' })
          .end(function (req, res) {
            expect(res.statusCode).to.equal(400)
            expect(res.body.message).to.equal('child "username" fails because ["username" is required]')
            done()
          })
      })
      it('should be return error params when username is not a string', done => {
        chai.request(app).post(`${BASE_URL_USER}/login`).send({ username: 1, password: 'abc' })
          .end(function (req, res) {
            expect(res.statusCode).to.equal(400)
            expect(res.body.message).to.equal('child "username" fails because ["username" must be a string]')
            done()
          })
      })
      it('should be return error params when username less than 3 characters', done => {
        chai.request(app).post(`${BASE_URL_USER}/login`).send({ username: '1', password: 'abc' })
          .end(function (req, res) {
            expect(res.statusCode).to.equal(400)
            expect(res.body.message).to.equal('child "username" fails because ["username" length must be at least 3 characters long]')
            done()
          })
      })
      it('should be return error params when username more than 50 characters', done => {
        chai.request(app).post(`${BASE_URL_USER}/login`).send({ username: '11asdasdajsgdasbdvasdais7dyasdgyi7asdvias7daysdasdbhayudgua7dgayidaiudgiauwgdiauwgdiawdiaw8dgia7gd32gv2fyv2j3fv2jh3vfj23fv23fv23fvjh23bfjh2f3vk23fgu2fg2b3k23bfk2fb2k3bfk233ufu2vb3fku2v3fk2vfky2v3f23vfjh2v3fj2vfyj2vj3fjv23fascyasfdcia', password: 'abc' })
          .end(function (req, res) {
            expect(res.statusCode).to.equal(400)
            expect(res.body.message).to.equal('child "username" fails because ["username" length must be less than or equal to 50 characters long]')
            done()
          })
      })
      it('should be return error params when missing password', done => {
        chai.request(app).post(`${BASE_URL_USER}/login`).send({ username: 'abc' })
          .end(function (req, res) {
            expect(res.statusCode).to.equal(400)
            expect(res.body.message).to.equal('child "password" fails because ["password" is required]')
            done()
          })
      })
      it('should be return error params when password is not a string', done => {
        chai.request(app).post(`${BASE_URL_USER}/login`).send({ username: 'abc', password: 1 })
          .end(function (req, res) {
            expect(res.statusCode).to.equal(400)
            expect(res.body.message).to.equal('child "password" fails because ["password" must be a string]')
            done()
          })
      })
    })
    describe('With valid parameters', () => {
      it('should be return error when wrong username', done => {
        chai.request(app).post(`${BASE_URL_USER}/login`).send({ username: 'abc', password: 'test' })
          .end(function (req, res) {
            expect(res.statusCode).to.equal(401)
            expect(res.body.message).to.equal('Username or password incorrect')
            done()
          })
      })
      it('should be return error when wrong username', done => {
        chai.request(app).post(`${BASE_URL_USER}/login`).send({ username: 'test', password: 'abc' })
          .end(function (req, res) {
            expect(res.statusCode).to.equal(401)
            expect(res.body.message).to.equal('Username or password incorrect')
            done()
          })
      })
      let accessToken;
      it('should be return token when correct username and password', done => {
        chai.request(app).post(`${BASE_URL_USER}/login`).send({ username: 'test', password: 'test' })
          .end(function (req, res) {
            expect(res.statusCode).to.equal(200)
            expect(res.body).to.have.property('accessToken')
            accessToken =  res.body.accessToken
            expect(res.body).to.have.property('refreshToken')
            expect(res.body).to.have.property('userName')
            done()
          })
      })
      it('should be return valid token', done => {
        chai.request(app).post(BASE_URL_TASK).set('Authorization', 'bearer ' + accessToken)
          .end(function (req, res) {
            expect(res.statusCode).to.not.equal(401)
            expect(res.statusCode).to.not.equal(403)
            done()
          })
      })
    })
  })
  describe('Request refresh token', () => {
    let accessToken;
    let refreshToken;
    before('Login user', done => {
      chai.request(app).post(`${BASE_URL_USER}/login`).send({ username: 'test', password: 'test' })
        .end(function (req, res) {
          expect(res.statusCode).to.equal(200)
          expect(res.body).to.have.property('accessToken')
          expect(res.body).to.have.property('refreshToken')
          expect(res.body).to.have.property('userName')
          accessToken = res.body.accessToken
          refreshToken = res.body.refreshToken
          done()
        })
    })
    describe('Without parameter', () => {
      it('should be return error params', done => {
        chai.request(app).post(`${BASE_URL_USER}/refreshtoken`)
          .end(function (req, res) {
            expect(res.statusCode).to.equal(400)
            expect(res.body.message).to.equal('child "token" fails because ["token" is required]')
            done()
          })
      })
    })
    describe('Without invalid parameters', () => {
      it('should be return error params', done => {
        chai.request(app).post(`${BASE_URL_USER}/refreshtoken`).send({ token: 'abc' })
          .end(function (req, res) {
            expect(res.statusCode).to.equal(400)
            expect(res.body.message).to.equal('child "token" fails because ["token" length must be at least 10 characters long]')
            done()
          })
      })
    })
    describe('With valid parameters', () => {
      it('should be return error params when token not existed', done => {

        chai.request(app).post(`${BASE_URL_USER}/refreshtoken`).send({ token: refreshToken + '1' })
          .end(function (req, res) {
            expect(res.statusCode).to.equal(403)
            done()
          })
      })
      it('should be return new token', done => {
        chai.request(app).post(`${BASE_URL_USER}/refreshtoken`).send({ token: refreshToken })
          .end(function (req, res) {
            expect(res.statusCode).to.equal(200)
            expect(res.body).to.have.property('accessToken')
            accessToken = res.body.accessToken
            done()
          })
      })
      it('should be return valid token', done => {
        chai.request(app).post(BASE_URL_TASK).set('Authorization', 'bearer ' + accessToken)
          .end(function (req, res) {
            expect(res.statusCode).to.not.equal(401)
            expect(res.statusCode).to.not.equal(403)
            done()
          })
      })
    })
  })
})
