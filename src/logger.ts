import pino from 'pino';
import { SERVICE_NAME, LOG_LEVEL } from './config';

const logger = pino({
  level: LOG_LEVEL,
  name: SERVICE_NAME
});

export default logger;
