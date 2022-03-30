import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';

import { User } from '../entities/user.entity';
import { Todo } from '../entities/todo.entity';

import { TodoDomain } from '../domain/todo.domain';
import { UserDomain } from '../domain/user.domain';
import { CreateTodoDto } from '../dto/create-todo.dto';

@Injectable()
export class TodoApplication {
  constructor(
    @InjectRepository(User) private readonly userRepository: Repository<User>,
    @InjectRepository(Todo) private readonly todoRepository: Repository<Todo>,
  ) {}
  
  public async create(createTodoDto: CreateTodoDto) {
    const todoDomain = new TodoDomain(this.todoRepository);
    const userDomain = new UserDomain(this.userRepository);
    let userId: number = createTodoDto.user_id;

    if (!userId) {
      const user = await userDomain.create(createTodoDto);
      userId = user.id;
    }

    if (createTodoDto.limit_task) {
      await userDomain.update(createTodoDto, userId);
    }

    return await todoDomain.create({ ...createTodoDto, user_id: userId });
  }
}