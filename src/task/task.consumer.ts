import { Message } from 'kafkajs';
import KafkaService from '../common/kafka';
import logger from '../logger';
import taskService from './task.service';
import { TASK_CONSUMER_TOPIC } from './task.topic';

const createTaskConsumer = async (): Promise<void> => {
  await KafkaService.consumeMessage(
    TASK_CONSUMER_TOPIC,
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
