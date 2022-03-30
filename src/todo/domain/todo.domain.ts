import { Repository } from 'typeorm';

import { Todo } from '../entities/todo.entity';
import { CreateTodoDto } from '../dto/create-todo.dto';

export class TodoDomain {
  constructor(private readonly todoRepository: Repository<Todo>) {}

  public async create(createTodoDto: CreateTodoDto) {
    return await this.todoRepository.save({
      task: createTodoDto.task,
      userId: createTodoDto.user_id,
    });
  }
}
