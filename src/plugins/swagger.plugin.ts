import * as HapiSwagger from 'hapi-swagger';
import * as Package from '../../package.json';
import { displayErrorsDescription } from '../common/util';
import { ErrorList } from '../error/error.list';

const swaggerOptions: HapiSwagger.RegisterOptions = {
  info: {
    title: 'CMS API Documentation',
    version: Package.version,
    description: displayErrorsDescription(ErrorList)
  }
};

export default {
  plugin: HapiSwagger,
  options: swaggerOptions
};
