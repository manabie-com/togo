const express = require('express');
const todolist = require('../routes/todolist');
const users = require('../routes/users');
const auth = require('../routes/auth');
const error = require('../middleware/error');

module.exports = function(app) {
  app.use(express.json());
  app.use('/api/todolist', todolist);
  app.use('/api/users', users);
  app.use('/api/auth', auth);
  app.use(error);
}