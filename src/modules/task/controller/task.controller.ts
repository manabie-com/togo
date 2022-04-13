import { Body, Controller, Post, UseGuards } from '@nestjs/common';
import { JwtAuthGuard } from '../../../common/guards/auth.guards';
import { UserAccount } from '../../../common/guards/user-account.class';
import { User } from '../../../common/guards/user.decorator';
import { CreateTaskRequestDto } from '../dto/create-task-request.dto';
import { TaskService } from '../service/task.service';

@Controller('v1/task')
export class TaskController {
  constructor(private readonly service: TaskService) {}

  @Post('')
  @UseGuards(JwtAuthGuard)
  async create(
    @Body() request: CreateTaskRequestDto,
    @User() user: UserAccount,
  ) {
    return await this.service.create(request, user.id);
  }
}
