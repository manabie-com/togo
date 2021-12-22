const { check } = require("express-validator");

module.exports = [
  check("title")
    .exists()
    .withMessage("Please provide title")
    .notEmpty()
    .withMessage("Please do not leave the title blank"),

    check("description")
    .exists()
    .withMessage("Please provide description")
    .notEmpty()
    .withMessage("Please do not leave the description blank"),
];
