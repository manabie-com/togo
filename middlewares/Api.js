const {validationResult} = require('express-validator');

const BaseMiddleware = require('./BaseMiddleware');
const {Authentication} = require('../libs');
const {errorFormatter} = require('../helpers/validation');

module.exports = class Api extends BaseMiddleware {
  /**
   * Represents a Api middleware.
   * @constructor
   * @param {object} req - this is request.
   * @param {object} res - this is response the request
   * @param {object} next - this is next of the request
   */
  constructor(req, res, next) {
    super(req, res, next);
    this._auth = new Authentication;
  }

  /**
   * check token from header
   * @return {function}
   */
  auth() {
    if (this._auth.verifyTokenFromHeader(this._req)) {
      this._next();
    } else {
      return this.errorWithMessage(401, 'Unauthorized.');
    }
  }

  /**
   * validate request
   */
  validate() {
    const errors = validationResult(this._req).formatWith(errorFormatter);
    if (errors.isEmpty()) {
      this._next();
    } else {
      this.badRequest(errors.array());
    }
  }
};
