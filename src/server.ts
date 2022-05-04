import hapi from '@hapi/hapi';
import Inert from '@hapi/inert';
import Vision from '@hapi/vision';
import {
  HOST,
  PORT,
  MONGO_URI,
  MONGO_USER,
  MONGO_PASS,
  MONGO_DB_NAME
} from './config';

import logger from './logger';
import routes from './routes';

import errorHandler from './plugins/errorHandler.plugin';
import ResponseWrapper from './plugins/responseWrapper.plugin';
import RequestWrapper from './plugins/requestWrapper.plugin';
import Swagger from './plugins/swagger.plugin';
import { connectMongo } from './connectMongo';
import taskConsumer from './task/task.consumer';

const createServer = async () => {
  const server = new hapi.Server({
    port: PORT,
    host: HOST,
    routes: {
      validate: {
        options: {
          abortEarly: false
        },
        failAction: errorHandler
      }
    }
  });
  // Register routes
  server.route(routes);

  const plugins: any[] = [
    Inert,
    Vision,
    Swagger,
    RequestWrapper,
    ResponseWrapper
  ];
  await server.register(plugins);

  return server;
};

export const init = async () => {
  await connectMongo({
    dbUri: MONGO_URI,
    user: MONGO_USER,
    pass: MONGO_PASS,
    dbName: MONGO_DB_NAME
  });

  const server = await createServer();
  await server.initialize().then(() => {
    logger.info(`server started at ${server.info.host}:${server.info.port}`);
    taskConsumer.createTaskConsumer();
  });
  return server;
};

export const start = async (module: NodeModule) => {
  if (!module.parent) {
    logger.info('Start server');
    try {
      const server = await init();
      await server.start();
    } catch (err: unknown) {
      logger.error('Server cannot start', { message: (err as Error).stack });
    }
  }
};

start(module);
