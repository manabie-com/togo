import { Module } from '@nestjs/common';
import { LimitTaskModule } from 'src/model/limit-task/limitTask.module';
import { TaskModule } from 'src/model/task/task.module';
import { TaskController } from './controller/task/task.controller';

@Module({
  imports: [TaskModule, LimitTaskModule],
  controllers: [TaskController]
})
export class ApiModule {}
