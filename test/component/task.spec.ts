import axios from 'axios';
import { CreateTaskRequestDto } from 'src/modules/task/dto/create-task-request.dto';
import { getConnection } from 'typeorm';
import InitTest from '../init';

describe('task', () => {
  let token;
  let baseUrl;
  let init: InitTest;
  let config;

  beforeAll(async () => {
    init = new InitTest();
    await init.initialize();
    token = await init.getToken();
    baseUrl = init.baseUrl + '/v1/task';
    config = {
      validateStatus: () => true,
      headers: { Authorization: `Bearer ${token}` },
    };
  });

  afterAll(async () => {
    await getConnection().synchronize(true);
    return Promise.resolve();
  });

  it('#POST user create task should work', async () => {
    const request: CreateTaskRequestDto = {
      title: 'task',
      content: 'task',
      startDate: new Date(),
    };

    const result = await axios.post(`${baseUrl}`, request, config);

    expect(result.status).toEqual(201);
  });
});
