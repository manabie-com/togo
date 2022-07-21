import { Test, TestingModule } from '@nestjs/testing';
import { Repository } from 'typeorm';
import { TodoEntity } from './entities/todo.entity';
import { TodosService } from './todos.service';

describe('TodosService', () => {
  let service: TodosService;
  let todosRepository: Repository<TodoEntity>;

  beforeEach(async () => {
    todosRepository = jest.mock as any;

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
});
