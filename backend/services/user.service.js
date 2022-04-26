const httpStatus = require('http-status');
const { users } = require('../models');
const ApiError = require('../utils/ApiError');

/**
 * Create an user
 * @param {Object} userBody
 * @returns {Promise<User>}
 */
const createUser = async (userBody) => {
  if (await users.isEmailTaken(userBody.email)) {
    throw new ApiError(httpStatus.BAD_REQUEST, 'Email address already in use!');
  }
  return users.create({ ...userBody });
};

/**
 * Get user by pk
 */
const getUserByPk = async (id) => {
  return users.findByPk(id);
};

/**
 * Update user by pk
 * @param {number} userId
 * @param {Object} updateBody
 * @returns {Promise<users>}
 */
const updateUserByPk = async (userId, updateBody) => {
  const user = await getUserByPk(userId);
  if (!user) {
    throw new ApiError(httpStatus.NOT_FOUND, 'User not found');
  }
  if (updateBody.email && (await users.isEmailTaken(updateBody.email))) {
    throw new ApiError(httpStatus.BAD_REQUEST, 'Email address already in use!');
  }
  if (updateBody.contact && (await users.isContactTaken(updateBody.contact))) {
    throw new ApiError(httpStatus.BAD_REQUEST, 'Contact already in use!');
  }
  Object.assign(user, updateBody);
  await user.save();
  return user;
};

/**
 * Delete user by Pk
 */
const deleteUserByPk = async (userId) => {
  const user = await getUserByPk(userId);
  if (!user) {
    throw new ApiError(httpStatus.NOT_FOUND, 'User not found');
  }
  await user.destroy();
  return user;
};

/**
 * Update password by pk
 * @param {ObjectId} userId
 * @param {Object} body
 * @returns {Promise<User>}
 */
const changePasswordByPk = async (userId, body) => {
  const user = await getUserByPk(userId);
  if (!user) {
    throw new ApiError(httpStatus.NOT_FOUND, 'User not found');
  }
  const isOldPasswordCorrect = await user.checkPassword(body.oldPassword, user.password);
  if (!isOldPasswordCorrect) {
    throw new ApiError(httpStatus.BAD_REQUEST, 'Old password is not correct');
  }
  user.password = body.newPassword;
  await user.save();
  return user;
};

module.exports = {
  getUserByPk,
  updateUserByPk,
  deleteUserByPk,
  changePasswordByPk,
  createUser,
};
