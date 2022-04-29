import Hapi from '@hapi/hapi';
import logger from '../logger';
import { AppError } from '../error/error.service';

const documentPathRegex = /^\/(documentation|swagger)/;

const handleHapiResponse = (
  hapiRequest: Hapi.Request,
  hapiResponse: Hapi.ResponseToolkit
) => {
  // ignore document ui path
  if (documentPathRegex.test(hapiRequest.url.pathname))
    return hapiResponse.continue;

  const responseData = hapiResponse.request.response;

  if (responseData instanceof AppError) {
    const error = responseData.getErrors();
    logger.error(error.message, { error });
    return hapiResponse.response(error).code(error.statusCode);
  }
  if (responseData instanceof Error) {
    logger.error(responseData.message, { message: responseData.stack });
    return hapiResponse
      .response(responseData.output)
      .code(responseData.output.statusCode);
  }

  if (responseData.variety === 'plain') {
    return hapiResponse
      .response({ data: responseData.source } as Hapi.ResponseValue)
      .code(responseData.statusCode);
  }
  return hapiResponse.continue;
};

// eslint-disable-next-line @typescript-eslint/ban-types
const responseWrapper: Hapi.Plugin<{}> = {
  name: 'responseWrapper',
  version: '1.0.0',
  register: (server: Hapi.Server) => {
    server.ext('onPreResponse', handleHapiResponse);
  },
  once: true
};

export default responseWrapper;
