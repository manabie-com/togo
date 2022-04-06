mongoose = require 'mongoose'
config   = require  '../config/development.coffee'

Schema = new mongoose.Schema(
  taskName:
    type: String
  taskCode:
    type: String
  taskDescription:
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

Tasks = mongoose.model("#{config.prefixModel}Tasks", Schema, "#{config.prefixCollection}_tasks")

class TasksModel
  constructor: () ->
    @Collection = Tasks

  create: (data, callback) =>
    newModel = new @Collection(data)
    newModel.save (err, newModel) ->
      callback(err, newModel)

  createMany: (query, callback) =>
    @Collection
      .insertMany query, callback
  
  upsert: (filters, data, callback) =>
    @Collection
      .findOneAndUpdate filters, data, { upsert: true
      new: true
      setDefaultsOnInsert: true }, callback

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

module.exports = TasksModel
