const validateObjectId = require('../middleware/validateObjectId');
const auth = require('../middleware/auth');
const admin = require('../middleware/admin');
const mongoose = require('mongoose');
const { Todo, validate } = require('../models/todo');
const { User } = require('../models/user');
const express = require('express');
const router = express.Router();


router.post('/', [auth], async (req, res) => {
  const { error } = validate(req.body); 
  if (error) return res.status(400).send(error.details[0].message);
  const {name, user_id} = req.body;

  const user = await User.findById(user_id);

  if(!user) return res.status(404).send('Invalid user ID.');

  const todoInDay = await Todo.find({
    "user_id": user_id,"created_date":{$gt:new Date(Date.now() - 24*60*60 * 1000)}
  });
  if(todoInDay.length >= user.max_todo)
    return res.status(400).send('Can not add ToDo because the daily amount of Todo has reached the limit.');

  let todo = new Todo({ name, user_id});
  todo = await todo.save();
  
  res.send(todo);
});

module.exports = router;