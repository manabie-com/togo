let mongoose = require('mongoose')
let todoSchema = new mongoose.Schema({
  title: String,
  description: String,
  user_id: String,
  created_at: String
})
module.exports = mongoose.model('Todo', todoSchema)