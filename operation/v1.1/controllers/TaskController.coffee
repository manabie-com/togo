_              = require 'lodash'
fibrous        = require 'fibrous'
async          = require 'async'
mongoose       = require 'mongoose'
code           = require "#{process.cwd()}/config/code"
config         = require "#{process.cwd()}/config/development"
Models         = require "#{process.cwd()}/models"
TasksModel     = Models.TasksModel
UsersModel     = Models.UsersModel
UserTasksModel = Models.UserTasksModel

ObjectId = mongoose.Types.ObjectId

buildUserFilters = ( params, options = {}) ->
  try
    taskFilters             = []
    stepFilter = {}
    sortParams = {}
    { limit, page, userId, sort, sd } = params
    
    if userId? then stepFilter = { _id: new ObjectId userId }
    if sort?
      sortParams[sort] = if sd is "asc" then 1 else -1
    else sortParams = { _id: -1 }
    
    limit = +limit || 25
    page  = +page || 1
    skip  = limit * (page - 1)
    
    taskFilters = [
      { $match: stepFilter }
      { $sort: sortParams }
      { $limit: limit }
      { $skip: skip }
      {
        $lookup: {
          "from": "#{config.prefixCollection}_user_tasks"
          "localField": "_id"
          "foreignField": "userId"
          "as": "tasks"
        }
      }
      
    ]
    return taskFilters
  catch e then throw e

buildTaskFilters = ( params, options = {}) ->
  try
    stepFilter = {}
    sortParams = {}
    { limit, page, userId, sort, sd, taskId } = params
    
    if taskId? then stepFilter = { _id: new ObjectId taskId }
    if sort?
      sortParams[sort] = if sd is "asc" then 1 else -1
    else sortParams = { taskId: -1 }
    
    taskFilters = [
      { $match: stepFilter }
      { $sort: sortParams }
    ]

    if limit? and page?
      skip  = +limit * (+page - 1)
      taskFilters = _.concat taskFilters, [
        { $limit: limit }
        { $skip: skip }
      ]

    return taskFilters
  catch e then throw e

normalizationResponseData = (users, tasks, options = {} ) ->
  try
    tasks = _.groupBy tasks, '_id'

    results = _.map users, (user) ->
      user.userId = user._id
      user = _.pick user, ['userId', 'userName', 'email', 'dailyTaskLimit', 'tasks']
      
      userTasks = _.map user?.tasks || [], (assignedTask) ->
        taskInfo = tasks?[assignedTask?.taskId]?[0] || {}
        if taskInfo?._id? then return {
          taskId: assignedTask.taskId
          taskName: taskInfo.taskName
          taskCode: taskInfo.taskCode
          taskDescription: taskInfo.taskDescription
          status: assignedTask.status
        }
        else return null

      userTasks              = _.compact userTasks
      user.totalTaskAssigned = userTasks?.length || 0
      user.tasks             = userTasks
      return user

    return results
  catch e then return []

class TaskController

  validateTaskAssigned: ( userId, taskId, callback) ->
    try
      
      async.parallel {
        users: (done) -> UsersModel.Collection.find { _id: userId }, done
        userTasks: (done) -> UserTasksModel.Collection.find { userId: userId }, done
      }, (err, result) ->
        if err then console.error "TaskController::validateTaskAssigned::err", err;  throw err

        { users, userTasks } = result
        validateResult = {}
        
        if users?.length < 0 then validateResult = { isValid: false, message: 'userId not exists' }
        else
          dailyTaskLimit = _.first(users)?.dailyTaskLimit || 0
          if userTasks?.length >= dailyTaskLimit then validateResult = { isValid: false, message: 'Task limit exceeded', dailyTaskLimit }
          else
            validateResult = { isValid: true, message: 'Already to assign', dailyTaskLimit }

        return callback null, validateResult
  
    catch error then throw error

  getAllTasks: (req, res, callback) ->
    try

      fibrous.run () ->
        tasks = TasksModel.Collection.sync.find()
        tasks = _.map tasks, (task) ->
          task.taskId = task._id
          task = _.pick task, ['taskId', 'taskName', 'taskDescription', 'ctime', 'utime']
          task

        return {
          code: code.CODE_SUCCESS
          entries: tasks || []
          meta: {
            total: tasks?.length || 0
          }
        }
      , (err, rs) ->
        if err? then return callback err
        return res.json rs
    catch error
      console.error "TaskController::getAllTasks::error", error
      return callback error
  
  assignTask: (req, res, callback) =>
    try
      body = req.body
      { userId, taskId, taskCode } = body
      fibrous.run () =>
        result = {
          code: code.CODE_ERROR
          message: 'Something went wrong'
        }

        try
          userId = new ObjectId userId
          taskId = new ObjectId taskId
        catch e
          result.message = 'Invalid type of userId or taskId'
          return result
        
        validateResult = @validateTaskAssigned.sync userId, taskId

        if not validateResult?.isValid
          result.message = validateResult.message
          return result

        newTask = TasksModel.sync.upsert { taskCode: taskCode } , body

        if newTask?._id?
          taskId = newTask?._id
          body.taskId = taskId
          assignTaskToUser = UserTasksModel.sync.upsert { $and: [{ userId }, { taskId }] }, body
          
        else result.message = 'Create task failed'; return result

        if not assignTaskToUser?._id then result.message = 'Assign failed'; return result

        return {
          code: code.CODE_SUCCESS
          message: 'Success'
          entries: assignTaskToUser
        }

      , (err, rs) ->
        if err? then return callback err
        return res.json rs
    catch error
      console.error "TaskController::assignTask::error", error
      return callback error

  getUserTasks: (req, res, callback) ->
    try
      { userId, limit, page, sort, sd } = req.query
      
      fibrous.run () ->

        userFilters = buildUserFilters { userId, limit, page, sort, sd }
        users       = UsersModel.Collection.aggregate(userFilters).sync.exec()

        taskIds = []
        users.forEach (user) ->
          userTaskIds = _.map user?.tasks || [], (task) -> return new ObjectId(task.taskId)
          taskIds = _.concat taskIds, userTaskIds

        taskFilters = buildTaskFilters { taskIds: _.compact(_.uniq taskIds ) }
        tasks       = TasksModel.Collection.aggregate(taskFilters).sync.exec()

        results     = normalizationResponseData users, tasks

        return {
          code: code.CODE_SUCCESS
          entries: results || []
          meta: {
            total: results?.length || 0
          }
        }
      , (err, rs) ->
        if err? then return callback err
        return res.json rs
    catch error
      console.error "TaskController::getUserTasks::error", error
      return callback error

module.exports = TaskController
