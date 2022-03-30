import {
  Inject,
  Injectable,
  HttpStatus,
} from '@nestjs/common';

import { TodoDomain } from '../domain/todo.domain';
import { UserDomain } from '../domain/user.domain';
import { CreateTodoDto } from '../dto/create-todo.dto';

@Injectable()
export class TodoApplication {
  constructor() {}

  public create(createTodoDto: CreateTodoDto) {
    const todoDomain = new TodoDomain();
    const userDomain = new UserDomain();


  }
}