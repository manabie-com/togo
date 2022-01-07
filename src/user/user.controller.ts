import { Body, Controller, Get, Headers, Post } from '@nestjs/common';
import { ApiTags } from '@nestjs/swagger';
import { CreateUserDto } from './dto/create-user-dto';
import { User } from './schemas/user.schema';
import { UserService } from './user.service';

@Controller('users')
@ApiTags('users')
export class UserController {

  constructor(private userService: UserService) { }

  @Post()
  //@UsePipes(ValidationPipe)
  //@UsePipes(new EmployeeTierValidationPipe())
  create(@Body() userCreateDto: CreateUserDto): Promise<User> {
    return this.userService.create(userCreateDto)
  }

  @Get()
  //@UsePipes(ValidationPipe)
  //@UsePipes(new EmployeeTierValidationPipe())
  findAll(): Promise<User[]> {
    return this.userService.findAll();
  }
}
