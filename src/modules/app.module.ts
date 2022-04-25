import { Module } from '@nestjs/common';
import { SequelizeModule } from '@nestjs/sequelize';
import { config } from 'dotenv';
import { Task, User, UserSettingTask } from './common/entities';
import { AuthModule } from './features/auth/auth.module';
config();
@Module({
  imports: [
    SequelizeModule.forRootAsync({
      useFactory: () => ({
        dialect: 'postgres',
        host: process.env.DB_HOST,
        port: +process.env.DB_PORT,
        username: process.env.DB_USER_NAME,
        password: process.env.DB_PASSWORD,
        database: process.env.DB_DATABASE,
        models: [Task, UserSettingTask, User],
      }),
    }),
    AuthModule,
  ],
})
export class AppModule {}
