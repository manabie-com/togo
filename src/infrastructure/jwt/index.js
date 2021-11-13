const jwt = require('jsonwebtoken');
const config = require('../../config/constants');

const jwtConfig = config.jwt;

const verifyToken = ({ token, publicKey }) => {
  const secretKey = jwtConfig.accessSecretKey;

  return jwt.verify(token, publicKey || secretKey);
};

const generateToken = (payload, expires) => {
  let expiresIn = null;
  let secretKey = null;

  expiresIn = expires || jwtConfig.accessTokenExpiration;
  secretKey = jwtConfig.accessSecretKey;

  return jwt.sign(payload, secretKey, { expiresIn });
};

module.exports = { verifyToken, generateToken };