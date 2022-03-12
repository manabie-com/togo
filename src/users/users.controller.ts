import { Controller, Get, Post, Delete, Param, Body } from '@nestjs/common';
import { User } from './users.entity';
import { UserService } from './users.service';
import { CreateUserDto } from '../dto/create-user.dto';

@Controller()
export class UserController {
  constructor(private readonly userService: UserService) {}

  @Post('user')
  async createUser(@Body() user: CreateUserDto) {
    await this.userService.createUser(user);
  }
}
