const { DataTypes } = require('sequelize');
const now = require('moment');
const { sequelize } = require('./configDB');
const logger = require('../config/logger');
const { users } = require('./index');

const tokens = sequelize.define('tokens', {
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
  token: {
    type: DataTypes.TEXT,
    allowNull: true,
    defaultValue: null,
  },
  type: {
    type: DataTypes.TEXT,
    allowNull: true,
    defaultValue: null,
  },
  expires: {
    type: DataTypes.DATE,
    allowNull: true,
    defaultValue: null,
  },
  createdAt: DataTypes.DATE(now()),
  updatedAt: DataTypes.DATE(now()),
});

users.hasMany(tokens, { foreignKey: 'user_id' });
tokens.belongsTo(users, { foreignKey: 'user_id', targetKey: 'id' });

(async function () {
  // await sequelize.sync({ alter: true }).then(() => { // alter to edit DB after run server
  await sequelize.sync().then(() => {
    logger.info('Sync tokens table success!');
  });
})().catch((error) => {
  logger.error('Sync tokens Table fail');
  logger.error(error);
});

module.exports = tokens;
