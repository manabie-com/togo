const auth = require('../middleware/auth');
const admin = require('../middleware/admin');
const { Todo, validate } = require('../models/todo');
const mongoose = require('mongoose');
const express = require('express');
const router = express.Router();

router.get('/', async (req, res) => {
  const todoli = await Todo.find().sort('name');
  res.send(todolist);
});

router.post('/', auth, async (req, res) => {
    const { error } = validate(req.body); 
  if (error) return res.status(400).send(error.details[0].message);

  let todo = new Todo({ name: req.body.name });
  todo = await todo.save();
  
  res.send(todo);
});

router.put('/:id', async (req, res) => {
  const { error } = validate(req.body); 
  
  if (error) return res.status(400).send(error.details[0].message);

  const todo = await Todo.findByIdAndUpdate(req.params.id, { name: req.body.name }, {
    new: true
  });

  if (!todo) return res.status(404).send('The todo with the given ID was not found.');
  
  res.send(todo);
});

router.delete('/:id', [auth, admin], async (req, res) => {
    const todo = await Todo.findByIdAndRemove(req.params.id);

  if (!todo) return res.status(404).send('The todo with the given ID was not found.');

  res.send(todo);
});

router.get('/:id', async (req, res) => {
  const todo = await Todo.findById(req.params.id);

  if (!todo) return res.status(404).send('The todo with the given ID was not found.');

  res.send(todo);
});

module.exports = router;