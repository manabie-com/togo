import { ValidationPipe } from '@nestjs/common';
import { NestFactory } from '@nestjs/core';
import * as express from 'express';
import { AppModule } from './app.module';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  app.useGlobalPipes(
    new ValidationPipe({
      whitelist: true,
    }),
  );

  app.use(express.urlencoded({ limit: '50mb', extended: true }));

  const port = process.env.PORT || 8080;

  await app.listen(port, () => {
    console.log(`Server listening on http://localhost:${port}`);
  });
}
bootstrap();
