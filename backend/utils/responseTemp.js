const responseTemp = (httpStatus, mess, data = null) => {
  return {
    status: httpStatus,
    message: mess,
    data,
  };
};

module.exports = responseTemp;
