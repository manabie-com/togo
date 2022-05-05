import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { SequelizeModule } from '@nestjs/sequelize';
import { config } from 'dotenv';
import { Task, User, UserSettingTask } from './common/entities';
import { AuthModule } from './features/auth/auth.module';
import { InitDataModule } from './features/init-data/init-data.module';
import { TaskModule } from './features/tasks/task.module';
config();
@Module({
  imports: [
    SequelizeModule.forRootAsync({
      useFactory: () => ({
        dialect: 'postgres',
        host: process.env.DB_HOST,
        port: +process.env.DB_PORT,
        username: process.env.DB_USER,
        password: process.env.DB_PASSWORD,
        database: process.env.DB_DATABASE,
        autoLoadModels: true,
        synchronize: true,

        models: [Task, UserSettingTask, User],
      }),
    }),
    AuthModule,
    TaskModule,
    InitDataModule,
    ConfigModule.forRoot(),
  ],
})
export class AppModule {}
