const { roles } = require('../config/roles');

const password = (value, helpers) => {
  if (value.length < 8) {
    return helpers.message('password must be at least 8 characters');
  }
  if (!value.match(/\d/) || !value.match(/[a-zA-Z]/)) {
    return helpers.message('password must contain at least 1 letter and 1 number');
  }
  return value;
};

const isRole = (value, helpers) => {
  if (!roles.includes(value)) {
    return helpers.message('Role is invalid, should be user, admin ...');
  }
  return value;
};

module.exports = {
  password,
  isRole,
};
