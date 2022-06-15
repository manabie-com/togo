
import * as dotenv from 'dotenv';
import { IDatabaseConfig } from './database.interface';

dotenv.config();
export const databaseConfig: IDatabaseConfig = {
  development: {
      username: process.env.DB_USER,
      password: process.env.DB_PASS,
      database: process.env.DB_NAME,
      host: process.env.IS_DOCKER,
      port: process.env.DB_PORT,
      dialect: process.env.DB_DIALECT,
  },
  production: {
      username: process.env.DB_USER_PROD,
      password: process.env.DB_PASS_PROD,
      database: process.env.DB_NAME,
      host: process.env.DB_HOST_PROD,
      port: process.env.DB_PORT,
      dialect: process.env.DB_DIALECT,
  },
};