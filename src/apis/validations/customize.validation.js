const mongoose = require("mongoose");

const objectId = (value, helpers) => {
  if (!mongoose.Types.ObjectId.isValid(value)) {
    return helpers.message('{{#label}} must be a valid id format');
  }
  return value;
};

const password = (value, helpers) => {
  if (value.length < 6) {
    return helpers.message("password must be at least 6 characters");
  }
  if (!value.match(/\d/) || !value.match(/[a-zA-Z]/)) {
    return helpers.message(
      "password must contain at least 1 letter and 1 number"
    );
  }
  return value;
};

module.exports = {
  objectId,
  password,
};
