import { Test, TestingModule } from '@nestjs/testing';
import { Repository } from 'typeorm';
import { TodoEntity } from './entities/todo.entity';
import { TodosService } from './todos.service';

describe('TodosService', () => {
  let service: TodosService;
  let todosRepository: Repository<TodoEntity>;
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
    todosRepository = {
      create: jest.fn(() => creatingTodo),
      save: jest.fn(),
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
    expect(todosRepository.save).toBeCalledWith(creatingTodo);
  });

});
