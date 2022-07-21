import { BadRequestException, Body, Controller, Delete, Get, Param, Post, Put } from '@nestjs/common';
import { ApiOkResponse, ApiTags } from '@nestjs/swagger';
import { UsersService } from 'src/users/users.service';
import { CreateTodoDto } from './dto/create-todo.dto';
import { TodoDto } from './dto/todo.dto';
import { UpdateTodoDto } from './dto/update-todo.dto';
import { TodosService } from './todos.service';

@ApiTags('Todos')
@Controller('users/:userId/todos')
export class TodosController {
  constructor(
    private readonly usersService: UsersService,
    private readonly todosService: TodosService) { }

  @ApiOkResponse({ type: TodoDto })
  @Post()
  async create(@Param('userId') userId: string, @Body() createTodoDto: CreateTodoDto) {
    await this.usersService.findOne(userId);
    return this.todosService.create({ ...createTodoDto, user: { id: userId } });
  }

  @ApiOkResponse({ type: TodoDto, isArray: true })
  @Get()
  async findAll(@Param('userId') userId: string) {
    await this.usersService.findOne(userId);
    return this.todosService.findByUserId(userId);
  }

  @ApiOkResponse({ type: TodoDto })
  @Get(':id')
  async findOne(@Param('userId') userId: string, @Param('id') id: string) {
    const todo = await this.todosService.findOne(id);

    if (todo.user.id !== userId) {
      throw new BadRequestException('Todo not belong to user')
    }

    return todo;
  }

  @ApiOkResponse({ type: TodoDto })
  @Put(':id')
  async update(@Param('userId') userId: string, @Param('id') id: string, @Body() updateTodoDto: UpdateTodoDto) {
    const todo = await this.todosService.findOne(id);

    if (todo.user.id !== userId) {
      throw new BadRequestException('Todo not belong to user')
    }

    return this.todosService.update(id, updateTodoDto);
  }

  @ApiOkResponse({ type: TodoDto })
  @Delete(':id')
  async remove(@Param('userId') userId: string, @Param('id') id: string) {
    const todo = await this.todosService.findOne(id);

    if (todo.user.id !== userId) {
      throw new BadRequestException('Todo not belong to user')
    }

    return this.todosService.remove(id);
  }
}
