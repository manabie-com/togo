mongoose = require 'mongoose'
config   = require  '../config/development.coffee'

Schema = new mongoose.Schema(
  userName:
    type: String
  email:
    type: String
  password:
    type: String
  dailyTaskLimit:
    type: Number
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

Users = mongoose.model("#{config.prefixModel}Users", Schema, "#{config.prefixCollection}_users")


class UsersModel
  constructor: () ->
    @Collection = Users

  create: (data, callback) =>
    if not data?.dailyTaskLimit? then data?.dailyTaskLimit = 0
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

module.exports = UsersModel
