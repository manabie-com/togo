const TodoRepository = require('../repositories/todo.repository');

exports.postTodoCreate = (req, res, next) => {
    TodoRepository.postTodoCreate(req, res, next);
}

exports.getTodos = (req, res, next) => {
    TodoRepository.getTodos(req, res, next);
}

