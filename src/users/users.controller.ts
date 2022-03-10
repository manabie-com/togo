import { Controller, Get, Post, Delete, Param, Body } from '@nestjs/common';
import { User } from './users.entity.js';
import { UserService } from './users.service.js';

@Controller()
export class UserController {
  constructor(private readonly userService: UserService) {}

  @Post('user')
  async createUser(@Body() user: User) {
    await this.userService.createUser(user);
  }
}
