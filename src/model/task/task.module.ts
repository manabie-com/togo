import { Module } from '@nestjs/common';
import { DatabaseModule } from 'src/database/database.module';
import { taskProviders } from './provider/task.provider';
import { TaskService } from './service/task.service';

@Module({
  imports: [DatabaseModule],
  providers: [
    TaskService,
    ...taskProviders,
  ],
  exports: [TaskService]
})
export class TaskModule {}