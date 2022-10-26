const jwt = require('jsonwebtoken');
const moment = require('moment');
const httpStatus = require('http-status');
const config = require('../config/config');
const userService = require('./user.service');
const { tokens } = require('../models');
const ApiError = require('../utils/ApiError');
const { tokenTypes } = require('../config/tokens');

/**
 * Generate token
 * @param {ObjectId} userId
 * @param {moment.Moment} expires
 * @param {string} type
 * @param {string} [secret]
 * @returns {string}
 */
const generateToken = (userId, expires, type, secret = config.jwt.secret) => {
  const payload = {
    sub: userId,
    iat: moment().unix(),
    exp: expires.unix(),
    type,
  };
  return jwt.sign(payload, secret);
};

/**
 * Save token into SQL
 * @param {Number} user_id
 * @param {string} token
 * @param {Moment} expires
 * @param {string} type
 * @returns {Promise<tokens>>}
 */
const saveToken = async ( user_id, token, type, expires) => {
  const tokenDoc = await tokens.create({
    user_id,
    token,
    type,
    expires: expires.toDate(),
  });
  return tokenDoc;
};

/**
 * Verify refresh token and return token doc (or throw an error if it is not valid)
 * @param {string} refreshToken
 * @returns {Promise<tokens>}
 */
const verifyRefreshTokenSQL = async (refreshToken) => {
  const payload = jwt.verify(refreshToken, config.jwt.secret);
  const user_id = payload.sub;
  if (!user_id) {
    throw new ApiError(httpStatus.NOT_FOUND, 'user_id not found');
  }
  const tokenDoc = await tokens.findOne({ where: { user_id, token: refreshToken } });
  if (!tokenDoc) {
    throw new ApiError(httpStatus.NOT_FOUND, 'Token not found');
  }
  return tokenDoc;
};

/**
 * Verify email token and return token doc (or throw an error if it is not valid)
 * @param {string} verifyEmailToken
 * @returns {Promise<tokens>}
 */
const verifyEmailToken = async (verifyEmailToken,) => {
  const payload = jwt.verify(verifyEmailToken, config.jwt.secret);
  const user_id = payload.sub;
  if (!user_id) {
    throw new ApiError(httpStatus.NOT_FOUND, 'userID not found');
  }
  const tokenDoc = await tokens.findOne({ where: { user_id, token: verifyEmailToken, type: tokenTypes.VERIFY_EMAIL } });
  if (!tokenDoc) {
    throw new ApiError(httpStatus.NOT_FOUND, 'Token not found');
  }
  return tokenDoc;
};
/**
 * Generate auth tokens and save into DB
 * @param {Object} user
 * @returns {Promise<Object>}
 */
const generateAuthTokens = async (user) => {
  const accessTokenExpires = moment().add(config.jwt.accessExpirationMinutes, 'minutes');
  const accessToken = generateToken(user.id, accessTokenExpires, tokenTypes.ACCESS);

  const refreshTokenExpires = moment().add(config.jwt.refreshExpirationDays, 'days');
  const refreshToken = generateToken(user.id, refreshTokenExpires, tokenTypes.REFRESH);

  const tokenDoc = await saveToken(user.id, refreshToken, tokenTypes.REFRESH, refreshTokenExpires);
  return {
    access: {
      token: accessToken,
      expires: accessTokenExpires.toDate(),
    },
    refresh: {
      token: refreshToken,
      expires: refreshTokenExpires.toDate(),
    },
  };
};

/**
 * Generate verify email token
 * @param {Object} user
 * @returns {Promise<tokens>>}
 */
const generateVerifyEmailToken = async (user) => {
  const expires = moment().add(config.jwt.verifyEmailExpirationMinutes, 'minutes');
  const verifyEmailToken = generateToken(user.id, expires, tokenTypes.VERIFY_EMAIL);
  let tokenDoc = await tokens.findOne({ where: { user_id: user.id, token: verifyEmailToken, type: tokenTypes.VERIFY_EMAIL } });
  if (!tokenDoc){
    //create new Token
    tokenDoc = await saveToken(user.id, verifyEmailToken, tokenTypes.VERIFY_EMAIL, expires);
  } else {
    //update Token
    tokenDoc = await userService.updateEmailTokenByUserId(user.id, verifyEmailToken);
  }
  return tokenDoc;
};

/**
 * Generate phone token
 * @param {ObjectId} userId
 * @param {String} code
 * @param {Moment} expires
 * @param {string} type
 * @param {string} [secret]
 * @returns {string}
 */
const generatePhoneToken = (userId, code, expires, type, secret = config.jwt.secret) => {
  const payload = {
    sub: userId,
    iat: moment().unix(),
    exp: expires.unix(),
    type,
    code,
  };
  return jwt.sign(payload, secret);
};


module.exports = {
  generateToken,
  saveToken,
  verifyRefreshTokenSQL,
  verifyEmailToken,
  generateAuthTokens,
  generateVerifyEmailToken,

};
