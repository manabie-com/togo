var express = require('express');
var router = express.Router();
const todoController = require('../controller/todo.controller');
var {validate} = require('../app/validator');

router.get('/list', todoController.getTodos);
router.post('/add', validate.validateTodo(), todoController.postTodoCreate);

module.exports = router;
