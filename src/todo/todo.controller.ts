import { Controller, Post, Body} from '@nestjs/common';

import { TodoApplication } from './application/todo.application';
import { CreateTodoDto } from './dto/create-todo.dto';

@Controller('todo')
export class TodoController {
  constructor(private readonly todoApplication: TodoApplication) {}

  @Post()
  create(@Body() createTodoDto: CreateTodoDto) {
    return this.todoApplication.create(createTodoDto);
  }
}
