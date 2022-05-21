const { User, preSaveFunc } = require("./user.model");
const Task = require("./task.model");

module.exports = {
  User,
  Task,
  preSaveFunc,
};
