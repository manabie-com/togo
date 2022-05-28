require('dotenv').config();

const commonEnv = {
  projectName: process.env.PROJECT_NAME || 'todo-app',
  environment: process.env.NODE_ENV,
  timeZone: 'Asia/Bangkok',

  debug: process.env.DEBUG || true,
  applicationMode: process.env.APPLICATION_MODE || 'api',

  swagger: {
    path: '/_swagger',
  },

  server: {
    host: process.env.HOST || '0.0.0.0',
    domainURL: process.env.DOMAIN_URL || 'http://localhost:3000',
    port: process.env.PORT || 3000,
  },

  jwt: {
    expiration: 24 * 60 * 60 * 7, // 7 days
    rememberMe: 24 * 60 * 60 * 30, // 30 days
    issuer: process.env.HOST || 'localhost',
    audience: 'TodoApp',
    secretKey: process.env.JWT_SECRET_KEY || '2a4Xj9vQjfxV3rw2Usi2qMKEiMlOIjN6',
    refreshTokenTTL: 3600 * 24 * 90, // 90 days
    refreshTokenKeyPrefix: 'refresh_token',
  },
};

export const environment = (() => {
  return commonEnv;
})();
