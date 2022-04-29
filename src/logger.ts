import pino from 'pino';
import { SERVICE_NAME, LOG_LEVEL } from './config';
import context from './common/context';
import { Tracing } from './common/constants';

const logger = pino({
  level: LOG_LEVEL,
  name: SERVICE_NAME,
  onLogging: () => {
    const requestId = context.get(Tracing.TRANSACTION_ID);
    return { requestId };
  }
});

export default logger;
