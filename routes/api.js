const express = require('express');
const {body} = require('express-validator');
const controllers = require(`${process.cwd()}/controllers`);
const middlewares = require(`${process.cwd()}/middlewares`);

const router = new express.Router();

router.get('/', function(req, res, next) {
  res.send('Welcome to Manabie Assignment');
});

router.post('/login', controllers('AuthController', 'login'));

router.get('/users', middlewares('Api', ['auth']), controllers('UserController', 'index'));
router.post('/user', middlewares('Api', ['auth']), controllers('UserController', 'store'));
router.get('/user/:id', middlewares('Api', ['auth']), controllers('UserController', 'show'));
router.put('/user/:id', middlewares('Api', ['auth']), controllers('UserController', 'update'));

router.get('/tasks', middlewares('Api', ['auth']), controllers('TaskController', 'index'));
router.post('/task', middlewares('Api', ['auth']), controllers('TaskController', 'store'));
router.get('/task/:id', middlewares('Api', ['auth']), controllers('TaskController', 'show'));
router.put('/task/:id', middlewares('Api', ['auth']), controllers('TaskController', 'update'));

module.exports = router;
