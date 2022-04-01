import {
  Body,
  Controller,
  Get,
  NotFoundException,
  Param,
  ParseIntPipe,
  Post,
  Put,
  Query,
} from '@nestjs/common';
import { ApiOperation, ApiTags } from '@nestjs/swagger';
import { ValidationPipe } from 'src/common/pipes/validation.pipe';
import {
  ApiPaginatedResponse,
  CreateUserDto,
  GetUserListDto,
  UpdateUserDto,
} from 'src/dto';
import { ApiResponse } from 'src/dto/ApiReponse.dto';
import { User } from 'src/entities/user.entity';
import { UserService } from 'src/services/user.service';

@Controller('users')
@ApiTags('users')
export class UserController {
  constructor(private userService: UserService) {}

  @Get()
  @ApiOperation({ summary: 'Get user list' })
  @ApiPaginatedResponse(User)
  getUserList(
    @Query(new ValidationPipe()) query: GetUserListDto,
  ): Promise<{ items: User[]; total: number }> {
    return this.userService.find(query);
  }

  @Get(':userId')
  @ApiOperation({ summary: 'Get user info' })
  @ApiResponse(User)
  async getUser(
    @Param('userId', new ParseIntPipe()) id: number,
  ): Promise<User> {
    const user = await this.userService.findOne(id);
    if (!user) throw new NotFoundException('User not found!');
    return user;
  }

  @Put(':userId')
  @ApiOperation({ summary: 'update user' })
  @ApiResponse(User)
  updateUser(
    @Param('userId', new ParseIntPipe()) id: number,
    @Body(new ValidationPipe()) body: UpdateUserDto,
  ) {
    return this.userService.updateUser(id, body);
  }

  @Post()
  @ApiOperation({ summary: 'create user' })
  @ApiResponse(User)
  createUser(@Body(new ValidationPipe()) body: CreateUserDto) {
    return this.userService.createUser(body);
  }
}
