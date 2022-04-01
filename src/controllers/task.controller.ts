import {
  Body,
  Controller,
  Delete,
  Get,
  NotFoundException,
  Param,
  ParseIntPipe,
  Post,
  Put,
  Query,
} from '@nestjs/common';
import { ApiOperation, ApiTags } from '@nestjs/swagger';
import { ValidationPipe } from 'src/common/pipes/validation.pipe';
import {
  ApiPaginatedResponse,
  CreateTaskDto,
  CreateToDoDto,
  GetTasksDto,
  UpdateTaskDto,
  UpdateToDoDto,
} from 'src/dto';
import { ApiResponse } from 'src/dto/ApiReponse.dto';
import { Task } from 'src/entities/task.entity';
import { ToDoList } from 'src/entities/toDoList.entity';
import { TaskService } from 'src/services/task.service';
import { ToDoService } from 'src/services/todo.service';

@Controller('tasks')
@ApiTags('tasks')
export class TaskController {
  constructor(
    private taskService: TaskService,
    private todoService: ToDoService,
  ) {}

  @Get()
  @ApiOperation({ summary: 'Get task list' })
  @ApiPaginatedResponse(Task)
  getUserList(@Query(new ValidationPipe()) query: GetTasksDto) {
    return this.taskService.find(query);
  }

  @Get(':taskId')
  @ApiOperation({ summary: 'Get task info' })
  @ApiResponse(Task)
  async getTask(@Param('taskId', new ParseIntPipe()) id: number) {
    const task = await this.taskService.findOne(id, ['toDoList', 'user']);
    if (!task) throw new NotFoundException('Task not found!');
    return task;
  }

  @Post()
  @ApiOperation({ summary: 'Create task' })
  @ApiResponse(Task)
  createUser(@Body(new ValidationPipe()) body: CreateTaskDto) {
    return this.taskService.create(body);
  }

  @Put(':taskId')
  @ApiOperation({ summary: 'Update task' })
  @ApiResponse(Task)
  async updateTask(
    @Param('taskId', new ParseIntPipe()) id: number,
    @Body(new ValidationPipe()) body: UpdateTaskDto,
  ) {
    return this.taskService.update(id, body);
  }

  @Delete(':taskId')
  @ApiOperation({ summary: 'Delete a task' })
  @ApiResponse(ToDoList)
  async deleteTask(@Param('taskId', new ParseIntPipe()) id: number) {
    const task = await this.taskService.findOne(id);
    if (!task) throw new NotFoundException('Task not found!');

    await this.taskService.delete(id);
    return this.taskService.findOne(id);
  }

  @Post(':taskId/to-do')
  @ApiOperation({ summary: 'Create todo for a task' })
  @ApiResponse(ToDoList)
  async createToDo(
    @Param('taskId', new ParseIntPipe()) id: number,
    @Body(new ValidationPipe()) body: CreateToDoDto,
  ) {
    const task = await this.taskService.findOne(id);
    if (!task) throw new NotFoundException('Task not found!');

    return this.todoService.create(task, body);
  }

  @Put(':taskId/to-do/:todoId')
  @ApiOperation({ summary: 'update a todo' })
  @ApiResponse(ToDoList)
  async updateToDo(
    @Param('taskId', new ParseIntPipe()) taskId: number,
    @Param('todoId', new ParseIntPipe()) todoId: number,
    @Body(new ValidationPipe()) body: UpdateToDoDto,
  ) {
    const task = await this.taskService.findOne(taskId);
    if (!task) throw new NotFoundException('Task not found!');

    return this.todoService.update(taskId, todoId, body);
  }
}
