const models = require('../../models');

class TaskRepository {
  constructor() {
    this.model = models.tasks;
  }

  getTasks(options) {
    return this.model.findOne(options);
  }

  getList(options) {
    return this.model.findAll(options);
  }

  async isAvailableCapacity(userId, transaction) {
    const user = await models.users.findOne({
      where: { id: userId },
      attributes: ['max_todo'],
      transaction,
    });
    const totalTask = await this.model.count({ where: { user_id: userId }, transaction });
    
    if (!user) {
      throw new Error('User is not exist.');
    }

    return user.max_todo > totalTask;
  }

  createTask(data, options = {}) {
    return this.model.create(data, options);
  }
}

module.exports = new TaskRepository();
