const httpStatus = require('http-status');

module.exports = (status, success, data, error, count) => {
  const response = {
    code: status || httpStatus.OK,
    success,
    message: error ? error.message : null,
    data: data || null,
  };
  
  return response;
};
