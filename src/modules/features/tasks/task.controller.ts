import { Body, Controller, Post } from '@nestjs/common';
import { CreateMultiTaskRequest } from './request/create-task.dto';
import { TaskService } from './task.service';

@Controller('tasks')
export class TaskController {
  constructor(private taskService: TaskService) {}

  @Post()
  create(@Body() createTasksRequest: CreateMultiTaskRequest): Promise<void> {
    return this.taskService.createTask('', createTasksRequest);
  }
}
