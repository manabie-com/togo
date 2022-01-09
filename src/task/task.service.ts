import { Injectable } from '@nestjs/common';

import { UserService } from '../user/user.service';
import { CreateTaskDto } from './dto/create-task-dto';
import { Task } from './schemas/task.schema';
import { TaskRepository } from './task.repository';

@Injectable()
export class TaskService {

  constructor(
    private taskRepository: TaskRepository,
    private userService: UserService
  ) { }

  async create(createTaskDto: CreateTaskDto): Promise<Task> {
    await this.userService.incrementTask(createTaskDto.userId);
    return this.taskRepository.create(createTaskDto);
  }

  async findAll(): Promise<Task[]> {
    return await this.taskRepository.findAll();
  }
}
