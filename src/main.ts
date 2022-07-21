import { NestFactory } from '@nestjs/core';
import { DocumentBuilder, SwaggerModule } from '@nestjs/swagger';
import { AppModule } from './app.module';
import { version } from 'package.json';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  if (process.env.ENV === 'development') {
    const config = new DocumentBuilder()
      .setTitle('Togo api documentation')
      .setDescription(`Togo api documentation v0`)
      .setVersion(version)
      .addBearerAuth()
      .build();
    const document = SwaggerModule.createDocument(app, config);
    SwaggerModule.setup('swagger', app, document);
  }

  await app.listen(3000);
}
bootstrap();
