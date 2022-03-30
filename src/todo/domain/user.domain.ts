import { Repository } from 'typeorm';

import { User } from '../entities/user.entity';
import { CreateTodoDto } from '../dto/create-todo.dto';

export class UserDomain {
  constructor(private readonly userRepository: Repository<User>) {}

  public async create(createTodoDto: CreateTodoDto) {
    return await this.userRepository.save({
      limit: createTodoDto.limit_task || null,
    });
  }

  public async update(createTodoDto: CreateTodoDto, userId: number) {
    return await this.userRepository.save({
      id: userId, 
      limit: createTodoDto.limit_task,
    });
  }
}
