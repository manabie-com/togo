import { Injectable } from '@nestjs/common';

import { unixTime } from '../../common/utils';
import { Task } from '../../shared/entities/task.entity';
import { User } from '../../shared/entities/user.entity';
import { TaskRepository } from '../../shared/repositories/task.repository';
import { CreateTaskDto, UpdateTaskDto } from './task.dto';

@Injectable()
export class TaskService {
  constructor(private _taskRepository: TaskRepository) {}

  async createNewTask(
    userId: number,
    taskDto: CreateTaskDto
  ): Promise<boolean> {
    await this._taskRepository.save({
      ...taskDto,
      user: new User({ id: userId }),
      createdAt: unixTime(),
      createdBy: userId
    });

    return true;
  }

  getTasks(userId: number): Promise<Task[]> {
    return this._taskRepository.find({
      where: { user: { id: userId } },
      order: { id: 'ASC' }
    });
  }

  async updateTask({ userId, taskId, taskDto }: IUpdateTask): Promise<boolean> {
    const task = await this._taskRepository.findOne({ id: taskId });
    if (!task) {
      return false;
    }

    await this._taskRepository.update(
      { id: taskId },
      {
        ...taskDto,
        updatedAt: unixTime(),
        updatedBy: userId
      }
    );

    return true;
  }
}

interface IUpdateTask {
  userId: number;
  taskId: number;
  taskDto: UpdateTaskDto;
}
