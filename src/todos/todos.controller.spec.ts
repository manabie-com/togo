import { Test, TestingModule } from '@nestjs/testing';
import { UsersService } from 'src/users/users.service';
import { Repository } from 'typeorm';
import { TodoEntity } from './entities/todo.entity';
import { TodosController } from './todos.controller';
import { TodosService } from './todos.service';

describe('TodosController', () => {
  let controller: TodosController;
  let todosRepository: Repository<TodoEntity>;
  let usersService: UsersService;

  beforeEach(async () => {
    todosRepository = jest.mock as any;
    usersService = jest.mock as any;

    const module: TestingModule = await Test.createTestingModule({
      controllers: [TodosController],
      providers: [
        {
          provide: UsersService,
          useValue: usersService
        },
        {
          provide: 'TodoEntityRepository',
          useValue: todosRepository
        },
        TodosService],
    }).compile();

    controller = module.get<TodosController>(TodosController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
