const express = require('express');
const router = express.Router();

const todolist = [
  { id: 1, name: 'Action' },  
  { id: 2, name: 'Horror' },  
  { id: 3, name: 'Romance' },  
];

router.get('/', (req, res) => {
  res.send(todolist);
});

router.post('/', (req, res) => {
  const { error } = validateGenre(req.body); 
  if (error) return res.status(400).send(error.details[0].message);

  const todo = {
    id: todolist.length + 1,
    name: req.body.name
  };
  todolist.push(todo);
  res.send(todo);
});

router.put('/:id', (req, res) => {
  const todo = todolist.find(c => c.id === parseInt(req.params.id));
  if (!todo) return res.status(404).send('The todo with the given ID was not found.');

  const { error } = validateGenre(req.body); 
  if (error) return res.status(400).send(error.details[0].message);
  
  todo.name = req.body.name; 
  res.send(todo);
});

router.delete('/:id', (req, res) => {
  const todo = todolist.find(c => c.id === parseInt(req.params.id));
  if (!todo) return res.status(404).send('The todo with the given ID was not found.');

  const index = todolist.indexOf(todo);
  todolist.splice(index, 1);

  res.send(todo);
});

router.get('/:id', (req, res) => {
  const todo = todolist.find(c => c.id === parseInt(req.params.id));
  if (!todo) return res.status(404).send('The todo with the given ID was not found.');
  res.send(todo);
});

function validateGenre(todo) {
  const schema = {
    name: Joi.string().min(3).required()
  };

  return Joi.validate(todo, schema);
}

module.exports = router;