import 'dotenv/config';

export const config = {
  dbHost: process.env.DB_HOST,
  dbPort: +process.env.DB_PORT,
  dbUsername: process.env.DB_USERNAME,
  dbPassword: process.env.DB_PASSWORD,
  dbDatabase: process.env.DB_DATABASE,
};
