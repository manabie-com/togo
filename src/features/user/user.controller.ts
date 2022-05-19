import { Controller, Post, Body, UseGuards } from '@nestjs/common';
import { UserService } from './user.service';
import { CreateUserDto } from './dto/create-user.dto';
import { AdminRoleGuard } from '../../guards/role.guard';


@Controller('users')
export class UserController {
  constructor(private readonly userService: UserService) {}

  @UseGuards(AdminRoleGuard)
  @Post()
  async create(@Body() createTaskDto: CreateUserDto) {
    return await this.userService.create(createTaskDto);
  }
}
