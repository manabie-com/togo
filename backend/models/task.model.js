const { DataTypes } = require('sequelize');
const now = require('moment');
const { sequelize } = require('./configDB');
const logger = require('../config/logger');
const { users } = require('./index');

const tasks = sequelize.define('tasks', {
  id: {
    type: DataTypes.INTEGER,
    autoIncrement: true,
    allowNull: false,
    primaryKey: true,
  },
  user_id: {
    type: DataTypes.INTEGER,
    allowNull: false,
  },
  task_name: {
    type: DataTypes.TEXT,
    allowNull: false,
    defaultValue: null,
  },
  task_priority: {
    type: DataTypes.INTEGER,
    allowNull: false,
    defaultValue: 0,
  },
  createdAt: DataTypes.DATE(now()),
  updatedAt: DataTypes.DATE(now()),
});

users.hasMany(tasks, { foreignKey: 'user_id' });
tasks.belongsTo(users, { foreignKey: 'user_id', targetKey: 'id' });

(async function () {
  // await sequelize.sync({ alter: true }).then(() => { // alter to edit DB after run server
  await sequelize.sync().then(() => {
    logger.info('Sync tokens table success!');
  });
})().catch((error) => {
  logger.error('Sync tokens Table fail');
  logger.error(error);
});

module.exports = tasks;
