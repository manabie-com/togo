import { Module } from '@nestjs/common';
import { LimitTaskModule } from '../model/limit-task/limitTask.module';
import { TaskModule } from '../model/task/task.module';
import { TaskController } from './controller/task/task.controller';

@Module({
  imports: [TaskModule, LimitTaskModule],
  controllers: [TaskController],
})
export class ApiModule {}
