require('dotenv').config();

const {
  SERVICE_NAME,
  HOST,
  PORT,
  LOG_LEVEL,
  MONGO_URI,
  MONGO_USER,
  MONGO_PASS,
  MONGO_DB_NAME
} = process.env;

export {
  SERVICE_NAME,
  HOST,
  PORT,
  LOG_LEVEL,
  MONGO_URI,
  MONGO_USER,
  MONGO_PASS,
  MONGO_DB_NAME
};
