mongoose = require 'mongoose'
config   = require  '../config/development.coffee'

Schema = new mongoose.Schema(
  userId:
    type: mongoose.Schema.ObjectId
    ref: "#{config.prefixModel}Users"
  taskId:
    type: mongoose.Schema.ObjectId
    ref: "#{config.prefixModel}Tasks"
  status:
    type: String
  ctime: Date
  utime: Date
)

Schema.pre "save", (next) ->
  if this.isNew is true
    this.ctime = Date.now()

  this.utime = Date.now()
  next()

Schema.pre "update", () ->
  this.update( {}, { $set: { utime: Date.now() } } )

UserTasks = mongoose.model("#{config.prefixModel}UserTasks", Schema, "#{config.prefixCollection}_user_tasks")

class UserTasksModel
  constructor: () ->
    @Collection = UserTasks

  create: (data, callback) =>
    newModel = new @Collection(data)
    newModel.save (err, newModel) ->
      callback(err, newModel)

  upsert: (filters, data, callback) =>
    @Collection
      .findOneAndUpdate filters, data, { upsert: true
      new: true
      setDefaultsOnInsert: true }, callback


  createMany: (query, callback) =>
    @Collection
      .insertMany query, callback

  find: (query, callback) =>
    @Collection
      .find query
      .exec callback

  findOne: (query, callback) =>
    @Collection
      .findOne query
      .exec callback

  update: (query, params, callback) =>
    @Collection
      .update query, params
      .exec callback

  remove: (query, callback) =>
    @Collection
      .remove query
      .exec callback

module.exports = UserTasksModel
