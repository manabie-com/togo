import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { ConfigModule } from '@nestjs/config';
import databaseConfig from './config/database.config';
import { DatabaseModule } from './database/database.module';
import { UserController } from './controllers/user.controller';
import { ToDoService } from './services/todo.service';
import { TaskService } from './services/task.service';
import { TypeOrmModule } from '@nestjs/typeorm';
import { Task } from './entities/task.entity';
import { User } from './entities/user.entity';
import { ToDoList } from './entities/toDoList.entity';
import { UserService } from './services/user.service';

@Module({
  imports: [
    ConfigModule.forRoot({
      load: [databaseConfig],
    }),
    DatabaseModule,
    TypeOrmModule.forFeature([Task, User, ToDoList]),
  ],
  controllers: [AppController, UserController],
  providers: [AppService, UserService, TaskService, ToDoService],
})
export class AppModule {}
