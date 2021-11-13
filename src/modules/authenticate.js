const httpStatus = require('http-status');
const AuthService = require('../domain/authentication');
const resError = require('../utils/rest-error');
const resResponse = require('../utils/handle-response');

class AuthModule {
  async register(req, res, next) {
    try {
      const {
        body: { username, password },
      } = req;
      const result = await AuthService.register(username, password);
      const response = resResponse(httpStatus.OK, true, result);

      res.send(response);
    } catch (error) {
      next(resError(httpStatus.BAD_REQUEST, error));
    }
  }

  async login(req, res, next) {
    try {
      const {
        body: { username, password },
      } = req;

      const result = await AuthService.login(username, password);
      const response = resResponse(httpStatus.OK, true, result);

      return res.send(response);
    } catch (error) {
      return next(resError(httpStatus.BAD_REQUEST, error));
    }
  }
}

module.exports = new AuthModule();
