import { Module } from '@nestjs/common';
import { TaskService } from './task.service';
import { TaskController } from './task.controller';
import { TypeOrmModule } from '@nestjs/typeorm';
import { TaskRepository } from './task.repository';
import { UserModule } from '../user/user.module';
import { CurrentUserInterceptor } from '../../interceptor/user.interceptor';

@Module({
  imports: [
    UserModule,
    TypeOrmModule.forFeature([ TaskRepository ])
  ],
  controllers: [TaskController],
  providers: [TaskService, CurrentUserInterceptor]
})
export class TaskModule {}
