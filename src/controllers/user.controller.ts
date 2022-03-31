import { Controller, Get, Query, UsePipes } from '@nestjs/common';
import { ValidationPipe } from 'src/common/pipes/validation.pipe';
import { GetUserListDto } from 'src/dto';

@Controller('users')
export class UserController {
  @Get()
  getUserList(@Query(new ValidationPipe()) query: GetUserListDto) {
    return query;
  }
}
