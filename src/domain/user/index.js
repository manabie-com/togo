const UserRepository = require('../../infrastructure/repositories/users');

class User {
  async updateUserInfo(payload, user) {
    try {
      delete payload.password;
      
      return UserRepository.updateUser(payload, { where: { id: user.id } });
    } catch (error) {
      throw new Error(error);
    }
  }
}

module.exports = new User();
