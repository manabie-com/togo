const UserService = require('../domain/user');
const resResponse = require('../utils/handle-response');
const httpStatus = require('http-status');
const resError = require('../utils/rest-error');

class User {
  async getUserInfo(req, res, next) {
    try {
      const response = resResponse(httpStatus.OK, true, req.user);

      res.send(response);
    } catch (error) {
      next(resError(httpStatus.BAD_REQUEST, error));
    }
  }

  async updateUserInfo(req, res, next) {
    try {
      const { body, user } = req;

      await UserService.updateUserInfo(body, user);

      const response = resResponse(httpStatus.OK, true, true);

      res.send(response);
    } catch (error) {
      next(resError(httpStatus.BAD_REQUEST, error));
    }
  }
}

module.exports = new User();
