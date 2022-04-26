const { DataTypes } = require('sequelize');
const now = require('moment');
const bcrypt = require('bcryptjs');
const { sequelize } = require('./configDB');
const logger = require('../config/logger');

const users = sequelize.define('users', {
  id: {
    type: DataTypes.INTEGER,
    autoIncrement: true,
    allowNull: false,
    primaryKey: true,
  },
  email: {
    type: DataTypes.TEXT,
    allowNull: false,
    unique: {
      args: true,
      msg: 'Msg from db: email address already in use!',
    },
    validate: {
      isEmail: { msg: 'Must be a valid email address' },
    },
  },
  username: {
    type: DataTypes.TEXT,
    allowNull: true,
  },
  limit_daily_task: {
    type: DataTypes.INTEGER,
    allowNull: true,
  },
  password: {
    type: DataTypes.TEXT,
    allowNull: false,
  },
  contact: {
    type: DataTypes.TEXT,
    allowNull: true,
    unique: {
      args: true,
      msg: 'Msg from db: contact already in use!',
    },
    validate: {
      isMobilePhone: { msg: 'Must be a valid number phone' },
    },
  },
  avatar: {
    type: DataTypes.TEXT,
    allowNull: true,
    defaultValue: 'http://cdn.onlinewebfonts.com/svg/img_339542.png',
  },
  is_email_verified: {
    type: DataTypes.BOOLEAN,
    allowNull: true,
    defaultValue: false,
  },
  is_contact_verified: {
    type: DataTypes.BOOLEAN,
    allowNull: true,
    defaultValue: false,
  },

  createdAt: DataTypes.DATE(now()),
  updatedAt: DataTypes.DATE(now()),
});

users.beforeCreate(async (user, options) => {
  user.password = await bcrypt.hash(user.password, 8);
});

users.beforeUpdate(async (user, options) => {
  if (user.changed('password')) {
    user.password = await bcrypt.hash(user.password, 8);
  }
  if (user.changed('email')) {
    user.is_email_verified = false;
  }
});

users.isEmailTaken = async (_email) => {
  const user = await users.findOne({ where: { email: _email } });
  return !!user;
};

users.isContactTaken = async (_contact) => {
  const user = await users.findOne({ where: { contact: _contact } });
  return !!user;
};

users.prototype.checkPassword = async (password, truePassword) => {
  return bcrypt.compare(password, truePassword);
};

(async function () {
  // await sequelize.sync({ alter: true }).then(() => { // alter to edit DB after run server
  await sequelize.sync().then(() => {
    logger.info('Sync users Table success!');
  });
})().catch((error) => {
  logger.error('Sync users Table fail');
  logger.error(error);
});

module.exports = users;
