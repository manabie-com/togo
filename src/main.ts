import { NestFactory } from '@nestjs/core';
import { DocumentBuilder, SwaggerModule } from '@nestjs/swagger';
import { AppModule } from './app.module';
import { HttpExceptionFilter } from './shared/filter/http-exception.filter';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  //app.useGlobalPipes(new ValidationPipe());
  app.useGlobalFilters(new HttpExceptionFilter());
  const options = new DocumentBuilder()
    .setTitle('Tasks')
    .setDescription('The Tasks API')
    .setVersion('1.0')
    .build();
  
  const document = SwaggerModule.createDocument(app, options);

  SwaggerModule.setup('api', app, document);
    
  await app.listen(3000);
}
bootstrap();
