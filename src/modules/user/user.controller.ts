import {
  Controller,
  Get,
  Param,
  Post,
  Body,
  Put,
  Delete,
  UsePipes,
} from '@nestjs/common';
import { UserService } from './user.service';
import { User } from './user.entity';
import { UpsertUserDTO } from './context/index';
import { ResponseInterface } from '@common/context';
import { UserPipe } from './user.pipe';

@Controller('user')
export class UserController {
  constructor(private userService: UserService) {}

  @Get(':id')
  async findOne(@Param('id') id: number): Promise<User> {
    return this.userService.findOne({ id });
  }

  @Post()
  @UsePipes(new UserPipe())
  async create(
    @Body() context: UpsertUserDTO,
  ): Promise<ResponseInterface<User>> {
    const data = await this.userService.create(context);
    return {
      message: 'Success',
      data: data,
    };
  }

  @Put(':id')
  async update(
    @Param('id') id: number,
    @Body() context: UpsertUserDTO,
  ): Promise<ResponseInterface<User>> {
    const data = await this.userService.update(id, context);
    return {
      status: 'success',
      data: data,
    };
  }

  @Delete(':id')
  async delete(@Param('id') id: number): Promise<ResponseInterface<number>> {
    const data = await this.userService.delete(id);
    return {
      status: 'success',
      data: data,
    };
  }
}
