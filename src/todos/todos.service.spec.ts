import { Test, TestingModule } from '@nestjs/testing';
import { Repository } from 'typeorm';
import { TodoEntity } from './entities/todo.entity';
import { TodosService } from './todos.service';

describe('TodosService', () => {
  let service: TodosService;
  let todosRepository: Repository<TodoEntity> = {} as any;
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

    const module: TestingModule = await Test.createTestingModule({
      providers: [
        {
          provide: 'TodoEntityRepository',
          useValue: todosRepository
        },
        TodosService],
    }).compile();

    service = module.get<TodosService>(TodosService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });

  it('should create todo', async () => {
    await service.create(createTodoDto);

    const saveFn = jest.spyOn(trxMock, 'save')
    expect(saveFn).toBeCalledWith(creatingTodo);
  });

  it('should not create todos greater than in setting', async () => {
    jest.spyOn(trxMock, 'findAndCount').mockImplementation(() => [, 1]);

    await expect(async () => {
      await service.create(createTodoDto);
    }).rejects.toThrowError("Exceed number of todos per day");

  });

});
