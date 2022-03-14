const jwt = require('jsonwebtoken');
const config = require('../configs');
const jwtSecret = config.getENV('JWT_SECRET') || 'a_secret';
const expiresIn = config.getENV('JWT_EXPIRES_IN') || '1d';
const tokenType = 'Bearer';

module.exports = class Authentication {
  /**
   * constructor
  */
  constructor() {
    this._jwtSecret = jwtSecret;
    this._signOptions = {
      expiresIn: expiresIn,
    };
  }

  /**
   * generate token
   * @param {object} payload
   * @return {string}
   */
  generateToken(payload) {
    const token = jwt.sign(payload, this._jwtSecret, this._signOptions);
    return token;
  }


  /**
   * verify token from header
   * @param {object} req
   * @return {boolean}
   */
  verifyTokenFromHeader(req) {
    let token = req.headers.authorization;
    if (!token || !token.startsWith(`${tokenType} `)) {
      return false;
    }
    token = token.slice(`${tokenType} `.length, token.length).trimLeft();
    if (this.verifyToken(req, token)) {
      return true;
    }
    return false;
  }

  /**
   * verify token
   * @param {object} req
   * @param {string} token
   * @return {boolean}
   */
  verifyToken(req, token) {
    try {
      const payload = jwt.verify(token, this._jwtSecret);
      req.user = {
        id: payload.id,
        name: payload.name,
        email: payload.email,
        role: payload.role,
      };
      return true;
    } catch (err) {
      return false;
    }
  }
};
