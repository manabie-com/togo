import { Injectable } from '@nestjs/common';
import { v1 as uuid } from 'uuid'

import { CreateTaskDto } from './dto/create-task-dto';
import { Task } from './schemas/task.schema';

@Injectable()
export class TaskRepository {

  private tasks: Task[] = [];

  async create(createTaskDto: CreateTaskDto): Promise<Task> {
    const { title,
        content,
        dateTime,
        } = createTaskDto;
    const task = {
      id: uuid(),
      title,
      content,
      dateTime,
      createdAt: new Date(),
      updatedAt: new Date(),
    }
    this.tasks.push(task)
    return task;
  }

  async findAll(): Promise<Task[]> {
    return this.tasks;
  }

}
