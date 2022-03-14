const createError = require('http-errors');
const express = require('express');
const path = require('path');
const cookieParser = require('cookie-parser');
const useragent = require('express-useragent');
const {api} = require('./routes');
const cors = require('cors');

const app = express();

app.use(useragent.express());
app.use(express.json());
app.use(express.urlencoded({extended: true}));
app.use(cookieParser());
app.use(express.static(path.join(__dirname, 'public')));
app.use(cors());
// router
app.use('/', api);

// catch 404 and forward to error handler
app.use(function(req, res, next) {
  const error = new Error(`Page not found: ${req.originalUrl}`);
  next(createError(404, error));
});

// error handler
app.use(function(error, req, res, next) {
  const status = error.status || 500;
  const message = error.message;

  const info = {
    status: status,
    error: error.stack,
  };

  res.status(status);
  res.send({message: message});
});

module.exports = app;
