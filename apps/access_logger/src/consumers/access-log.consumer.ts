import {repository} from '@loopback/repository';
import {ConsumeMessage, Nack, rabbitConsume} from 'loopback-rabbitmq';
import {AccessLog} from '../models';
import {AccessLogRepository} from '../repositories';

// interface Message {
//   createdAt: Date;
//   serviceName: string;
//   url: string;
//   method: string;
//   payload: object;
// }

export class AccessLogConsumer {
  constructor(
    @repository(AccessLogRepository)
    private accessLogRepository: AccessLogRepository,
  ) {
    console.log('AccessLogConsumer');
  }

  @rabbitConsume({
    exchange: 'messaging.direct',
    routingKey: 'access-log',
    queue: 'access-log-queue',
  })
  async handle(message: AccessLog, rawMessage: ConsumeMessage) {
    try {
      await this.accessLogRepository.create(message);
    } catch (error) {
      console.error(error);
      // retry at least once.
      if (rawMessage?.fields?.redelivered) {
        return new Nack(false);
      }
      return new Nack(true);
    }
  }
}
