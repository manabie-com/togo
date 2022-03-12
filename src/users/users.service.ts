import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import bcrypt from 'bcrypt';
import { saltRounds } from '../auth/constants';
import { User } from './users.entity';
import { Todo } from '../todo/todo.entity';
import { CreateUserDto } from '../dto/create-user.dto';

@Injectable()
export class UserService {
  constructor(
    @InjectRepository(User)
    private usersRepository: Repository<User>,
  ) {}

  async createUser(user: CreateUserDto) {
    if (await this.usersRepository.findOne(user.username))
      throw new HttpException({
        status: HttpStatus.BAD_REQUEST,
        error: 'Username already existed',
      }, HttpStatus.BAD_REQUEST);
    else {
      user.password = await bcrypt.hash(user.password, saltRounds);
      user = this.usersRepository.create(user);
      await this.usersRepository.save(user);
    }
  }

  async createTodo(username: string, content: string, todoCount: number): Promise<Todo> {
    const user = await this.findUser(username);

    if (todoCount >= user.limitPerDay)
      throw new HttpException({
        status: HttpStatus.BAD_REQUEST,
        error: "Daily todo's limit exceeded",
      }, HttpStatus.BAD_REQUEST);

    const todo = new Todo();
    todo.content = content;
    user.todos.push(todo);
    await this.usersRepository.save(user);
    return todo;
  }

  async findUser(username: string): Promise<User> {
    return this.usersRepository.findOne({
      where: { username: username },
      relations: ['todos'],
    });
  }
}
