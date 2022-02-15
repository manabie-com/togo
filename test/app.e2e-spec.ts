import { Test, TestingModule } from '@nestjs/testing';
import { INestApplication } from '@nestjs/common';
import * as request from 'supertest';
import { AppModule } from '../src/app.module';
import { UsersModule } from '../src/users/users.module';
import { TasksModule } from '../src/tasks/tasks.module';
import { MongooseModule } from '@nestjs/mongoose';
import { constants } from '../src/constants';

describe('AppController (e2e)', () => {
  let app: INestApplication;
  const newAccount = {
    username: 'admin123',
    password: 'admin123',
    displayName: 'admin',
  };
  const account = {
    username: 'admin123',
    password: 'admin123',
  };
  const newTask = {
    description: 'solve challenge',
  };
  const mockCreatedUserResp = {
    limitedTaskPerDay: 3,
    displayName: 'admin',
    username: 'admin123',
    __v: 0,
  };

  let token = null;

  beforeEach(async () => {
    const moduleFixture: TestingModule = await Test.createTestingModule({
      imports: [AppModule, UsersModule, TasksModule],
    }).compile();

    app = moduleFixture.createNestApplication();
    await app.init();
  });

  it('create new user', () => {
    return request(app.getHttpServer())
      .post('/users')
      .send(newAccount)
      .expect(201)
      .then((res) => {
        expect(res.body).toEqual(expect.objectContaining(mockCreatedUserResp));
      });
  });

  it('get access token', () => {
    return request(app.getHttpServer())
      .post('/auth/login')
      .send(account)
      .expect(201)
      .then((res) => {
        token = res.body?.access_token;
        expect(res.body).toHaveProperty('access_token');
      });
  });

  it('create new task', () => {
    return request(app.getHttpServer())
      .post('/tasks')
      .set('Authorization', 'Bearer ' + token)
      .send(newTask)
      .expect([201, 400]);
  });

  // it('get tasks by owner', () => {
  //   return request(app.getHttpServer())
  //     .get('/tasks')
  //     .set('Authorization', 'Bearer ' + token)
  //     .query({ page: 1, size: 10 })
  //     .expect([200, 401])
  //     .then((res) => {
  //       expect(res.body).toHaveProperty(['data', 'page', 'size', 'total']);
  //     });
  // });
});
