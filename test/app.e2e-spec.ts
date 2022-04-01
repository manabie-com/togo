import { Test } from '@nestjs/testing';
import { INestApplication, ValidationPipe } from '@nestjs/common';
import * as request from 'supertest';

import { AppModule } from './../src/app.module';

describe('TodoController (e2e)', () => {
  let app: INestApplication;
  let userId: number = 0;

  beforeAll(async () => {
    const moduleFixture = await Test.createTestingModule({
      imports: [AppModule],
    }).compile();

    app = moduleFixture.createNestApplication();
    app.useGlobalPipes(new ValidationPipe({
      transform: true,
      whitelist: true,
      validationError: {
        target: true,
        value: true,
      }
    }));
    await app.init();
  });

  it('/todo (POST) Create task faild with data empty', () => {
    const data = {};
    return request(app.getHttpServer())
      .post('/todo')
      .send(data)
      .expect(400)
      .expect({
        "statusCode": 400,
        "message": [ 'task must be a string', 'task should not be empty' ],
        "error": "Bad Request"
      });
  });

  it('/todo (POST) Create task faild with task empty', () => {
    const data = { 'task': '' };
    return request(app.getHttpServer())
      .post('/todo')
      .send(data)
      .expect(400)
      .expect({
        "statusCode": 400,
        "message": ["task should not be empty"],
        "error": "Bad Request"
      });
  });

  it('/todo (POST) Create task faild with user id not exist', () => {
    const data = {
      'task': 'string',
      'user_id': 0
    };
    return request(app.getHttpServer())
      .post('/todo')
      .send(data)
      .expect(400)
      .expect({
        'statusCode': 400,
        'message': 'User not exist',
        'error': 'Bad Request'
      });
  });

  it('/todo (POST) Create task and user success with limit task for user', async () => {
    const data = {
      'task': 'test 1',
      'limit_task': 1,
    };
    const respronse = await request(app.getHttpServer())
      .post('/todo')
      .send(data)
      .expect(201)

    userId = respronse.body && respronse.body.userId ? respronse.body.userId : 0;
    return;
  });

  it('/todo (POST) Create task falid with length task bigger 1', () => {
    const data = {
      'task': 'test 1',
      'user_id': userId
    };
    return request(app.getHttpServer())
      .post('/todo')
      .send(data)
      .expect(400)
      .expect({
        "statusCode": 400,
        "message": "Overload task in 1 day",
        "error": "Bad Request"
      });
  });

  it('/todo (POST) Create task and create user success with user id empty', async () => {
    const data = {
      'task': 'test 1',
    };
    const respronse = await request(app.getHttpServer())
      .post('/todo')
      .send(data)
      .expect(201)

    userId = respronse.body && respronse.body.userId ? respronse.body.userId : 0;
    return;
  });

  it('/todo (POST) Create task success with user id just created', () => {
    const data = {
      'task': 'test 2',
      'user_id': userId,
    };
    return request(app.getHttpServer())
      .post('/todo')
      .send(data)
      .expect(201)
  });

  it('/todo (POST) Create task falid with user id just created and limit by 2', () => {
    const data = {
      'task': 'test 3',
      'user_id': userId,
      'limit_task': 2,
    };
    return request(app.getHttpServer())
      .post('/todo')
      .send(data)
      .expect(400)
      .expect({
        "statusCode": 400,
        "message": "Overload task in 1 day",
        "error": "Bad Request"
      });
  });

  it('/todo (POST) Create task success with user id just created and limit by 3', () => {
    const data = {
      'task': 'test 4',
      'user_id': userId,
      'limit_task': 3,
    };
    return request(app.getHttpServer())
      .post('/todo')
      .send(data)
      .expect(201)
  });

  it('/todo (POST) Create task falid with overload limit in 1 day', () => {
    const data = {
      'task': 'test 5',
      'user_id': userId
    };
    return request(app.getHttpServer())
      .post('/todo')
      .send(data)
      .expect(400)
      .expect({
        "statusCode": 400,
        "message": "Overload task in 1 day",
        "error": "Bad Request"
      });
  });

  it('/todo (POST) Create task with user id just created and limit by 4', () => {
    const data = {
      'task': 'test 6',
      'user_id': userId,
      'limit_task': 4,
    };
    return request(app.getHttpServer())
      .post('/todo')
      .send(data)
      .expect(201)
  });

  afterAll(async () => {
    if (app)
      await app.close();
  });
});
