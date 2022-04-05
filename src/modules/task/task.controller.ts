import {
  Controller,
  Get,
  Param,
  Query,
  Post,
  Body,
  UsePipes,
  Put,
  Delete,
} from '@nestjs/common';
import { TaskService } from './task.service';
import { Task } from './task.entity';
import { TaskQuery, CreateTaskDTO } from './context/index';
import { ResponseInterface } from '@common/context';
import { TaskPipe } from './task.pipe';

@Controller('task')
export class TaskController {
  constructor(private taskService: TaskService) {}

  @Get()
  async findAll(@Query() query: TaskQuery): Promise<Task[]> {
    return this.taskService.findAll(query);
  }

  @Get(':id')
  async findOne(@Param('id') id: number): Promise<Task> {
    return this.taskService.findOne({ id });
  }

  @Post()
  @UsePipes(new TaskPipe())
  async create(
    @Body() context: CreateTaskDTO,
  ): Promise<ResponseInterface<Task>> {
    const data = await this.taskService.create(context);
    return {
      status: 'success',
      data: data,
    };
  }

  @Put(':id')
  @UsePipes(new TaskPipe())
  async update(
    @Param('id') id: number,
    @Body() context: CreateTaskDTO,
  ): Promise<ResponseInterface<Task>> {
    const data = await this.taskService.update(id, context);
    return {
      status: 'success',
      data: data,
    };
  }

  @Delete(':id')
  async delete(@Param('id') id: number): Promise<ResponseInterface<number>> {
    const data = await this.taskService.delete(id);
    return {
      status: 'success',
      data: data,
    };
  }
}
