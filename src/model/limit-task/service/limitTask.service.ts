import { Injectable, Inject } from '@nestjs/common';
import { LIMIT_TASK } from 'src/constance/variable';
import { LimitTask } from '../schema/limitTask.entity';

@Injectable()
export class LimitTaskService {
  constructor(
    @Inject(LIMIT_TASK)
    private limitTaskModel: typeof LimitTask
  ) {}

  async findAll(): Promise<LimitTask[]> {
    return this.limitTaskModel.findAll<LimitTask>();
  }

  async getByUserIdAndDate(userId: string, date: Date) {
    const result = await this.limitTaskModel.findOne({ where: { userId, creationDate: date } });
    if(!result) return await this.limitTaskModel.create({ userId, creationDate: date, limitNumber: 5 });
  }
}