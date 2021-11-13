const passport = require('passport');
const httpStatus = require('http-status');
const resError = require('../utils/rest-error');
const config = require('../config/constants');

const { authService } = config;

class Authenticate {
  constructor(req, res, next) {
    this.req = req;
    this.res = res;
    this.next = next;
  }

  authorize() {
    passport.authenticate(authService.jwt, { session: false }, (error, userInfo) =>
      this.checkExist(userInfo, error),
    )(this.req, this.res, this.next);
  }

  checkExist(userInfo, error) {
    try {
      const hasError = error || !userInfo;

      if (hasError) {
        throw resError(httpStatus.UNAUTHORIZED, error || 'Unauthorized');
      }

      this.req.user = userInfo;

      this.next();
    } catch (e) {
      this.next(e);
    }
  }
}

module.exports = (method) => async (req, res, next) => {
  try {
    const authenticate = new Authenticate(req, res, next);

    await authenticate[method]();
  } catch (error) {
    next(error);
  }
};
