import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { TypeOrmModule } from '@nestjs/typeorm';
import { ConfigModule } from '@nestjs/config';
import { UserData } from './interfaces/user_data.interface';
import { UserModule } from './modules/user-module/user.module';

declare module "express" {
  export interface Request {
    user: UserData
  }
}
@Module({
  imports: [
    TypeOrmModule.forRoot({
      type: 'mysql',
      host: '127.0.0.1',
      port: 3306,
      username: 'root',
      password: '1',
      database: 'togo',
      synchronize: true,
      entities: ["dist/**/*.entity{.ts,.js}"],
    }),
    ConfigModule.forRoot({
      isGlobal: true,
      envFilePath: ".env"
    }),
    UserModule
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
