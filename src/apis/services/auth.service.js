const httpStatus = require("http-status");
const userService = require("./user.service");

const ApiError = require("../../utils/api-error");

/**
 * Login
 * @param {string} email
 * @param {string} password
 * @returns {Promise<User>}
 */
const login = async (email, password) => {
  const user = await userService.getUserByEmail(email);
  if (!user || !(await user.isPasswordMatch(password))) {
    throw new ApiError(httpStatus.UNAUTHORIZED, "Incorrect email or password");
  }

  return user;
};

module.exports = {
  login,
};
