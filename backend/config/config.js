const dotenv = require('dotenv');
const path = require('path');
const Joi = require('joi');

dotenv.config({ path: path.join(__dirname, '../../.env') });

const envVarsSchema = Joi.object()
  .keys({
    NODE_ENV: Joi.string().valid('production', 'development', 'test').required(),
    PORT: Joi.number().default(3000),
    SQL_URL: Joi.string().required().description('SQL DB URL'),
    POSTGRESQL_HOSTNAME: Joi.string().required().description('POSTGRESQL host'),
    POSTGRESQL_USERNAME: Joi.string().required().description('POSTGRESQL username'),
    POSTGRESQL_PASSWORD: Joi.string().required().description('POSTGRESQL pass'),
    POSTGRESQL_PORT: Joi.string().required().description('POSTGRESQL port'),
    POSTGRESQL_DB: Joi.string().required().description('POSTGRESQL Database'),
    JWT_SECRET: Joi.string().required().description('JWT secret key'),
    JWT_ACCESS_EXPIRATION_MINUTES: Joi.number().default(30).description('minutes after which access tokens expire'),
    JWT_REFRESH_EXPIRATION_DAYS: Joi.number().default(30).description('days after which refresh tokens expire'),
    JWT_RESET_PASSWORD_EXPIRATION_MINUTES: Joi.number()
      .default(10)
      .description('minutes after which reset password token expires'),
    JWT_VERIFY_EMAIL_EXPIRATION_MINUTES: Joi.number()
      .default(10)
      .description('minutes after which verify email token expires'),
    JWT_VERIFY_PHONE_EXPIRATION_MINUTES: Joi.number()
      .default(1)
      .description('minutes after which verify phone token expires'),
    SMTP_HOST: Joi.string().description('server that will send the emails'),
    SMTP_PORT: Joi.number().description('port to connect to the email server'),
    SMTP_USERNAME: Joi.string().description('username for email server'),
    SMTP_PASSWORD: Joi.string().description('password for email server'),
    EMAIL_FROM: Joi.string().description('the from field in the emails sent by the app'),
    AWS_ACCESS_KEY_ID: Joi.string().description('access key id for AWS Server'),
    AWS_SECRET_ACCESS_KEY: Joi.string().description('secret key id for AWS Server'),
    REGION: Joi.string().description('region for AWS Server'),
    BUCKET: Joi.string().description('bucket for in AWS for this app'),
    REDIS_PORT: Joi.number().description('redis port'),
    REDIS_HOST: Joi.string().description('redis host'),
    TIME_TO_LIVE: Joi.number().description('redis time to live'),
  })
  .unknown();

const { value: envVars, error } = envVarsSchema.prefs({ errors: { label: 'key' } }).validate(process.env);

if (error) {
  throw new Error(`Config validation error: ${error.message}`);
}

module.exports = {
  env: envVars.NODE_ENV,
  port: envVars.PORT,
  poolPostGre: {
    user: envVars.POSTGRESQL_USERNAME,
    host: envVars.POSTGRESQL_HOSTNAME,
    database: envVars.POSTGRESQL_DB,
    password: envVars.POSTGRESQL_PASSWORD,
    port: envVars.POSTGRESQL_PORT,
    // ssl: envVars.POSTGRESQL_SSL,
  },
  sql: {
    url: envVars.SQL_URL + (envVars.NODE_ENV === 'test' ? '-test' : ''),
    options: {
      useCreateIndex: true,
      useNewUrlParser: true,
      useUnifiedTopology: true,
    },
  },
  jwt: {
    secret: envVars.JWT_SECRET,
    accessExpirationMinutes: envVars.JWT_ACCESS_EXPIRATION_MINUTES,
    refreshExpirationDays: envVars.JWT_REFRESH_EXPIRATION_DAYS,
    resetPasswordExpirationMinutes: envVars.JWT_RESET_PASSWORD_EXPIRATION_MINUTES,
    verifyEmailExpirationMinutes: envVars.JWT_VERIFY_EMAIL_EXPIRATION_MINUTES,
    verifyPhoneExpirationMinutes: envVars.JWT_VERIFY_PHONE_EXPIRATION_MINUTES,
  },
  email: {
    smtp: {
      host: envVars.SMTP_HOST,
      port: envVars.SMTP_PORT,
      auth: {
        user: envVars.SMTP_USERNAME,
        pass: envVars.SMTP_PASSWORD,
      },
    },
    from: envVars.EMAIL_FROM,
  },
  aws: {
    accessKey: envVars.AWS_ACCESS_KEY_ID,
    secretKey: envVars.AWS_SECRET_ACCESS_KEY,
    region: envVars.REGION,
    bucket: envVars.BUCKET,
  },
  redis: {
    redisPort: envVars.REDIS_PORT,
    redisHost: envVars.REDIS_HOST,
    timeToLive: envVars.TIME_TO_LIVE,
  },
};
