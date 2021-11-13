const httpStatus = require('http-status');

module.exports = (err, req, res, next) => {
  const error = err || {};
  const code = error.status || httpStatus.INTERNAL_SERVER_ERROR;
  const response = {
    success: false,
    message: error.message || 'System error. Please try again later !',
    data: null,
  };

  delete response.stack;

  res.status(code).json(response);
};
