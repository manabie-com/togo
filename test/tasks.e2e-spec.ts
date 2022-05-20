import { Test, TestingModule } from '@nestjs/testing';
import { HttpStatus, INestApplication } from '@nestjs/common';
import * as request from 'supertest';
import { Repository } from 'typeorm';
import { AuthModule } from '../src/features/auth/auth.module';
import { CoreModule } from '../src/globalModule/core.module';
import { UserModule } from '../src/features/user/user.module';
import { AppModule } from '../src/app.module';
import { TaskEntity } from '../src/features/task/entities/task.entity';

describe('TaskController (e2e)', () => {
  let app: INestApplication;

  beforeEach(async () => {
    const moduleFixture: TestingModule = await Test.createTestingModule({
      imports: [AppModule, UserModule, AuthModule, CoreModule]
    })
    .compile();

    app = moduleFixture.createNestApplication();
    app.setGlobalPrefix('/api/v1');
    await app.init();
  });

  describe('create', () => {
    it('shoud create a new task record', async () => {
      const signinResponse = await request(app.getHttpServer())
      .post('/api/v1/signin')
      .send({ email: 'admin@gmail.com', password: 'admin123' })
      .expect(HttpStatus.CREATED);
  
      const { token } = signinResponse.body;

      return request(app.getHttpServer())
        .post('/api/v1/tasks')
        .set('Authorization', 'Bearer ' + token)
        .send({
          content: "Task One"
        })
        .expect(HttpStatus.CREATED);
    });
  })

  afterAll(async () => {
    await app.close();
  });
});
