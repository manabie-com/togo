import { Body, Controller, Param, Post } from '@nestjs/common'
import { AppService } from './app.service'

@Controller('user')
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Post(':userId/task')
  async takeTask(
    @Param('userId') userId,
    @Body() input: any
  ): Promise<any> {
    return this.appService.takeTask(userId, input)
  }
}
