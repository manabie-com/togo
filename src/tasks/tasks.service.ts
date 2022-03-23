import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { TaskInterface } from './interfaces/task.interface';
import { Task, TaskDocument } from './schemas/task.schema';
import * as moment from 'moment';
import { UsersService } from '../users/users.service';

@Injectable()
export class TasksService {
  constructor(
    @InjectModel(Task.name) private readonly taskModel: Model<TaskDocument>,
    private readonly usersService: UsersService,
  ) {}

  create(data: TaskInterface) {
    return this.taskModel.create(data);
  }

  async getAllByOwner(owner: string, page: number, size: number) {
    const data = await this.taskModel.aggregate([
      {
        $match: {
          owner,
        },
      },
      {
        $sort: {
          createdAt: -1,
        },
      },
      {
        $facet: {
          data: [{ $skip: (page - 1) * size }, { $limit: size }],
          count: [{ $count: 'totalRecord' }],
        },
      },
      {
        $unwind: {
          path: '$count',
        },
      },
      {
        $project: {
          data: 1,
          total: '$count.totalRecord',
        },
      },
    ]);

    if (data?.length) {
      return data[0];
    }

    return null;
  }

  countTasks(owner: string, fromDate: Date, toDate: Date) {
    const startOfDay = moment(fromDate).startOf('day');
    const endOfDay = moment(toDate).endOf('day');
    return this.taskModel.countDocuments({
      $and: [
        {
          owner,
        },
        { createdAt: { $gte: startOfDay } },
        { createdAt: { $lte: endOfDay } },
      ],
    });
  }

  async reachLimitedTaskPerDay(owner: string) {
    const user = await this.usersService.findOneById(owner);
    const today = new Date();
    const todayTasks = await this.countTasks(owner, today, today);

    if (todayTasks >= user.limitedTaskPerDay) {
      return true;
    }
    return false;
  }
}
