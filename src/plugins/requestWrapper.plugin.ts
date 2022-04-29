import Hapi from '@hapi/hapi';
import context from '../common/context';
import { v4 as uuidv4 } from 'uuid';
import { Tracing } from '../common/constants';

const handleHapiRequest = async (
  _hapiRequest: Hapi.Request,
  hapiResponse: Hapi.ResponseToolkit
) => {
  context.set(Tracing.TRANSACTION_ID, uuidv4());
  return hapiResponse.continue;
};

// eslint-disable-next-line @typescript-eslint/ban-types
const requestWrapper: Hapi.Plugin<{}> = {
  name: 'requestWrapper',
  version: '1.0.0',
  register: (server: Hapi.Server) => {
    server.ext('onRequest', handleHapiRequest);
  },
  once: true
};

export default requestWrapper;
