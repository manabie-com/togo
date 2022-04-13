import { ValidationPipe } from '@nestjs/common';
import { NestFactory } from '@nestjs/core';
import { NestExpressApplication } from '@nestjs/platform-express';
import * as express from 'express';
import { AppModule } from '../src/app.module';

async function startApp() {
  const app = await NestFactory.create<NestExpressApplication>(AppModule);
  app.useGlobalPipes(
    new ValidationPipe({
      whitelist: true,
    }),
  );
  app.enableCors();
  app.use(express.urlencoded({ limit: '50mb', extended: true }));
  app.enableShutdownHooks();

  const port = 8081;

  await app.listen(port, () => {
    console.log(`Server listening on http://localhost:${port}`);
  });

  return app;
}

export default startApp;
