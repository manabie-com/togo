import { Module } from '@nestjs/common';
import { TodoController } from './todo.controller.js';
import { TodoService } from './todo.service.js';
import { UserService } from '../users/users.service.js';
import { TypeOrmModule } from '@nestjs/typeorm';
import { Todo } from './todo.entity.js';
import { UserModule } from '../users/users.module.js';

@Module({
  imports: [TypeOrmModule.forFeature([Todo]), UserModule],
  controllers: [TodoController],
  providers: [TodoService, UserService],
})
export class TodoModule {}
