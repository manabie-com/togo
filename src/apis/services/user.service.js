const httpStatus = require("http-status");

const ApiError = require("../../utils/api-error");
const { User } = require("../models");

/**
 * Get user by email
 * @param {string} email
 * @returns {Promise<User>}
 */
const getUserByEmail = async (email) => {
  return User.findOne({ email });
};

/**
 * Create a user
 * @param {Object} userBody
 * @returns {Promise<User>}
 */
const createUser = async (userBody) => {
  if (await User.isEmailTaken(userBody.email)) {
    throw new ApiError(httpStatus.BAD_REQUEST, "Email already taken");
  }

  return User.create(userBody);
};

module.exports = {
  getUserByEmail,
  createUser,
};
