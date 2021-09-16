import {
  Body,
  Controller,
  Get,
  Param,
  ParseIntPipe,
  Post,
  Put,
  UseGuards
} from '@nestjs/common';
import { ApiOperation, ApiTags } from '@nestjs/swagger';

import { CurrentUser, ICurrentUser } from '../../decorators/user.decorator';
import { AuthGuard } from '../../guards/auth.guard';
import { CreateTaskDto, UpdateTaskDto } from './task.dto';
import { TaskService } from './task.service';

@Controller('tasks')
@ApiTags('Task')
export class TaskController {
  constructor(private _taskService: TaskService) {}

  @Post()
  @ApiOperation({ summary: 'Create a new task' })
  @UseGuards(new AuthGuard('jwt'))
  createTask(
    @CurrentUser() user: ICurrentUser,
    @Body() taskDto: CreateTaskDto
  ): Promise<unknown> {
    return this._taskService.createNewTask(user.id, taskDto);
  }

  @Get()
  @ApiOperation({ summary: 'Get user tasks' })
  @UseGuards(new AuthGuard('jwt'))
  getTasks(@CurrentUser() user: ICurrentUser): Promise<unknown> {
    return this._taskService.getTasks(user.id);
  }

  @Put(':taskId')
  @ApiOperation({ summary: 'Update a task' })
  @UseGuards(new AuthGuard('jwt'))
  updateTask(
    @CurrentUser() user: ICurrentUser,
    @Body() taskDto: UpdateTaskDto,
    @Param('taskId', ParseIntPipe) taskId: number
  ): Promise<boolean> {
    return this._taskService.updateTask({ userId: user.id, taskId, taskDto });
  }
}
