import { Body, Controller, Get, Headers, Post } from '@nestjs/common';
import { CreateTaskDto } from './dto/create-task-dto';
import { Task } from './schemas/task.schema';
import { TaskService } from './task.service';

@Controller('task')
export class TaskController {

  constructor(private taskService: TaskService) { }

  @Post()
  //@UsePipes(ValidationPipe)
  //@UsePipes(new EmployeeTierValidationPipe())
  create(@Body() taskCreateDto: CreateTaskDto): Promise<Task> {
    return this.taskService.create(taskCreateDto)
  }

  @Get()
  //@UsePipes(ValidationPipe)
  //@UsePipes(new EmployeeTierValidationPipe())
  gettAll(): Promise<Task[]> {
    return this.taskService.getAll();
  }
}
