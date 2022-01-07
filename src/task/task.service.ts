import { Injectable } from '@nestjs/common';
import { CreateTaskDto } from './dto/create-task-dto';
import { Task } from './schemas/task.schema';
import { TaskRepository } from './task.repository';

@Injectable()
export class TaskService {

  constructor(private taskRepository: TaskRepository) { }

  async create(createTaskDto: CreateTaskDto): Promise<Task> {
    return this.taskRepository.create(createTaskDto);
  }

  async findAll(): Promise<Task[]> {
    return await this.taskRepository.findAll();
  }
}
