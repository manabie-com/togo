const mongoose = require('mongoose');
const todolist = require('./routes/todolist');
const express = require('express');
const app = express();

mongoose.connect('mongodb://localhost/togo')
  .then(() => console.log('Connected to MongoDB...'))
  .catch(err => console.error('Could not connect to MongoDB...'));

app.use(express.json());
app.use('/api/todolist', todolist);

const port = process.env.PORT || 3000;
app.listen(port, () => console.log(`Listening on port ${port}...`));