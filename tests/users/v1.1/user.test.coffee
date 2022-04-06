config                  = require  "#{process.cwd()}/config/development"
config.prefixCollection = 'mnb_test'
config.prefixModel      = 'mnbTest'

chai                    = require 'chai'
chaiHttp                = require 'chai-http'
code                    = require "#{process.cwd()}/config/code"
server                  = require "#{process.cwd()}/app"
should                  = chai.should()
Models                  = require "#{process.cwd()}/models"
TasksModel              = Models.TasksModel
UsersModel              = Models.UsersModel


chai.use(chaiHttp)
{ expect } = chai

USER_INFO = {
  userName: 'member11113'
  email: 'member11113@gmail.com'
  password: 's324@hahj!$'
  dailyTaskLimit: 2
}


describe('User', () ->
  before (done) ->
    # Empty the database
    UsersModel.Collection.deleteMany {}, (error, result) ->
      if error then throw error
    done()

  # POST
  describe '/POST /users/v1.1/sign-up', () ->
    # it 'it should return err when userName already exists', (done) ->
    it '[pattern validate] it should be blocked when missing param inputs', (done) ->
      chai.request(server)
      .post '/users/v1.1/sign-up'
      .send {}
      .end (err, res) ->
        apiResponse = res.body

        res.should.have.status(200)
        apiResponse.should.have.property('code').eql(406)
        apiResponse.should.have.property('message').eql("data must have required property 'userName', data must have required property 'email', data must have required property 'password'")
      done()

    it '[pattern validate] it should be blocked when userName is number', (done) ->
      chai.request(server)
      .post '/users/v1.1/sign-up'
      .send {
        userName: 11111
        email: USER_INFO.email
        password: USER_INFO.password
        dailyTaskLimit: USER_INFO.dailyTaskLimit
      }
      .end (err, res) ->
        apiResponse = res.body

        res.should.have.status(200)
        apiResponse.should.have.property('code').eql(500)
        apiResponse.should.have.property('message').eql("data/userName must be string")
      done()
    
    it '[pattern validate] it should be blocked when email has less than minimum (6)', (done) ->
      chai.request(server)
      .post '/users/v1.1/sign-up'
      .send {
        userName: USER_INFO.userName
        email: 'email'
        password: USER_INFO.password
        dailyTaskLimit: USER_INFO.dailyTaskLimit
      }
      .end (err, res) ->
        apiResponse = res.body

        res.should.have.status(200)
        apiResponse.should.have.property('code').eql(500)
        apiResponse.should.have.property('message').eql("data/email must NOT have fewer than 6 characters")
      done()
    
    it '[pattern validate] it should be blocked when email has more than maximum (100)', (done) ->
      chai.request(server)
      .post '/users/v1.1/sign-up'
      .send {
        userName: USER_INFO.userName
        email: "data must have required property 'userName', data must have required property 'email', data must have required property 'password'"
        password: USER_INFO.password
        dailyTaskLimit: USER_INFO.dailyTaskLimit
      }
      .end (err, res) ->
        apiResponse = res.body

        res.should.have.status(200)
        apiResponse.should.have.property('code').eql(500)
        apiResponse.should.have.property('message').eql("data/email must NOT have more than 100 characters")
      done()
    
    it 'it should be success - create new user ', (done) ->
      chai.request(server)
      .post '/users/v1.1/sign-up'
      .send USER_INFO
      .end (err, res) ->
        apiResponse = res.body
        res.should.have.status(200)
        apiResponse.should.have.property('code').eql(200)
        apiResponse.entries.should.have.property('_id')
        apiResponse.entries.should.have.property('userName')
        apiResponse.entries.should.have.property('email')
        apiResponse.entries.should.have.property('password')
        apiResponse.entries.should.have.property('dailyTaskLimit')
      done()

    it 'it should be bloked when create new user with userName already exists ', (done) ->
      setTimeout () ->
        chai.request(server)
        .post '/users/v1.1/sign-up'
        .send USER_INFO
        .end (err, res) ->
          apiResponse = res.body
          res.should.have.status(200)
          apiResponse.should.have.property('code').eql(500)
          apiResponse.should.have.property('message').eql("USER_NAME OR EMAIL HAS BEEN USED")
      , 2000
      done()

  # GET
  describe '/GET /users/v1.1/get-all', () ->
    setTimeout () ->
      it 'it should GET all the users', (done) ->
        chai.request(server)
        .get '/users/v1.1/get-all'
        .end (err, res) ->
          apiResponse = res.body
          totalUser = apiResponse?.entries?.length || 0

          res.should.have.status(200)
          apiResponse.should.have.property('code').eql(200)
          apiResponse?.meta?.should.have.property('total').eql(totalUser)
          apiResponse?.meta?.should.have.property('total').to.be.at.least(1)
        done()
    , 2000
  
)
