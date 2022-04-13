import { Connection, createConnection } from 'typeorm';
import { UserEntity } from '../src/modules/user/entity/user.entity';
import { TaskEntity } from '../src/modules/task/entity/task.entity';
import { EnvironmentService } from '../src/common/environment/services/environment.service';
import axios from 'axios';

class InitTest {
  user = {
    username: 'togo',
    password: '123456',
  };
  baseUrl = 'http://localhost:8081';
  connection: Connection;
  userToken: string;
  environmentService = new EnvironmentService();

  async initialize(): Promise<void> {
    await this.initDatabase();
    await this.initBasicData();
  }

  async initDatabase(): Promise<void> {
    this.connection = await createConnection({
      type: 'sqlite',
      database: './test.db',
      dropSchema: true,
      entities: [UserEntity, TaskEntity],
      synchronize: true,
    });
  }

  async initBasicData(): Promise<void> {
    const user = await axios.post(
      `${this.baseUrl}/v1/auth/register`,
      this.user,
    );

    expect(user.status).toEqual(201);
  }

  async getToken(): Promise<string> {
    const token = await axios.post(`${this.baseUrl}/v1/auth/login`, this.user);

    expect(token.status).toEqual(201);

    return token.data.token as unknown as string;
  }
}

export default InitTest;
