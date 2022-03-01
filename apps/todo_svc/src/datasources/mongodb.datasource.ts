import {inject, lifeCycleObserver, LifeCycleObserver} from '@loopback/core';
import {juggler} from '@loopback/repository';


const {MONGO_HOST, MONGO_PORT, MONGO_USER, MONGO_PASS} = process.env;

const config = {
  name: 'mongodb',
  connector: 'mongodb',
  // uncomment to use url setting to override other settings, ex mongodb://username:password@hostname:port/database
  // url: 'mongodb:27017/simple-app',
  host: MONGO_HOST,
  port: MONGO_PORT,
  user: MONGO_USER,
  password: MONGO_PASS,
  database: 'simple-app',
  useNewUrlParser: true,
  authSource: 'admin',
};

// Observe application's life cycle to disconnect the datasource when
// application is stopped. This allows the application to be shut down
// gracefully. The `stop()` method is inherited from `juggler.DataSource`.
// Learn more at https://loopback.io/doc/en/lb4/Life-cycle.html
@lifeCycleObserver('datasource')
export class MongodbDataSource extends juggler.DataSource
  implements LifeCycleObserver {
  static dataSourceName = 'mongodb';
  static readonly defaultConfig = config;

  constructor(
    @inject('datasources.config.mongodb', {optional: true})
    dsConfig: object = config,
  ) {
    super(dsConfig);
  }
}
