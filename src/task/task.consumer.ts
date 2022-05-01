import { Message } from 'kafkajs';
import KafkaService from '../common/kafka';
import logger from '../logger';
import taskService from './task.service';

const createTaskConsumer = async (): Promise<void> => {
  await KafkaService.consumeMessage(
    'task-consumer',
    async (message: Message) => {
      logger.info(`createTaskConsumer >>>>`);
      const value = message.value?.toString() as string;
      await taskService.processTask(value);
    }
  );
};

export default {
  createTaskConsumer
};
