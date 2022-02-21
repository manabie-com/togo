const express = require('express');
const app = express();
const bodyParser = require('body-parser');

const task = require('./routes/task');

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({
    extended: true
}));

app.get('/', function (req, res) {
    return res.send({ message: 'hello' })
});

app.use('/tasks', task);

app.listen(3000, function () {
    console.log('Node app is running on port 3000');
});

module.exports = app;