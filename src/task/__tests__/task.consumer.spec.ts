import taskConsumer from '../task.consumer';
import taskService from '../task.service';
import KafkaService from '../../common/kafka';

jest.mock('../../common/kafka');
jest.mock('../task.service');

describe('task.consumer', () => {
  it('Should consume message successfully', async () => {
    const expectedPayload = { value: '_test' };
    const commitOffsets = jest.fn();
    const consumer = jest.fn().mockImplementationOnce(() => ({
      commitOffsets
    }));
    (taskService.processTask as jest.Mock).mockResolvedValueOnce(null);
    (KafkaService.consumeMessage as jest.Mock).mockImplementationOnce(
      (_topic, handler) => {
        handler(expectedPayload, consumer);
      }
    );
    await taskConsumer.createTaskConsumer();

    expect(KafkaService.consumeMessage.mock.calls[0][0]).toBe('task-consumer');
    expect(taskService.processTask).toBeCalledWith(expectedPayload.value);
  });
});
