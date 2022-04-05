import { Injectable, BadRequestException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { Task } from './task.entity';
import { TaskQuery, TaskFilter, CreateTaskDTO } from './context/index';
import { UserService } from './../user/user.service';
import * as moment from 'moment';

@Injectable()
export class TaskService {
  constructor(
    @InjectRepository(Task)
    private taskRepository: Repository<Task>,
    private userService: UserService,
  ) {}

  async findAll(query: TaskQuery): Promise<Task[]> {
    return this.taskRepository.find({
      where: { ...query.filter },
      order: {
        ...query.orderBy,
      },
      skip: query?.pagination?.offset,
      take: query?.pagination?.limit,
    });
  }

  async findOne(filter: TaskFilter): Promise<Task> {
    return this.taskRepository.findOne({ where: { ...filter } });
  }

  async create(context: CreateTaskDTO): Promise<Task> {
    const task = Object.assign(new Task(), context);
    const user = await this.userService.findOne({ id: context.assignee_id });
    if (!user) {
      throw new BadRequestException({
        message: 'User not found',
      });
    }
    const totalTaskInDay = await this.totalTaskInDay(context);
    if (totalTaskInDay >= user.limitTaskInDay) {
      throw new BadRequestException({
        message: 'Total task in day must be less than limit task per user',
      });
    }
    // task.code =
    //   'TASK-' + user.id + '/' + (Math.random() + 1).toString(36).substring(8);
    task.updated_time = new Date();
    task.created_time = new Date();
    return this.taskRepository.save(task);
  }

  async totalTaskInDay(context: CreateTaskDTO): Promise<number> {
    let { totalTaskInDay } = await this.taskRepository
      .createQueryBuilder()
      .select('count(id)', 'totalTaskInDay')
      .where('assignee_id = :assignee_id', { assignee_id: context.assignee_id })
      .where('created_time >= :created_time', {
        created_time: moment().format('YYYY-MM-DD 00:00:00'),
      })
      .where('created_time < :created_time', {
        created_time: moment().format('YYYY-MM-DD 23:59:59'),
      })
      .getRawOne<{ totalTaskInDay: number }>();
    return totalTaskInDay || 0;
  }

  async update(id: number, context: CreateTaskDTO): Promise<Task> {
    const user = await this.userService.findOne({ id: context.assignee_id });
    if (!user) {
      throw new BadRequestException({
        message: 'User not found',
      });
    }
    const task = await this.taskRepository.findOne({ where: { id } });
    if (!task) {
      throw new BadRequestException({
        message: 'Task does not exist',
      });
    }
    Object.assign(task, context, { updated_time: new Date() });
    return this.taskRepository.save(task);
  }

  async delete(id: number): Promise<number> {
    const result = await this.taskRepository.delete({ id });
    return result.affected || 0;
  }
}
