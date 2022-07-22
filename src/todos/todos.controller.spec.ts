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
  let todosService: TodosService;
  let trxMock: any;
  const createTodoDto = {
    title: "todo1",
    date: new Date(),
    user: {
      id: "userid1"
    }
  }
  const creatingTodo = {
    ...createTodoDto,
    id: '123'
  }

  const user = {
    id: 'userid1'
  }

  beforeEach(async () => {
    trxMock = {
      findAndCount: jest.fn(() => [{}, 0]),
      findOne: jest.fn(() => ({ todoPerday: 1 })),
      save: jest.fn()
    }

    todosRepository = {
      create: jest.fn(() => creatingTodo),
      save: jest.fn(),
      manager: {
        transaction: async (type, cb) => {
          await cb(trxMock);
          return ({ catch: jest.fn() })
        }
      }
    } as any;

    usersService = {
      findOne: jest.fn(() => user),
      save: jest.fn(),
    } as any;

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
    todosService = module.get<TodosService>(TodosService);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });

  it('should create todo', async () => {
    const createFn = jest.spyOn(todosService, 'create')

    await controller.create('userid1', {
      title: createTodoDto.title,
      date: createTodoDto.date
    })

    expect(createFn).toBeCalledWith(createTodoDto);
  });

  it('should not create todos greater than in setting', async () => {
    jest.spyOn(trxMock, 'findAndCount').mockImplementation(() => [, 1]);

    await expect(async () => {
      await controller.create('userid1', {
        title: createTodoDto.title,
        date: createTodoDto.date
      })
    }).rejects.toThrowError("Exceed number of todos per day");

  });
});
