const mongoose = require('mongoose')
const Schema = mongoose.Schema
const schemaName = 'task'
const ObjectID = mongoose.Types.ObjectId

let taskSchema = new Schema(
    {
        title: { type: String },
        content: { type: String },
        author: { type: ObjectID, ref: 'user' },
        created_at: { type: Date, default: new Date() },
        updated_at: { type: Date },
        deleted_at: { type: Date }
    },
    { strict: false }
)

const Task = mongoose.model(schemaName, taskSchema, schemaName)
module.exports = { Task }