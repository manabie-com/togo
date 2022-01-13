const mongoose = require("mongoose");

const Schema = mongoose.Schema;

const taskSchema = new Schema({
  task_name: { type: String, required: true },
  user_name: { type: String, required: true },
  date_created: { type: String, required: true },
});

module.exports = mongoose.model("Task", taskSchema);
