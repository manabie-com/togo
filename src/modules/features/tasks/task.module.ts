import { Module } from '@nestjs/common';
import { SequelizeModule } from '@nestjs/sequelize';
import { Task, UserSettingTask } from 'src/modules/common/entities';
import { TaskController } from './task.controller';
import { TaskService } from './task.service';

@Module({
  imports: [SequelizeModule.forFeature([Task, UserSettingTask])],
  providers: [TaskService],
  controllers: [TaskController],
})
export class TaskModule {}
