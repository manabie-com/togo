const httpStatus = require("http-status");
const _ = require("lodash");

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

/**
 * Get token from Bearer token + signature
 * @param {Object} headers
 * @returns {Promise<string>}
 */
const getTokenFromHeaders = (headers) => {
  const token = _.get(headers, "authorization");

  if (!_.includes(token, "Bearer "))
    throw new ApiError(httpStatus.UNAUTHORIZED, "Please authenticate");

  const parts = _.split(token, ".");
  if (_.size(parts) !== 3) throw new Error("Invalid token. Has no signature");

  return token.replace("Bearer ", "");
};

module.exports = {
  login,
  getTokenFromHeaders,
};
