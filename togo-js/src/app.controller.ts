import { Body, Controller, Get, Post, Query, UseGuards } from '@nestjs/common';
import { AppService } from './app.service';
import { TaskInput } from './dto/task.create';
import { JwtAuthGuard } from './services/jwt/auth.guard';
import { GetUser } from './services/jwt/get-user.decorator';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get('/login')
  login(@Query('user_id') userId: string, @Query('password') passport: string) {
    return this.appService.login(userId, passport);
  }

  @Post('/tasks')
  @UseGuards(JwtAuthGuard)
  async createTask(@GetUser() id: number, @Body() input: TaskInput) {
    await this.appService.checkDailyTask(id);
    const result = await this.appService.createTask(id, input);
    return result;
  }

  @Get('/tasks')
  @UseGuards(JwtAuthGuard)
  async getTasks(@GetUser() id: number) {
    const result = await this.appService.listTasks(id);
    return result;
  }
}
