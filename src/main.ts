import { NestFactory } from '@nestjs/core';
import { ValidationPipe } from '@nestjs/common';
import { DocumentBuilder, SwaggerModule } from '@nestjs/swagger';

import { AppModule } from './app.module';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  app.useGlobalPipes(new ValidationPipe({
    transform: true,
    whitelist: true,
    validationError: {
      target: true,
      value: true,
    }
  }));

  const document = SwaggerModule.createDocument(
    app,
    new DocumentBuilder()
        .setTitle('Manabie API')
        .setVersion('1.0')
        .build()
  );
  SwaggerModule.setup('v1', app, document);
  await app.listen(3000);
}
bootstrap();
