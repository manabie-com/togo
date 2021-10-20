import { Controller, Get, Post, Query } from '@nestjs/common';
import { AppService } from './app.service';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get('/login')
  login(
    @Query('user_id') userId: string,
    @Query('password') passport: string,
  ) {
    return this.appService.login(userId, passport);
  }

  @Get('/tasks')
  getTasks(): string {
    return 'tasks';
  }

  @Post('/tasks')
  createTask() {
    return 'new task';
  }
}
