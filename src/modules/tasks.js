const TaskService = require('../domain/task');
const resResponse = require('../utils/handle-response');
const httpStatus = require('http-status');
const resError = require('../utils/rest-error');

class User {
  async getTasks(req, res, next) {
    try {
      const tasks = await TaskService.getTasks(req.user.id);
      const response = resResponse(httpStatus.OK, true, tasks);

      res.send(response);
    } catch (error) {
      next(resError(httpStatus.BAD_REQUEST, error));
    }
  }

  async createTask(req, res, next) {
    try {
      const { body, user } = req;

      await TaskService.addTask(body, user.id)

      const response = resResponse(httpStatus.OK, true, true);

      res.send(response);
    } catch (error) {
      next(resError(httpStatus.BAD_REQUEST, error));
    }
  }
}

module.exports = new User();
