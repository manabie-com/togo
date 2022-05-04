import { KafkaMessage } from 'kafkajs';
import KafkaService from '../common/kafka';
import logger from '../logger';
import taskService from './task.service';
import { TASK_CONSUMER_TOPIC } from './task.topic';

const createTaskConsumer = async (): Promise<void> => {
  await KafkaService.consumeMessage(
    TASK_CONSUMER_TOPIC,
    async (message: KafkaMessage, consumer) => {
      logger.info(`createTaskConsumer >>>>`);
      const value = message.value?.toString() as string;
      await taskService.processTask(value);
      await consumer.commitOffsets([
        { topic: TASK_CONSUMER_TOPIC, offset: message.offset, partition: 0 }
      ]);
    }
  );
};

export default {
  createTaskConsumer
};
