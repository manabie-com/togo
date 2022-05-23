const Joi = require('joi');
const mongoose = require('mongoose');

const todoSchema = new mongoose.Schema({
  name: {
    type: String,
    required: true,
    minlength: 5,
    maxlength: 50
  }
});

const Todo = mongoose.model('Todo', todoSchema);

function validateTodo(todo) {
  const schema = {
    name: Joi.string().min(5).max(50).required()
  };

  return Joi.validate(todo, schema);
}

exports.todoSchema = todoSchema;
exports.Todo = Todo; 
exports.validate = validateTodo;