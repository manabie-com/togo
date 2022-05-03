import { Consumer, Kafka, Message } from 'kafkajs';
import { KAFKA_URL, KAFKA_GROUP_ID, SERVICE_NAME } from '../config';

interface IKafkaService {
  clientId?: string;
  brokers: string[];
}

class KafkaService {
  private _kafka: Kafka;
  constructor(options: IKafkaService) {
    this._kafka = new Kafka(options);
  }

  produceMessage = async (
    topic: string,
    messages: Message[]
  ): Promise<void> => {
    const producer = this._kafka.producer();
    await producer.connect();
    await producer.send({
      topic,
      messages
    });

    await producer.disconnect();
  };

  consumeMessage = async (
    selectedTopic: string,
    handler: (message: any, consumer: Consumer) => Promise<void> | void
  ) => {
    const consumer = this._kafka.consumer({ groupId: KAFKA_GROUP_ID || '' });

    await consumer.connect();
    await consumer.subscribe({ topic: selectedTopic, fromBeginning: true });

    await consumer.run({
      autoCommit: false,
      eachMessage: async ({ message }) => {
        handler(message, consumer);
      }
    });
  };
}

export default new KafkaService({
  clientId: SERVICE_NAME,
  brokers: [KAFKA_URL || '']
});
