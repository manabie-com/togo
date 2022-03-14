const BaseController = require('./BaseController.js');
const {users} = require('../models').mysql;
const { Authentication } = require('../libs');
const { DataConstant } = require('../constants');
const Message = require('../constants/Messages');
const Hash = require('../libs/Hash');
const validator = require('../libs/Validator');
const Op = require('sequelize').Op;

module.exports = class UserController extends BaseController {
    /**
     * constructor
     */
    constructor(req, res, next) {
      super(req, res, next);
      this._user = users;
      this._auth = new Authentication;
      this._hash = new Hash;
      this._validator = validator;
    }

    /**
     * Get List Users
     * GET /users
     */
    async index() {
      const page = this._req.query.page || 1;
      const limit = 20;

      try {
        const users = await this._user.findAll({
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
        const name = this._req.body.name;
        const email = this._req.body.email;
        const password = this._req.body.password;
        const role = this._req.body.role;
        const status = this._req.body.status;

        const model = {
          name,
          email,
          password: this._hash.make(password),
          role,
          status
        }

        const user = await this._user.create(model);

        return this.renderJson(user);
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
        const user = await this._user.findOne({ where: { id } });

        if (!user) {
          return this.errorWithMessage(404, DataConstant.user_not_exit);
        }

        return this.renderJson(user);
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
        const user = await this._user.findOne({ where: { id } });

        if (!user) {
          return this.errorWithMessage(404, DataConstant.user_not_exit);
        }

        const params = this._req.body;

        if (params['email']) {
          this._validator.validate(params, {'email': 'email'});
  
          const duplicateEmailInfluencer = await this._user.findOne({
            where: {
              [Op.and]: [
                {id: {[Op.ne]: id}},
                {email: params['email']},
              ],
            },
          });
  
          if (duplicateEmailInfluencer) {
            return this.errorWithMessage(422, Message.EMAIL_ALREADY_EXISTS);
          }
        }

        await user.update(params);

        const info = await this._user.findOne({ where: { id } });

        return this.renderJson(info);
      } catch (e) {
        return this.errorWithMessage(500, e);
      }
    }
}