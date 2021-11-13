const bcrypt = require('bcrypt');
const db = require('../../infrastructure/models');
const { generateToken } = require('../../infrastructure/jwt');
const UserRepository = require('../../infrastructure/repositories/users');
const config = require('../../config/constants');

const { saltRoundPassword } = config;

class AuthenticateService {
  encryptPassword(password) {
    if (!password) {
      return false;
    }

    return bcrypt.hashSync(password, saltRoundPassword);
  }

  comparePassword(newPass, pass) {
    if (!newPass || !pass) {
      return false;
    }

    return bcrypt.compareSync(newPass, pass);
  }

  async register(username, password) {
    if (!username || !password) {
      throw new Error('Username or password is missing.');
    }

    const user = await UserRepository.getUser({
      where: { username: `${username}` },
      attributes: ['id'],
    });

    if (user) {
      throw new Error('User already exists.');
    }

    const t = await db.sequelize.transaction();

    try {
      let newUser = await UserRepository.createUser(
        {
          username,
          password: this.encryptPassword(password),
        },
        { transaction: t },
      );

      newUser = newUser && newUser.toJSON();

      const token = await this.getAuthTokens(newUser);

      await t.commit();

      return { user: newUser, ...token };
    } catch (error) {
      await t.rollback();

      throw new Error('Sign up fail !');
    }
  }

  async login(username, password) {
    if (!username || !password) {
      throw new Error('Username or password is incorrect.');
    }

    let user = await UserRepository.getUser({
      where: { username: `${username}` },
      attributes: ['id', 'max_todo', 'username', 'password'],
    });

    user = user && user.toJSON();
    
    if (!user || !this.comparePassword(password, user.password)) {
      throw new Error('Username or password is incorrect.');
    }

    const token = await this.getAuthTokens(user);

    return { user, ...token };
  }

  async getAuthTokens(user) {
    delete user.password;

    const accessToken = await generateToken(user);

    return {
      accessToken: accessToken,
    };
  }
}

module.exports = new AuthenticateService();
