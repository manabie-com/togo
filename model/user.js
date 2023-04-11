const mongoose = require('mongoose')
const Schema = mongoose.Schema
const schemaName = 'user'
const {MAX_NUMBER_TASK_CREATED } = require('../util/constant')

let userSchema = new Schema(
    {
        email: { type: String, unique: true },
        password: { type: String },
        quota: {
            max_post_by_day: { type: Number, default: MAX_NUMBER_TASK_CREATED, min: 0 },
            last_task_created_at: { type: Date },
            remaining_post: { type: Number, default: MAX_NUMBER_TASK_CREATED, min: 0 }
        },
        created_at: { type: Date, default: new Date() },
        updated_at: { type: Date },
        deleted_at: { type: Date }
    },
    { strict: false }
)

const User = mongoose.model(schemaName, userSchema, schemaName)
module.exports = { User }