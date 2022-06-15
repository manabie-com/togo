import { Body, Controller, Get, Post } from '@nestjs/common';
import * as moment from 'moment';
import { Op } from 'sequelize';
import { LimitTaskService } from '../../../model/limit-task/service/limitTask.service';
import { TaskService } from '../../../model/task/service/task.service';
import { CreateTaskDto } from './dto/createTask.dto';

@Controller('task')
export class TaskController {
  constructor(
    readonly taskService: TaskService,
    readonly limitTaskService: LimitTaskService,
  ) {}

  @Get()
  public getTask() {
    return this.taskService.findAll();
  }

  @Post()
  public async createTask(@Body() body: CreateTaskDto) {
    const { userId } = body;
    const date = new Date();
    const timeHandle = moment(date);
    const startDate = timeHandle.startOf('day').toDate();
    const endDate = timeHandle.endOf('day').toDate();
    // find or create limit task
    const limitTask = await this.limitTaskService.getByUserIdAndDate(
      userId,
      date,
    );
    // count amount task
    const taskNumber = await this.taskService.countNumber({
      userId,
      creationDate: {
        [Op.and]: [{ [Op.gte]: startDate }, { [Op.lte]: endDate }],
      },
    });
    // check limit task and amount task per day
    if (limitTask?.limitNumber < taskNumber)
      return { message: 'The limited tasks in day' };
    return await this.taskService.create(body);
  }
}
