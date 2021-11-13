const models = require('../../models');

class UserRepository {
  constructor() {
    this.model = models.users;
  }

  getUser(options) {
    return this.model.findOne(options);
  }

  createUser(data, options = {}) {
    return this.model.create(data, options);
  }

  updateUser(data, options = {}) {
    return this.model.update(data, options);
  }
}

module.exports = new UserRepository();
