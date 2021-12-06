const { check } = require("express-validator");

module.exports = [
  check("content")
    .exists()
    .withMessage("Please provide content")
    .notEmpty()
    .withMessage("Please do not leave the content blank"),

];
