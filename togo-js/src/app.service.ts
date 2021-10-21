import { HttpException, HttpStatus, Injectable } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { TaskInput } from './dto/task.create';
import { Task } from './models/task.entity';
import { User } from './models/user.entity';
import { Raw } from 'typeorm';
// import { InjectRepository } from '@nestjs/typeorm';
// import { Brackets, Connection, getManager, IsNull, Repository } from 'typeorm';
// import { JwtPayload } from './services/jwt/strategy';

@Injectable()
export class AppService {
  constructor(private jwtService: JwtService) {}

  async checkDailyTask(id: number) {
    const user = await User.findOne(id);
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
