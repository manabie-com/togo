import { Injectable, Inject } from '@nestjs/common';
import { LIMIT_TASK } from '../../../constance/variable';
import { LimitTask } from '../schema/limitTask.entity';
import * as moment from 'moment';
import { Op } from 'sequelize';

@Injectable()
export class LimitTaskService {
  constructor(
    @Inject(LIMIT_TASK)
    private limitTaskModel: typeof LimitTask,
  ) {}

  async findAll(): Promise<LimitTask[]> {
    return this.limitTaskModel.findAll<LimitTask>();
  }

  async getByUserIdAndDate(userId: string, date: Date) {
    const timeHandle = moment(date);
    const startDate = timeHandle.startOf('day').toDate();
    const endDate = timeHandle.endOf('day').toDate();
    const result = await this.limitTaskModel.findOne({
      where: {
        userId,
        creationDate: {
          [Op.and]: [{ [Op.gte]: startDate }, { [Op.lte]: endDate }],
        },
      },
    });
    if (!result)
      return await this.limitTaskModel.create({
        userId,
        creationDate: date,
        limitNumber: 5,
      });
    return result;
  }
}
