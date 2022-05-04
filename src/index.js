const express = require('express');
const bodyParser = require('body-parser');
const app = express();
app.use(bodyParser.urlencoded({ extended: true }))
app.use(bodyParser.json())
const port = 3000;
let routes = require('./api/routes') 
routes(app);

app.listen(port)

module.exports = app;