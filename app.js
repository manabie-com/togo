const express = require('express');
const app = express();
const bodyParser = require('body-parser');

const taskRoutes = require('./routes/tasks');
const userRoutes = require('./routes/users');

app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());

app.use('/tasks', taskRoutes);
app.use('/users', userRoutes);

// Open app on port 3000
const port = 3000;
app.listen(port, () => {
    console.log(`Listening on port ${port}`);
})

// For testing
module.exports = app;