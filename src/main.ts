import { Logger, ValidationPipe } from '@nestjs/common';
import { NestFactory } from '@nestjs/core';
import * as cookieParser from 'cookie-parser';
import { NestExpressApplication } from '@nestjs/platform-express';
import { AppModule } from './app.module';
import { AllExceptionFilter } from './common/exceptions/exception.filter';
import { TransformInterceptor } from './common/interceptor/transform.interceptor';
import { DocumentBuilder, SwaggerModule } from '@nestjs/swagger';
import { User } from './entities/user.entity';
import { Task } from './entities/task.entity';
import { ToDoList } from './entities/toDoList.entity';

async function bootstrap() {
  const app = await NestFactory.create<NestExpressApplication>(AppModule);
  const config = new DocumentBuilder()
    .setTitle('Todo List example')
    .setDescription('The todo list API description')
    .setVersion('1.0')
    .addServer('/api')
    .addTag('todo')
    .build();
  const document = SwaggerModule.createDocument(app, config, {
    extraModels: [User, Task, ToDoList],
  });
  SwaggerModule.setup('api', app, document);
  app.use(cookieParser());
  app.setGlobalPrefix('api');
  const logger = new Logger();
  app.useGlobalFilters(new AllExceptionFilter(logger));
  app.useGlobalInterceptors(new TransformInterceptor());
  app.enableCors();
  await app.listen(3000);
}
bootstrap();
