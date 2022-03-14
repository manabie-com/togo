const env = process.env.NODE_ENV || 'dev';
const rabbitHost = process.env.RABBITMQ_HOST || 'rabbit';
const rabbitUsername = process.env.RABBITMQ_USER || 'guest';
const rabbitPassword = process.env.RABBITMQ_PASS || 'guest';

// eslint-disable-next-line max-len
const mongodbHost = process.env.MONGODB_HOST || '127.0.0.1';

const mysqlHost = process.env.MYSQL_HOST || '127.0.0.1';
const mysqlPort = process.env.MYSQL_PORT || '3306';
const mysqlUsername = process.env.MYSQL_USER || 'root';
const mysqlPassword = process.env.MYSQL_PASSWORD || 'admin123';
const mysqlDatabase = process.env.MYSQL_DATABASE || 'manabie';

environments = {
  test: {
    amqp: {
      host: `amqp://${rabbitUsername}:${rabbitPassword}@${rabbitHost}`,
      port: '5672',
    },
    mongo: {
      host: `mongodb://${mongodbHost}`,
      port: '27017',
      name: 'manabie',
    },
    mysql: {
      host: mysqlHost,
      port: mysqlPort,
      database: mysqlDatabase,
      username: mysqlUsername,
      password: mysqlPassword,
      dialect: 'mysql',
    },
  },
  dev: {
    amqp: {
      host: `amqp://${rabbitUsername}:${rabbitPassword}@${rabbitHost}`,
      port: '5672',
    },
    mongo: {
      host: `mongodb://${mongodbHost}`,
      port: '27017',
      name: 'manabie',
    },
    mysql: {
      host: mysqlHost,
      port: mysqlPort,
      database: mysqlDatabase,
      username: mysqlUsername,
      password: mysqlPassword,
      dialect: 'mysql',
    },
  },
  staging: {
    amqp: {
      host: `amqp://${rabbitUsername}:${rabbitPassword}@${rabbitHost}`,
      port: '5672',
    },
    mongo: {
      host: `mongodb://${mongodbHost}`,
      port: '27017',
      name: 'notification-staging',
    },
    mysql: {
      host: mysqlHost,
      port: mysqlPort,
      database: mysqlDatabase,
      username: mysqlUsername,
      password: mysqlPassword,
      dialect: 'mysql',
    },
  },
  production: {
    amqp: {
      host: `amqp://${rabbitUsername}:${rabbitPassword}@${rabbitHost}`,
      port: '5672',
    },
    mongo: {
      host: `mongodb://${mongodbHost}`,
      port: '27017',
      name: 'notification',
    },
    mysql: {
      host: mysqlHost,
      port: mysqlPort,
      database: mysqlDatabase,
      username: mysqlUsername,
      password: mysqlPassword,
      dialect: 'mysql',
    },
  },
};

module.exports = environments[env];
module.exports.getENV = (key) => {
  return process.env[key] || '';
};
