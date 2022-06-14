let mongoose = require('mongoose')
let userSchema = new mongoose.Schema({
  user_name: String,
  email: String,
  limit_task_per_day: Number
})
module.exports = mongoose.model('User', userSchema)