const { Sequelize } = require('sequelize');
const config = require('../config/config');
const logger = require('../config/logger');

let sequelize;

if (process.env.NODE_ENV === 'production') {
  logger.info('NODE_ENV = production');
  sequelize = new Sequelize(config.poolPostGre.database, config.poolPostGre.user, config.poolPostGre.password, {
    host: config.poolPostGre.host,
    dialect: 'postgres',
    port: config.poolPostGre.port,
    dialectOptions: {},
  });
} else {
  logger.info('NODE_ENV = dev');
  sequelize = new Sequelize(config.sql.url);
}

module.exports = { sequelize };
