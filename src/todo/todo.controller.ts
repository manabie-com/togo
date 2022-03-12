import {
  Controller,
  Get,
  Post,
  Put,
  Delete,
  Request,
  Body,
  Param,
  Query,
  UseGuards,
} from '@nestjs/common';
import { JwtAuthGuard } from '../auth/jwt-auth.guard';
import { Todo, TodoStatus } from './todo.entity';
import { TodoService } from './todo.service';
import { UserService } from '../users/users.service';

@Controller()
export class TodoController {
  constructor(
    private readonly todoService: TodoService,
    private readonly userService: UserService,
  ) {}

  @UseGuards(JwtAuthGuard)
  @Get('tasks?')
  async getAllTodo(
    @Request() req,
    @Query('status') status: TodoStatus,
  ): Promise<{ data: Todo[] }> {
    return { data: await this.todoService.getTodos(req.user.username, status) };
  }

  @UseGuards(JwtAuthGuard)
  @Post('tasks')
  async createTodo(
    @Request() req,
    @Body('content') content: string,
  ): Promise<{ data: Todo }> {
    const username = req.user.username;
    const todoCount = await this.todoService.countTodoDaily(username);
    return {
      data: await this.userService.createTodo(username, content, todoCount),
    };
  }

  @UseGuards(JwtAuthGuard)
  @Put('tasks')
  async updateStatus(@Body() body: { id: string; status: TodoStatus }) {
    await this.todoService.setTodoStatus(body.id, body.status);
  }

  @UseGuards(JwtAuthGuard)
  @Put('many-tasks')
  async updateManyStatus(@Body() body: { ids: string[]; status: TodoStatus }) {
    await this.todoService.setManyTodoStatus(body.ids, body.status);
  }

  @UseGuards(JwtAuthGuard)
  @Delete('tasks/:id')
  async deleteTodo(@Param('id') id: string) {
    await this.todoService.deleteTodoById(id);
  }

  @UseGuards(JwtAuthGuard)
  @Delete('tasks')
  async deleteAllTodos(@Request() req) {
    await this.todoService.deleteAllTodos(req.user.username);
  }
}
