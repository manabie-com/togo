const Joi = require('joi');
const todolist = require('./routes/todolist');
const express = require('express');
const app = express();

app.use(express.json());
app.use('/api/todolist', todolist);

const port = process.env.PORT || 3000;
app.listen(port, () => console.log(`Listening on port ${port}...`));