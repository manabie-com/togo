import { Injectable } from '@nestjs/common';
import { v1 as uuid } from 'uuid'

import { CreateTaskDto } from './dto/create-task-dto';
import { Task } from './schemas/task.schema';

@Injectable()
export class TaskRepository {

  private tasks: Task[] = [
    {
      id: "00ca6888-af27-4f6f-95bc-4735c326447d",
      title: "task1",
      userId: "305c8624-a214-4a25-93b4-d54dcc411150",
      content: "content1",
      dateTime: new Date(),
      createdAt: new Date(),
      updatedAt: new Date(),
    },
    {
      id: "fa1ef149-be12-4af9-a76b-e9ce9e7b0d37",
      title: "task2",
      userId: "305c8624-a214-4a25-93b4-d54dcc411150",
      content: "content2",
      dateTime: new Date(),
      createdAt: new Date(),
      updatedAt: new Date(),
    },
    {
      id: "fd38518d-c4bc-492a-b9e6-785d40bb3fd2",
      title: "task3",
      userId: "d864710e-c343-413d-8885-d2d0053fa75b",
      content: "content3",
      dateTime: new Date(),
      createdAt: new Date(),
      updatedAt: new Date(),
    }
  ];

  async create(createTaskDto: CreateTaskDto): Promise<Task> {
    const { title, content, dateTime, userId } = createTaskDto;
    const task = {
      id: uuid(),
      userId,
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
