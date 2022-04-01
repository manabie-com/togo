import { Controller, UseInterceptors, Post, Body } from '@nestjs/common';

import { Validation } from './interceptor/validation.interceptor';

import { TodoApplication } from './application/todo.application';
import { CreateTodoDto } from './dto/create-todo.dto';

@Controller('todo')
export class TodoController {
  constructor(private readonly todoApplication: TodoApplication) {}

  @Post()
  @UseInterceptors(Validation)
  async create(@Body() createTodoDto: CreateTodoDto) {
    return this.todoApplication.create(createTodoDto);
  }
}
