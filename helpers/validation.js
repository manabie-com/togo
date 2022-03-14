
exports.errorFormatter = ({msg, param}) => {
  return {
    message: msg,
    field: param,
  };
};