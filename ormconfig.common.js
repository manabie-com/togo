require('dotenv').config();

if (
  !(process.env.DB_HOST && process.env.DB_PORT && process.env.DB_USER && process.env.DB_PASSWORD && process.env.DB_NAME)
) {
  const common = require('@nestjs/common');

  common.Logger.error(
    `Database environment values are not configured. Please prepare .env file or export database environment variables.`,
  );
  process.exit(1);
}

module.exports = {
  type: 'mssql',
  host: process.env.DB_HOST,
  port: parseInt(process.env.DB_PORT, 10),
  username: process.env.DB_USER,
  password: process.env.DB_PASSWORD,
  database: process.env.DB_NAME,
  timezone: 'utc',
  synchronize: false,
  charset: 'utf8mb4',
  options: {
    trustedConnection: true,
    enableArithAbort: true,
    trustServerCertificate: true,
  },
  pool: {
    min: 1,
    max: 20,
  },
};
