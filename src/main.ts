import { Logger, ValidationPipe } from '@nestjs/common'
import { NestFactory } from '@nestjs/core'
import { AppModule } from './app.module'
import { AllExceptionsFilter } from './all-exceptions.filter'
import * as morgan from 'morgan';

async function bootstrap() {
  const app = await NestFactory.create(AppModule)
  app.use(morgan(
    ':remote-addr - :remote-user [:date[clf]] ":method :url HTTP/:http-version" :status :response-time ms :res[content-length] ":referrer" ":user-agent"',
  ))
  app.useGlobalPipes(new ValidationPipe())
  app.useGlobalFilters(new AllExceptionsFilter())

  const port = 3000
  await app.listen(port)
  Logger.log(`app listening on port: ${port}`)
}
bootstrap()
