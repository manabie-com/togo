export const Config = {
  server: {
    port: parseInt(process.env.PORT, 10),
    env: process.env.NODE_ENV,
  },
  database: {
    dbConnection: 'DATABASE_CONNECTION',
    host: process.env.DB_HOST,
    port: parseInt(process.env.DB_PORT, 10),
    username: process.env.DB_USERNAME,
    password: process.env.DB_PASSWORD,
    databaseName: process.env.DB_NAME,
  }
};
