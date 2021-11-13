const createError = require('http-errors');

module.exports = (status, error) => {
  const errorFormat = {};
  const message = (error && error.message) || error;

  errorFormat.status = status;
  errorFormat.message = message;

  return createError(status, message, errorFormat);
};
