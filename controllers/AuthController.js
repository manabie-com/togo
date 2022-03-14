const BaseController = require('./BaseController.js');
const {users} = require('../models').mysql;
const Hash = require('../libs/Hash');
const { DataConstant, Message } = require('../constants');
const validator = require('../libs/Validator');
const {Authentication} = require('../libs');


module.exports = class AuthController extends BaseController {
  /**
   * Constructor.
   *
   * @param {object} req
   * @param {object} res
   * @param {object} next
  */
  constructor(req, res, next) {
    super(req, res, next);
    // libs
    this._hash = new Hash;
    this._validator = validator;
    this._auth = new Authentication;

    // models
    this._user = users;
  }

  /**
   * POST /auth/login
   */
  async login() {
    try {
      const params = this._req.body;
      const {email, password} = params;

      const user = await this._user.findOne({ where: { email } });

      if (!user) {
        return this.errorWithMessage(401, DataConstant.user_not_exit);
      }

      if (!user.password) {
        return this.errorWithMessage(401, DataConstant.user_is_invalid);
      }

      if (!this._hash.check(password, user.password)) {
        return this.errorWithMessage(401, DataConstant.password_is_invalid);
      }

      const userData = {
        id: user.id,
        email: user.email,
        role: user.role
      };

      return this.renderJson({token: this._auth.generateToken(userData)});
    } catch (e) {
      return this.errorWithMessage(500, e);
    }
  }
}
