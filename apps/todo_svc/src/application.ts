import {AuthenticationComponent} from '@loopback/authentication';
import {
  JWTAuthenticationComponent,
  MyUserService,
  UserServiceBindings,
} from '@loopback/authentication-jwt';
import {BootMixin} from '@loopback/boot';
import {ApplicationConfig, createBindingFromClass} from '@loopback/core';
import {RepositoryMixin} from '@loopback/repository';
import {RestApplication} from '@loopback/rest';
import {
  RestExplorerBindings,
  RestExplorerComponent,
} from '@loopback/rest-explorer';
import {ServiceMixin} from '@loopback/service-proxy';
import {
  ConsumersBooter,
  MessageHandlerErrorBehavior,
  QueueComponent,
  RabbitmqBindings,
  RabbitmqComponent,
  RabbitmqComponentConfig,
} from 'loopback-rabbitmq';
import path from 'path';
// import {DbDataSource} from './datasources';
import {MongodbDataSource} from './datasources';
import {AccessLoggerHandlerMiddlewareProvider} from './middlewares';
import {MySequence} from './sequence';
export {ApplicationConfig};

export class TodoListApplication extends BootMixin(
  ServiceMixin(RepositoryMixin(RestApplication)),
) {
  constructor(options: ApplicationConfig = {}) {
    super(options);
    // add middleware
    this.add(createBindingFromClass(AccessLoggerHandlerMiddlewareProvider));

    // Set up the custom sequence
    this.sequence(MySequence);

    // Set up default home page
    this.static('/', path.join(__dirname, '../public'));

    // Customize @loopback/rest-explorer configuration here
    this.configure(RestExplorerBindings.COMPONENT).to({
      path: '/explorer',
    });
    this.component(RestExplorerComponent);

    this.projectRoot = __dirname;
    // Customize @loopback/boot Booter Conventions here
    this.bootOptions = {
      controllers: {
        // Customize ControllerBooter Conventions here
        dirs: ['controllers'],
        extensions: ['.controller.js'],
        nested: true,
      },
    };

    this.configRabbitmq();

    // ------ ADD SNIPPET AT THE BOTTOM ---------
    // Mount authentication system
    this.component(AuthenticationComponent);
    // Mount jwt component
    this.component(JWTAuthenticationComponent);
    // Bind datasource
    // this.dataSource(DbDataSource, UserServiceBindings.DATASOURCE_NAME);
    this.dataSource(MongodbDataSource, UserServiceBindings.DATASOURCE_NAME);
    // ------------- END OF SNIPPET -------------

    //new
    this.bind(UserServiceBindings.USER_SERVICE).toClass(MyUserService);

    //
  }

  configRabbitmq() {
    this.configure<RabbitmqComponentConfig>(RabbitmqBindings.COMPONENT).to({
      options: {
        protocol: process.env.RABBITMQ_PROTOCOL ?? 'amqp',
        hostname: process.env.RABBITMQ_HOST ?? 'rabbitmq',
        port:
          process.env.RABBITMQ_PORT === undefined
            ? 5672
            : +process.env.RABBITMQ_PORT,
        username: process.env.RABBITMQ_USER ?? 'admin',
        password: process.env.RABBITMQ_PASS ?? 'admin',
        vhost: process.env.RABBITMQ_VHOST ?? '/',
      },
      // configurations below are optional, (Default values)
      producer: {
        idleTimeoutMillis: 10000,
      },
      consumer: {
        retries: 0, // number of retries, 0 is forever
        interval: 1500, // retry interval in ms
      },
      defaultConsumerErrorBehavior: MessageHandlerErrorBehavior.ACK,
      prefetchCount: 10,
      exchanges: [
        {
          name: 'loopback.direct',
          type: 'direct', // A direct exchange delivers messages to queues based on the message routing key.
          // type: 'fanout' // A fanout exchange routes messages to all queues that are linked
        },
        {
          name: 'messaging.direct',
          type: 'direct',
        },
      ],
    });
    this.component(RabbitmqComponent);
    this.booters(ConsumersBooter);
    this.component(QueueComponent);
  }
}
