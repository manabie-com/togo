import { Test } from '@nestjs/testing';
import { INestApplication, ValidationPipe } from '@nestjs/common';
import * as request from 'supertest';

import { AppModule } from './../src/app.module';

describe('TodoController (e2e)', () => {
  let app: INestApplication;

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

  it('/todo (POST) Create with data empty', () => {
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

  it('/todo (POST) Create with task empty', () => {
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

  it('/todo (POST) Create with user id not exist', () => {
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

  afterAll(async () => {
    await app.close();
  });
});
