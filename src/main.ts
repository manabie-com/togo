import { ValidationPipe } from '@nestjs/common';
import { NestFactory } from '@nestjs/core';
import { DocumentBuilder, SwaggerModule } from '@nestjs/swagger';
import { AppModule } from './modules/app.module';
import { DispatchError } from './modules/config/errors';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.useGlobalFilters(new DispatchError());

  app.enableCors();
  app.setGlobalPrefix('api');

  const options = new DocumentBuilder()
    .addBearerAuth()
    .setTitle('Tasks Applicants Service')
    .setDescription('The Enrolment Applicants Service API description')
    .setVersion('1.0')
    .build();

  const document = SwaggerModule.createDocument(app, options);

  SwaggerModule.setup('v1', app, document);

  app.useGlobalPipes(
    new ValidationPipe({
      transform: true,
      whitelist: true,
      validationError: {
        target: true,
        value: true,
      },
    })
  );
  await app.listen(process.env.PORT || 3002, '0.0.0.0');
}
bootstrap();
