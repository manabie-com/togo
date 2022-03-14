const BaseController = require('./BaseController.js');
const {users, tasks} = require('../models').mysql;
const { Authentication } = require('../libs');
const { DataConstant } = require('../constants');
const Message = require('../constants/Messages');

module.exports = class TaskController extends BaseController {
    /**
     * constructor
     */
    constructor(req, res, next) {
      super(req, res, next);
      this._user = users;
      this._task = tasks;
    }

    /**
     * Get List Users
     * GET /users
     */
     async index() {
      const page = this._req.query.page || 1;
      const limit = 20;

      try {
        const users = await this._task.findAll({
          limit,
          offset: (page - 1) * limit
        });

        return this.renderJson(users);
      } catch (e) {
        return this.errorWithMessage(500, e);
      }
    }

    /**
     * Create User
     * POST /user
     */
    async store() {
      try {
        const params = this._req.body;

        const task = await this._task.create(params);

        return this.renderJson(task);
      } catch (e) {
        return this.errorWithMessage(500, e);
      }
    }

    /**
     * Get Detail User
     * GET /user/:id
     */
    async show() {
      try {
        const id = this._req.params.id;
        const task = await this._task.findOne({ where: { id } });

        if (!task) {
          return this.errorWithMessage(404, DataConstant.task_not_exit);
        }

        return this.renderJson(task);
      } catch (e) {
        return this.errorWithMessage(500, e);
      }
    }

    /**
     * Update User
     * PUT /user/"id"
     */
    async update() {
      try {
        const id = this._req.params.id;
        const params = this._req.body;
        const task = await this._task.findOne({ where: { id } });

        if (!task) {
          return this.errorWithMessage(404, DataConstant.task_not_exit);
        }

        await task.update(params);

        const update = await this._task.findOne({ where: { id } });

        return this.renderJson(update);
      } catch (e) {
        return this.errorWithMessage(500, e);
      }
    }

}