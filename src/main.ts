import {
  ClassSerializerInterceptor,
  HttpStatus,
  UnprocessableEntityException,
  ValidationPipe
} from '@nestjs/common';
import { NestFactory, Reflector } from '@nestjs/core';
import {
  FastifyAdapter,
  NestFastifyApplication
} from '@nestjs/platform-fastify';
import { DocumentBuilder, SwaggerModule } from '@nestjs/swagger';
import compression from 'compression';
import fastifyCookie from 'fastify-cookie';
import fastifyCsrf from 'fastify-csrf';
import { fastifyHelmet } from 'fastify-helmet';
import fmp from 'fastify-multipart';

import { AppModule } from './app.module';
import { ErrorExceptionFilter } from './filters/error.filter';

export async function bootstrap(): Promise<NestFastifyApplication> {
  const app = await NestFactory.create<NestFastifyApplication>(
    AppModule,
    new FastifyAdapter()
  );

  await app.register(fastifyCookie);
  await app.register(fmp);
  await app.register(fastifyCsrf, { cookieKey: 'X-CSRF-Token' });
  await app.register(fastifyHelmet);

  app.use(compression());

  const reflector = app.get(Reflector);

  app.useGlobalFilters(new ErrorExceptionFilter());

  app.useGlobalInterceptors(new ClassSerializerInterceptor(reflector));

  app.useGlobalPipes(
    new ValidationPipe({
      whitelist: true,
      errorHttpStatusCode: HttpStatus.UNPROCESSABLE_ENTITY,
      transform: true,
      dismissDefaultMessages: true,
      exceptionFactory: (errors) => new UnprocessableEntityException(errors)
    })
  );

  const config = new DocumentBuilder()
    .setTitle('NestJs Core APIs')
    .setDescription('NestJs Core APIs')
    .setVersion('0.0.1')
    .addBearerAuth()
    .build();
  const document = SwaggerModule.createDocument(app, config);
  SwaggerModule.setup('swagger', app, document);

  const port = process.env.PORT || 8080;
  await app.listen(port, '0.0.0.0', (err: Error, address: string) => {
    if (!err) {
      console.log(`\n\n\nServer started at ${address}\n\n`);

      return;
    }

    console.log(err);
  });

  console.info(`server running on port ${port}`);

  return app;
}

void bootstrap();
