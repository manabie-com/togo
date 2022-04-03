import { Test, TestingModule } from '@nestjs/testing';
import { INestApplication, Logger } from '@nestjs/common';
import * as request from 'supertest';
import { AppModule } from './../src/app.module';
import * as cookieParser from 'cookie-parser';
import { AllExceptionFilter } from 'src/common/exceptions/exception.filter';
import { TransformInterceptor } from 'src/common/interceptor/transform.interceptor';

describe('AppController (e2e)', () => {
  let app: INestApplication;

  beforeEach(async () => {
    const moduleFixture: TestingModule = await Test.createTestingModule({
      imports: [AppModule],
    }).compile();

    app = moduleFixture.createNestApplication();
    app.use(cookieParser());
    app.setGlobalPrefix('api');
    const logger = new Logger();
    app.useGlobalFilters(new AllExceptionFilter(logger));
    app.useGlobalInterceptors(new TransformInterceptor());
    app.enableCors();
    await app.init();
  });

  it('/ (GET)', () => {
    return request(app.getHttpServer()).get('/api').expect(200).expect({
      success: true,
      message: null,
      data: 'Hello World!',
    });
  });
  it('/ GET USERS', () => {
    return request(app.getHttpServer()).get('/api/users').expect(200).expect({
      success: true,
      message: null,
      data: 'Hello World!',
    });
  });
});
