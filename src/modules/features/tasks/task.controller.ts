import { Body, Controller, Post, UseGuards } from '@nestjs/common';
import { ApiBearerAuth, ApiTags } from '@nestjs/swagger';
import { JwtAuthGuard } from '../auth/guards/jwt-auth.guard';
import { SessionUser } from '../auth/guards/user.session';
import { CreateMultiTaskRequest } from './request/create-task.dto';
import { TaskService } from './task.service';

@Controller('tasks')
@ApiTags('tasks')
@ApiBearerAuth()
export class TaskController {
  constructor(private taskService: TaskService) {}

  @Post()
  @UseGuards(JwtAuthGuard)
  create(
    @SessionUser() user: any,
    @Body() createTasksRequest: CreateMultiTaskRequest
  ): Promise<void> {
    return this.taskService.createTask(user.id, createTasksRequest);
  }
}
