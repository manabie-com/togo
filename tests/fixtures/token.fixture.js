const moment = require('moment');
const config = require('../../backend/config/config');
const { tokenTypes } = require('../../backend/config/tokens');
const tokenService = require('../../backend/services/token.service');
const { userOne } = require('./user.fixture');

const accessTokenExpires = moment().add(config.jwt.accessExpirationMinutes, 'minutes');
const userOneAccessToken = tokenService.generateToken(userOne._id, accessTokenExpires, tokenTypes.ACCESS);

module.exports = {
  userOneAccessToken,
};
