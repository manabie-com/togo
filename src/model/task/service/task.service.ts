import { Inject, Injectable } from '@nestjs/common';
import { WhereOptions } from 'sequelize/types';
import { TASK } from '../../../constance/variable';
import {
  Task,
  TaskAttributes,
  TaskCreationAttributes,
} from '../schema/task.entity';

@Injectable()
export class TaskService {
  constructor(
    @Inject(TASK)
    private taskModel: typeof Task,
  ) {}

  async findAll(): Promise<Task[]> {
    return this.taskModel.findAll();
  }

  async create(data: TaskCreationAttributes) {
    return this.taskModel.create(data);
  }

  async countNumber(where: WhereOptions<TaskAttributes>) {
    return this.taskModel.count({ where });
  }

  async destroyDB() {
    return this.taskModel.destroy({ where: {} });
  }
}
