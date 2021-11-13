const TaskRepository = require('../../infrastructure/repositories/tasks');

class User {
  async getTasks(userId) {
    try {
      return TaskRepository.getList({ where: { user_id: userId } });
    } catch (error) {
      throw new Error(error);
    }
  }

  async addTask(payload, userId) {
    try {
      if (!payload.content) {
        throw new Error('Content is missing.');
      }

      const isAvailableCapacity = await TaskRepository.isAvailableCapacity(userId);

      if (!isAvailableCapacity) {
        throw new Error('User is reached capacity.');
      }

      return TaskRepository.createTask({ ...payload, user_id: userId });
    } catch (error) {
      throw new Error(error);
    }
  }
}

module.exports = new User();
