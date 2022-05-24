'use strict';


const Koa = require('koa');
const logger = require('koa-logger')
const bodyParser = require('koa-bodyparser');
const cors = require('koa-cors')
const routes = require('./routes');
const app = new Koa();

const { errorHandler } = require('./middlewares/error-handler');
const { functionContext } = require('./middlewares/function-context');
const { passport } = require('./middlewares/passort');

app
  .use(errorHandler())
    // .use(logger())
  .use(cors())
  .use(functionContext())
  .use(passport())
  .use(bodyParser())
  .use(routes.routes())
  .use(routes.allowedMethods())
  .use((ctx) => {
    return ctx.showError(`Not found API ${ctx.request.method}: ${ctx.request.url}.`, 404);
  });

module.exports = app;
