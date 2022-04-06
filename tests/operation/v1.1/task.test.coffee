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
UserTasksModel          = Models.UserTasksModel



chai.use(chaiHttp)
{ expect } = chai

TASK_INFO = {
  taskId: null
  taskName: "Task 1"
  taskCode: "task_1"
  taskDescription: "Check in"
  status: "in-process"
}

describe('Task', () ->

  before (done) ->
    # Empty the database
    TasksModel.Collection.deleteMany {}, (error, result) ->
      if error then throw error
    UserTasksModel.Collection.deleteMany {}, (error, result) ->
      if error then throw error
    done()

  describe '/POST /operation/v1.1/tasks/assign', () ->
    
      # it 'it should return err when userName already exists', (done) ->
      it '[pattern validate] it should be blocked when missing param inputs', (done) ->
        chai.request(server)
        .post '/operation/v1.1/tasks/assign'
        .send {}
        .end (err, res) ->
          apiResponse = res.body
          res.should.have.status(200)
          apiResponse.should.have.property('code').eql(406)
          apiResponse.should.have.property('message').eql("data must have required property 'userId', data must have required property 'taskId', data must have required property 'taskName', data must have required property 'taskDescription', data must have required property 'status', data must have required property 'taskCode'")
        done()

      it '[pattern validate] it should be blocked when taskId is number', (done) ->
        chai.request(server)
        .post '/operation/v1.1/tasks/assign'
        .send {
          taskId: 111111
          userId: null
          taskName: TASK_INFO.taskName
          taskCode: TASK_INFO.taskCode
          taskDescription: TASK_INFO.taskDescription
          status: TASK_INFO.status
        }
        .end (err, res) ->
          apiResponse = res.body
          res.should.have.status(200)
          apiResponse.should.have.property('code').eql(500)
          apiResponse.should.have.property('message').eql("data/userId must be string, data/taskId must be string,null")
        done()
      
      it '[pattern validate] it should be blocked when missing userId', (done) ->
        chai.request(server)
        .post '/operation/v1.1/tasks/assign'
        .send {
          taskId: null
          taskName: TASK_INFO.taskName
          taskCode: TASK_INFO.taskCode
          taskDescription: TASK_INFO.taskDescription
          status: TASK_INFO.status
        }
        .end (err, res) ->
          apiResponse = res.body
          res.should.have.status(200)
          apiResponse.should.have.property('code').eql(406)
          apiResponse.should.have.property('message').eql("data must have required property 'userId'")
        done()
      
      it 'Get User info to assign task', (done) ->
        setTimeout () ->
          chai.request(server)
          .get '/users/v1.1/get-all'
          .end (err, res) ->
            apiResponse = res.body
            TASK_INFO.userId = apiResponse?.entries?[0]?.userId
            res.should.have.status(200)
            apiResponse.should.have.property('code').eql(200)
            apiResponse?.meta?.should.have.property('total').to.be.at.least(1)
        , 2000
        done()
      
      it 'Create first task and assign to user', (done) ->
        setTimeout () ->
          chai.request(server)
          .post '/operation/v1.1/tasks/assign'
          .send TASK_INFO
          .end (err, res) ->
            apiResponse = res.body
            res.should.have.status(200)
            apiResponse.should.have.property('code').eql(200)
        , 2500
        done()
      
      it 'Create second task and assign to user', (done) ->
        setTimeout () ->
          chai.request(server)
          .post '/operation/v1.1/tasks/assign'
          .send {
            taskId: null
            userId: TASK_INFO.userId
            taskName: 'Task 2'
            taskCode: 'task_2'
            taskDescription: "Check out"
            status: 'done'
          }
          .end (err, res) ->
            apiResponse = res.body
            res.should.have.status(200)
            apiResponse.should.have.property('code').eql(200)
        , 2500
        done()
      
      it 'It should be bloked when create third task with dailyTaskLimit(2)', (done) ->
        setTimeout () ->
          chai.request(server)
          .post '/operation/v1.1/tasks/assign'
          .send {
            taskId: null
            userId: TASK_INFO.userId
            taskName: 'Task 3'
            taskCode: 'task_3'
            taskDescription: "Payment"
            status: 'done'
          }
          .end (err, res) ->
            apiResponse = res.body
            res.should.have.status(200)
            apiResponse.should.have.property('code').eql(500)
            apiResponse.should.have.property('message').eql('Task limit exceeded')
        , 2700
        done()

  # GET
  describe '/GET /operation/v1.1/tasks', () ->
    it 'it should GET all tasks', (done) ->
      setTimeout () ->
        chai.request(server)
        .get '/operation/v1.1/tasks'
        .end (err, res) ->
          apiResponse = res.body
          total = apiResponse?.entries?.length || 0

          res.should.have.status(200)
          apiResponse.should.have.property('code').eql(200)
          apiResponse?.meta?.should.have.property('total').eql(total)
          apiResponse?.meta?.should.have.property('total').to.be.at.least(1)
      , 2700
      done()
  
  describe '/GET USER_TASK /operation/v1.1/tasks/by-user', () ->
    it 'it should GET users and tasks assigned', (done) ->
      setTimeout () ->
        chai.request(server)
        .get '/operation/v1.1/tasks/by-user'
        .end (err, res) ->
          apiResponse = res.body
          total = apiResponse?.entries?.length || 0

          res.should.have.status(200)
          apiResponse.should.have.property('code').eql(200)
          apiResponse?.meta?.should.have.property('total').eql(total)
          apiResponse?.meta?.should.have.property('total').to.be.at.least(1)
      , 2700
      done()
      
)

