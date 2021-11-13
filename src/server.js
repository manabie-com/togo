const express = require('express');
const cors = require('cors');
const httpStatus = require('http-status');
const passport = require('passport');
const passportStrategy = require('./infrastructure/passport');
const routes = require('./routes/router');
const errorConverter = require('./utils/rest-error');
const handleError = require('./utils/handle-error');
const config = require('./config/constants');
const swaggerUi = require('swagger-ui-express');

const swaggerDocument = require('./swagger.json');

const { authService } = config;

/**
 * Express instance
 * @public
 */
const app = express();

app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use(cors());

// authentiation
app.use(passport.initialize());
passport.use(authService.jwt, passportStrategy.jwtStrategy);

// Swagger
app.use('/api/v1/docs', swaggerUi.serve, swaggerUi.setup(swaggerDocument));

app.use('/api/v1', routes);
app.use((req, res, next) => {
  return next(errorConverter(httpStatus.NOT_FOUND, 'Not found'));
});

app.use(handleError);

module.exports = app;
