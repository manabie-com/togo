import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/sequelize';
import { Task } from '../../common/entities/task';
import { CreateMultiTaskRequest } from './request/create-task.dto';

@Injectable()
export class TaskService {
  constructor(
    @InjectModel(Task)
    private taskModel: typeof Task
  ) {}

  async createTask(
    userId: string,
    createTaskRequest: CreateMultiTaskRequest
  ): Promise<void> {
    const tasksCreate = createTaskRequest.tasks.map((task) =>
      this.taskModel.create({
        ...task,
        status: 'TO_DO',
        createdBy: userId,
      })
    );
    await Promise.all(tasksCreate);
  }
}
