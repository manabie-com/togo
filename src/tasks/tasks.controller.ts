import {
  Controller,
  Post,
  Body,
  UseGuards,
  Request,
  BadRequestException,
  Get,
  Query,
} from '@nestjs/common';
import { TasksService } from './tasks.service';
import { CreateTaskDto } from './dto/create-task.dto';
import { JwtAuthGuard } from 'src/auth/guards/jwt-auth.guard';
import { PaginationQueryDto } from './dto/pagination-query.dto';

@Controller('tasks')
export class TasksController {
  constructor(private readonly tasksService: TasksService) {}

  @UseGuards(JwtAuthGuard)
  @Get()
  async findAll(@Request() req, @Query() params: PaginationQueryDto) {
    const data = await this.tasksService.getAllByOwner(
      req.user.id,
      params.page,
      params.size,
    );
    const response = {
      data: [],
      page: params.page,
      size: params.size,
      total: 0,
    };

    if (data) {
      response['size'] = data.data.length;
      response['data'] = data.data;
      response['total'] = data.total;
    }

    return response;
  }

  @UseGuards(JwtAuthGuard)
  @Post()
  async create(@Request() req, @Body() createTaskDto: CreateTaskDto) {
    const data = {
      owner: req.user.id,
      description: createTaskDto.description,
    };
    const reachLimited = await this.tasksService.reachLimitedTaskPerDay(
      data.owner,
    );
    if (reachLimited) {
      throw new BadRequestException('reach limited task per day');
    } else {
      return await this.tasksService.create(data);
    }
  }
}
