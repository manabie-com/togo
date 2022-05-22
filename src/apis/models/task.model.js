const mongoose = require("mongoose");

const { priorityTypes } = require("../../configs/priority");

const taskSchema = mongoose.Schema({
  title: {
    type: String,
    required: true,
  },
  description: {
    type: String,
    required: false,
  },
  priority: {
    type: String,
    enum: [priorityTypes.HIGH, priorityTypes.MEDIUM, priorityTypes.LOW],
    required: false,
    default: priorityTypes.MEDIUM,
  },
  completed: {
    type: Boolean,
    required: false,
    default: false,
  },
  createdBy: {
    type: mongoose.Schema.Types.ObjectId,
    required: false,
  },
});

const Task = mongoose.model("Task", taskSchema);

module.exports = Task;
