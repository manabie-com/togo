import { Logger, ValidationPipe } from '@nestjs/common';
import { NestFactory } from '@nestjs/core';
import * as cookieParser from 'cookie-parser';
import { NestExpressApplication } from '@nestjs/platform-express';
import { AppModule } from './app.module';
import { AllExceptionFilter } from './common/exceptions/exception.filter';
import { TransformInterceptor } from './common/interceptor/transform.interceptor';

async function bootstrap() {
  const app = await NestFactory.create<NestExpressApplication>(AppModule);
  app.use(cookieParser());
  app.setGlobalPrefix('api');
  const logger = new Logger();
  app.useGlobalFilters(new AllExceptionFilter(logger));
  app.useGlobalInterceptors(new TransformInterceptor());
  app.enableCors();
  await app.listen(3000);
}
bootstrap();
