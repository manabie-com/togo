import { Test, TestingModule } from '@nestjs/testing';
import { INestApplication, ValidationPipe } from '@nestjs/common';
import * as request from 'supertest';
import { AppModule } from './../src/app.module';
import { TaskService } from '../src/model/task/service/task.service';

describe('AppController (e2e)', () => {
  let app: INestApplication;

  beforeEach(async () => {
    const moduleFixture: TestingModule = await Test.createTestingModule({
      imports: [AppModule],
    }).compile();

    app = moduleFixture.createNestApplication();
    await app.init();
  });

  it('/ (GET)', () => {
    return request(app.getHttpServer())
      .get('/')
      .expect(200)
      .expect('Hello World!');
  });

  afterAll(async () => {
    await app.close();
  });
});

describe('ApiController (e2e)', () => {
  let app: INestApplication;
  let taskService: TaskService;

  beforeAll(async () => {
    const moduleFixture: TestingModule = await Test.createTestingModule({
      imports: [AppModule],
    }).compile();
    taskService = moduleFixture.get<TaskService>(TaskService);
    app = moduleFixture.createNestApplication();
    app.useGlobalPipes(
      new ValidationPipe({
        transform: true,
      }),
    );
    await app.init();
    await taskService.destroyDB();
  });

  it('/task (POST) - missing title', () => {
    return request(app.getHttpServer())
      .post('/task')
      .send({ userId: 1, desc: 'lam rat nhieu bai tap' })
      .expect(400);
    // .catch((err) => console.log(err.message));
  });

  it('/task (POST) - missing desc', () => {
    return request(app.getHttpServer())
      .post('/task')
      .send({ userId: 1, title: 'lam bai tap' })
      .expect(400);
    // .catch((err) => console.log(err.message));
  });

  it('/task (POST) - missing desc, title', () => {
    return request(app.getHttpServer())
      .post('/task')
      .send({ userId: 1 })
      .expect(400);
    // .catch((err) => console.log(err.message));
  });

  it('/task (POST) - complete', () => {
    return request(app.getHttpServer())
      .post('/task')
      .send({ userId: 1, title: 'lam bai tap', desc: 'lam rat nhieu bai tap' })
      .then((result) => {
        expect(result.body?.title).toEqual('lam bai tap');
        expect(result.body?.desc).toEqual('lam rat nhieu bai tap');
      });
    // .catch((err) => console.log(err.message));
  });

  it('/task (POST) - limited task in on day (5 times)', async () => {
    for (let i = 0; i < 5; i++) {
      await request(app.getHttpServer()).post('/task').send({
        userId: 1,
        title: 'lam bai tap',
        desc: 'lam rat nhieu bai tap',
      });
    }
    return await request(app.getHttpServer())
      .post('/task')
      .send({ userId: 1, title: 'lam bai tap', desc: 'lam rat nhieu bai tap' })
      .then((result) => {
        expect(result.body?.message).toEqual('The limited tasks in day');
        expect(result.statusCode).toEqual(201);
      });
    // .catch((err) => console.log(err.message));
  });

  afterAll(async () => {
    await app.close();
  });
});
