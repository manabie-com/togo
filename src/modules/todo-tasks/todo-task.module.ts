import { Module } from '@nestjs/common';

import { TodoTaskService } from './todo-task.service';
import { TodoTaskController } from './todo-task.controller';
import { TodoTask } from './entities/todo-task.entity';
import { TypeOrmModule } from '@nestjs/typeorm';
import { UserService } from '@modules/users/user.service';
import { User } from '@modules/users/entities/user.entity';

@Module({
  imports: [TypeOrmModule.forFeature([TodoTask, User])],
  controllers: [TodoTaskController],
  providers: [TodoTaskService, UserService],
})
export class TodoTaskModule {}
