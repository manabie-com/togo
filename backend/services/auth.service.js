const httpStatus = require('http-status');
const tokenService = require('./token.service');
const userService = require('./user.service');
const ApiError = require('../utils/ApiError');
const { tokens } = require('../models');

/**
 * Login with email and password
 * @param {string} email
 * @param {string} password
 * @returns {Promise}
 */
const loginEmailAndPassword = async (email, password) => {
  const user = await userService.getUserByEmail(email.toLowerCase());
  if (!user) {
    throw new ApiError(httpStatus.NOT_FOUND, 'No user found with this email');
  }
  const isPasswordCorrect = await user.checkPassword(password, user.password);
  if (!isPasswordCorrect) {
    throw new ApiError(httpStatus.BAD_REQUEST, 'Incorrect email or password');
  }
  return user;
};

/**
 * Logout with refreshToken and deviceId
 * @param {string} refreshToken
 * @returns {Promise}
 */
const logout = async (refreshToken) => {
  const tokenDoc = await tokens.findOne({ where: { token: refreshToken } });
  if (!tokenDoc) {
    throw new ApiError(httpStatus.NOT_FOUND, 'Not found token for logout');
  }
  await tokenDoc.destroy();
};

/**
 * Refresh auth token with refreshToken and deviceId
 * @param {string} refreshToken
 * @returns {Promise}
 */
// eslint-disable-next-line no-shadow
const refreshToken = async (refreshToken) => {
  try {
    const tokenDoc = await tokenService.verifyRefreshTokenSQL(refreshToken);
    const user = await userService.getUserByPk(tokenDoc.user_id);
    if (!user) {
      throw new ApiError(httpStatus.NOT_FOUND, 'User not found with this refresh token');
    }
    await tokenDoc.destroy();
    return tokenService.generateAuthTokens(user);
  } catch (error) {
    throw new ApiError(httpStatus.UNAUTHORIZED, 'Please authenticate');
  }
};

/**
 * Verify email
 */
const verifyEmail = async (verifyEmailToken) => {
  try {
    const verifyEmailTokenDoc = await tokenService.verifyEmailToken(verifyEmailToken);
    const user = await userService.getUserByPk(verifyEmailTokenDoc.user_id);
    if (!user) {
      throw new ApiError(httpStatus.NOT_FOUND, 'User not found');
    }
    await verifyEmailTokenDoc.destroy();
    await userService.updateUserByPk(user.id, { is_email_verified: true });
  } catch (error) {
    throw new ApiError(httpStatus.UNAUTHORIZED, 'Email verification failed');
  }
};

module.exports = {
  loginEmailAndPassword,
  logout,
  refreshToken,
  verifyEmail,
};
