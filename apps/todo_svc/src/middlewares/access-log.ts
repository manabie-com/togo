import {inject, injectable, Next, Provider} from '@loopback/core';
import {
  asMiddleware,
  HttpErrors,
  Middleware,
  MiddlewareContext,
  Response,
  RestMiddlewareGroups,
} from '@loopback/rest';
import {RabbitmqBindings, RabbitmqProducer} from 'loopback-rabbitmq';

@injectable(
  asMiddleware({
    group: 'logger',
    upstreamGroups: RestMiddlewareGroups.PARSE_PARAMS,
    downstreamGroups: RestMiddlewareGroups.INVOKE_METHOD,
  }),
)
export class AccessLoggerHandlerMiddlewareProvider
  implements Provider<Middleware>
{
  constructor(
    @inject(RabbitmqBindings.RABBITMQ_PRODUCER)
    private rabbitmqProducer: RabbitmqProducer,
  ) {}

  async value() {
    const middleware: Middleware = async (
      ctx: MiddlewareContext,
      next: Next,
    ) => {
      const {request} = ctx;
      try {
        const messageData = {
          serviceName: 'product_svc',
          createdAt: new Date(),
          url: request.url,
          method: request.method,
          payload: request.body ?? null,
        };
        await this.rabbitmqProducer.publish(
          'messaging.direct',
          'access-log',
          Buffer.from(JSON.stringify(messageData)),
        );

        return await next();
      } catch (err) {
        // Any error handling goes here
        return this.handleError(ctx, err);
      }
    };
    return middleware;
  }

  handleError(context: MiddlewareContext, err: HttpErrors.HttpError): Response {
    // We simply log the error although more complex scenarios can be performed
    // such as customizing errors for a specific endpoint
    throw err;
  }
}
