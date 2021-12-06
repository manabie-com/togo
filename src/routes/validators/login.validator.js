const { check } = require("express-validator");

module.exports = [
  check("userName")
    .exists()
    .withMessage("Please provide username")
    .notEmpty()
    .withMessage("Please do not leave the username blank"),

  check("password")
    .exists()
    .withMessage("Please provide password")
    .notEmpty()
    .withMessage("Please do not leave the password blank")
    .isLength({ min: 6 })
    .withMessage("Password must be at least 6 characters"),
];
