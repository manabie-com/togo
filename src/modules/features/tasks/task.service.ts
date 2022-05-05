import { BadRequestException, Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/sequelize';
import * as moment from 'moment';
import { Op } from 'sequelize';
import { UserSettingTask } from '../../common/entities';
import { Task } from '../../common/entities/task';
import { CreateMultiTaskRequest } from './request/create-task.dto';

@Injectable()
export class TaskService {
  constructor(
    @InjectModel(Task)
    private taskModel: typeof Task,
    @InjectModel(UserSettingTask)
    private userSettingTask: typeof UserSettingTask
  ) {}

  async createTask(
    userId: string,
    createTaskRequest: CreateMultiTaskRequest
  ): Promise<void> {
    const taskOfUserInDay = await this.userSettingTask.findOne({
      where: { userId },
    });
    if (!taskOfUserInDay) {
      throw new BadRequestException('Not found setting of user');
    }
    const startDate = moment().startOf('day').toDate();
    const endDate = moment().endOf('day').toDate();
    const numberTaskInDate = await this.taskModel.count({
      where: {
        createdBy: userId,
        createdAt: {
          [Op.gte]: startDate,
          [Op.lte]: endDate,
        },
      },
    });
    const tasks = createTaskRequest.tasks;
    if (taskOfUserInDay.get().maximum < numberTaskInDate + tasks.length) {
      throw new BadRequestException(
        'Exceed the number of tasks created per day.'
      );
    }

    const tasksCreate = tasks.map((task) =>
      this.taskModel.create({
        ...task,
        status: 'TO_DO',
        createdBy: userId,
      })
    );
    await Promise.all(tasksCreate);
  }
}
