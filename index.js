const Joi = require('joi');
const express = require('express');
const app = express();

app.use(express.json());

const todolist = [
  { id: 1, name: 'Action' },  
  { id: 2, name: 'Horror' },  
  { id: 3, name: 'Romance' },  
];

app.get('/api/todolist', (req, res) => {
  res.send(todolist);
});

app.post('/api/todolist', (req, res) => {
  const { error } = validateGenre(req.body); 
  if (error) return res.status(400).send(error.details[0].message);

  const todo = {
    id: todolist.length + 1,
    name: req.body.name
  };
  todolist.push(todo);
  res.send(todo);
});

app.put('/api/todolist/:id', (req, res) => {
  const todo = todolist.find(c => c.id === parseInt(req.params.id));
  if (!todo) return res.status(404).send('The todo with the given ID was not found.');

  const { error } = validateGenre(req.body); 
  if (error) return res.status(400).send(error.details[0].message);
  
  todo.name = req.body.name; 
  res.send(todo);
});

app.delete('/api/todolist/:id', (req, res) => {
  const todo = todolist.find(c => c.id === parseInt(req.params.id));
  if (!todo) return res.status(404).send('The todo with the given ID was not found.');

  const index = todolist.indexOf(todo);
  todolist.splice(index, 1);

  res.send(todo);
});

app.get('/api/todolist/:id', (req, res) => {
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

const port = process.env.PORT || 3000;
app.listen(port, () => console.log(`Listening on port ${port}...`));