const server = require('./app/app');
const config = require('./config/config');
const logger = require('./config/logger');
const { sequelize } = require('./models/configDB');

sequelize
  .authenticate()
  .then(() => {
    logger.info('Connect DB success!!!');
    server.listen(process.env.PORT || config.port, () => {
      logger.info(`Listening to port ${config.port}`);
    });
  })
  // .then(function(){
  //   return sequelize.drop() // drop all tables in the db
  // }).then(() => logger.info('Drop success!'))
  .catch(() => logger.error(`Connect DB fail ${process.env.PORT}`));

const exitHandler = () => {
  if (server) {
    server.close(() => {
      logger.info('Server closed');
      process.exit(1);
    });
  } else {
    process.exit(1);
  }
};

const unexpectedErrorHandler = (error) => {
  logger.error(error);
  exitHandler();
};

process.on('uncaughtException', unexpectedErrorHandler);
process.on('unhandledRejection', unexpectedErrorHandler);

process.on('SIGTERM', () => {
  logger.info('SIGTERM received');
  if (server) {
    server.close();
  }
});
