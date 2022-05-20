import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { TypeOrmModule } from '@nestjs/typeorm';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { AuthModule } from './features/auth/auth.module';
import { UserModule } from './features/user/user.module';
import { TaskModule } from './features/task/task.module';
import databaseConfiguration from './configuration/database';
import { CoreModule } from './globalModule/core.module';

@Module({
  imports: [
    ConfigModule.forRoot(),
    TypeOrmModule.forRoot({
      ...databaseConfiguration
    }),
    CoreModule,
    AuthModule,
    UserModule,
    TaskModule
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
