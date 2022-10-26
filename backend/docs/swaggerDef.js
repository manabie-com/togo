const { version } = require('../../package.json');
const config = require('../config/config');

const swaggerDef = {
  openapi: '3.0.0',
  info: {
    title: 'Todos-App API',
    version,
  },
  servers: [
    {
      url: `http://localhost:${config.port}/v1`,
    },
  ],
};

if (process.env.NODE_ENV === 'development' || process.env.NODE_ENV === 'production') {
  swaggerDef.servers[0].url = `http://localhost:${config.port}/v1`;
}

module.exports = swaggerDef;
