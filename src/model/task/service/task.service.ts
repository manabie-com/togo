import { Inject, Injectable } from '@nestjs/common';
import { WhereOptions } from 'sequelize/types';
import { TASK } from 'src/constance/variable';
import { Task, TaskAttributes, TaskCreationAttributes } from '../schema/task.entity';

@Injectable()
export class TaskService {
  constructor(
    @Inject(TASK)
    private userModel: typeof Task
    ) {}

  async findAll(): Promise<Task[]> {
    return this.userModel.findAll()
  }

  async create(data: TaskCreationAttributes) {
    return this.userModel.create(data)
  }

  async countNumber(where: WhereOptions<TaskAttributes>) {
    return this.userModel.count({ where })
  }
}