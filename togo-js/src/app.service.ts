import { HttpException, HttpStatus, Injectable } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { TaskInput } from './dto/task.create';
import { Task } from './models/task.entity';
import { User } from './models/user.entity';
import { Raw } from 'typeorm';

@Injectable()
export class AppService {
  constructor(private jwtService: JwtService) {}

  async checkDailyTask(id: number) {
    const user = await User.findOne(id, {
      select: ['maxTodo']
    });
    const amountDailyTask = await Task.count({
      where: {
        user: {
          id,
        },
        createdAt: Raw((alias) => `${alias} >= NOW() - INTERVAL '24 HOURS'`),
      },
    });
    if (amountDailyTask >= user.maxTodo)
      throw new HttpException('DAILY_TASK_EXCEEDED', HttpStatus.BAD_REQUEST);
  }

  async createTask(id: number, input: TaskInput) {
    const user = await User.findOne(id);

    // we can just pass in like { ...input }
    // but we will have to make a global validation and transformation pipe first
    // i don't have time for that T_T
    const task = Task.create({
      content: input.content || '',
      user: user,
    });
    await Task.save(task);
    return {
      data: task,
    };
  }

  async listTasks(id: number) {
    // i was considering adding index to task.userId column
    // but it's not worth it
    const tasks = await Task.find({
      where: {
        user: {
          id,
        },
      },
    });
    return { data: tasks };
  }

  async login(userId: string, password: string) {
    const user = await User.findOne({
      where: {
        userId,
        password,
      },
    });
    const payload = this.jwtService.sign({
      id: user.id,
      user_name: user.userId,
    });
    return {
      data: payload,
    };
  }
}
