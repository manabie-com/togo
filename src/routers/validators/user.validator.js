const { check } = require("express-validator");

module.exports = [
  check("email")
    .exists()
    .withMessage("Please provide email")
    .notEmpty()
    .withMessage("Please do not leave the email blank"),

  check("password")
    .exists()
    .withMessage("Please provide password")
    .notEmpty()
    .withMessage("Please do not leave the password blank")
    .isLength({ min: 7 })
    .withMessage("Password must be at least 7 characters"),
];
