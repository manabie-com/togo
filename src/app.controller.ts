import { Body, Controller, Param, ParseIntPipe, Post } from '@nestjs/common'
import { AppService } from './app.service'

import { TakeTaskDTO } from './dto/dto'

@Controller('user')
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Post(':userId/task')
  async takeTask(
    @Param('userId', ParseIntPipe) userId: number,
    @Body() input: TakeTaskDTO,
  ): Promise<any> {
    return await this.appService.takeTask(userId, input)
  }
}
