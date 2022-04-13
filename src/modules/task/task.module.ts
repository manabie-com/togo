import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { UserModule } from '../user/user.module';
import { TaskController } from './controller/task.controller';
import { TaskEntity } from './entity/task.entity';
import { TaskService } from './service/task.service';

@Module({
  imports: [UserModule, TypeOrmModule.forFeature([TaskEntity])],
  providers: [TaskService],
  controllers: [TaskController],
  exports: [TaskService],
})
export class TaskModule {}
