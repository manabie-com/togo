const Joi = require('joi');
const { now } = require('lodash');
const mongoose = require('mongoose');

const todoSchema = new mongoose.Schema({
  name: {
    type: String,
    required: true,
    minlength: 5,
    maxlength: 50
  },
  user_id: {
    type: String,
    required: true,
    minlength: 5,
    maxlength: 255
  },
  content: {
    type: String,
  },
  created_date: {
    type: Date,
    default: Date.now
  }
});

const Todo = mongoose.model('Todo', todoSchema);

function validateTodo(todo) {
  const schema = {
    name: Joi.string().min(5).max(50).required(),
    user_id: Joi.string().min(5).max(255).required()
  };

  return Joi.validate(todo, schema);
}

exports.todoSchema = todoSchema;
exports.Todo = Todo; 
exports.validate = validateTodo;