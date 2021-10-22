import { Test, TestingModule } from '@nestjs/testing';
import { INestApplication } from '@nestjs/common';
import * as request from 'supertest';
import { AppModule } from './../src/app.module';
import { TypeOrmModule } from '@nestjs/typeorm';

describe('AppController (e2e)', () => {
  let app: INestApplication;

  beforeAll(async () => {
    // should use a STAGING or TESTING DB here
    const moduleFixture: TestingModule = await Test.createTestingModule({
      imports: [
        AppModule,
        TypeOrmModule.forRoot({
          type: 'postgres',
          host: 'localhost',
          port: 5432,
          username: 'postgres',
          password: 'postgres',
          database: 'togo',
          entities: ['./dist/**/**.entity{.ts,.js}'],
          synchronize: false,
          logging: false,
        }),
      ],
    }).compile();

    app = moduleFixture.createNestApplication();
    await app.init();
  });

  it('login successfully with right "user_id" and "password"', () => {
    const query = {
      user_name: 'firstUser',
      password: 'example',
    };
    return request(app.getHttpServer())
      .get(`/login?user_id=${query.user_name}&password=${query.password}`)
      .expect(200)
      .expect((res: any) => {
        expect(res.body).toHaveProperty('data');
      });
  });

  it('fails to create task without token', () => {
    return request(app.getHttpServer())
      .post('/tasks')
      .send({ content: 'content' })
      .set('Accept', 'application/json')
      .expect('Content-Type', /json/)
      .expect(401);
  });

  it('fails to list task without token', () => {
    return request(app.getHttpServer())
      .get('/tasks')
      .expect(401);
  });
  afterAll(async () => {
    await app.close();
  });
});
