import { Module } from '@nestjs/common';
import { SequelizeModule } from '@nestjs/sequelize';
import { Task } from 'src/modules/common/entities';
import { TaskController } from './task.controller';
import { TaskService } from './task.service';

@Module({
  imports: [SequelizeModule.forFeature([Task])],
  providers: [TaskService],
  controllers: [TaskController],
})
export class TaskModule {}
