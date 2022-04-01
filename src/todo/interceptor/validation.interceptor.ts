import {
  Injectable,
  NestInterceptor,
  ExecutionContext,
  CallHandler,
  BadRequestException,
} from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository, Between } from 'typeorm';

import { startOfDay, endOfDay } from 'date-fns';

import { User } from '../entities/user.entity';
import { Todo } from '../entities/todo.entity';

@Injectable()
export class Validation implements NestInterceptor {
  constructor(
    @InjectRepository(User) private readonly userRepository: Repository<User>,
    @InjectRepository(Todo) private readonly todoRepository: Repository<Todo>,
  ) {}

  async intercept(
    context: ExecutionContext,
    next: CallHandler,
  ) {
    const httpContext = context.switchToHttp();
    const req = httpContext.getRequest();

    const userId: number = req.body.user_id;
    if (!userId && userId !== 0) {
      return next.handle();
    }

    const user = await this.userRepository.findOneBy({ 
      id: userId
    });

    if (!user) {
      throw new BadRequestException('User not exist');
    }

    const nowDate = (date: Date) => Between(startOfDay(date), endOfDay(date));
    const tasks: Todo[] = await this.todoRepository.find({
      where: {
        userId: userId,
        createdAt: nowDate(new Date()),
      }
    });

    const limit: number = req.body.limit_task || user.limit;
    const quantityTask = tasks.length + 1;
    if (limit != null && quantityTask > limit) {
      throw new BadRequestException('Overload task in 1 day');
    }
 
    return next.handle();
  }
}
