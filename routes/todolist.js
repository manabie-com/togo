const validateObjectId = require('../middleware/validateObjectId');
const auth = require('../middleware/auth');
const admin = require('../middleware/admin');
const mongoose = require('mongoose');
const { Todo, validate } = require('../models/todo');
const express = require('express');
const router = express.Router();

router.get('/', async (req, res) => {
  const todolist = await Todo.find().sort('name');
  res.send(todolist);
});

router.post('/', auth, async (req, res) => {
  const { error } = validate(req.body); 
  if (error) return res.status(400).send(error.details[0].message);

  let todo = new Todo({ name: req.body.name });
  todo = await todo.save();
  
  res.send(todo);
});

router.put('/:id', [auth, validateObjectId], async (req, res) => {
  const { error } = validate(req.body); 
  
  if (error) return res.status(400).send(error.details[0].message);

  if(!mongoose.Types.ObjectId.isValid(req.params.id))
    return res.status(404).send('Invalid ID.');

  const todo = await Todo.findByIdAndUpdate(req.params.id, { name: req.body.name }, {
    new: true
  });

  if (!todo) return res.status(404).send('The todo with the given ID was not found.');
  
  res.send(todo);
});

router.delete('/:id', [auth, admin, validateObjectId], async (req, res) => {
    const todo = await Todo.findByIdAndRemove(req.params.id);

  if (!todo) return res.status(404).send('The todo with the given ID was not found.');

  res.send(todo);
});

router.get('/:id', validateObjectId, async (req, res) => {
  if(!mongoose.Types.ObjectId.isValid(req.params.id))
    return res.status(404).send('Invalid ID.');

  const todo = await Todo.findById(req.params.id);

  if (!todo) return res.status(404).send('The todo with the given ID was not found.');

  res.send(todo);
});

module.exports = router;