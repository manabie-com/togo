import { Test, TestingModule } from '@nestjs/testing';
import { getRepositoryToken } from '@nestjs/typeorm';
import { TodoTask } from './entities/todo-task.entity';

import { TodoTaskService } from './todo-task.service';

describe('TodoTaskService', () => {
  let service: TodoTaskService;

  const mockTodoTaskRepository = {};

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        TodoTaskService,
        {
          provide: getRepositoryToken(TodoTask),
          useValue: mockTodoTaskRepository,
        },
      ],
    }).compile();

    service = module.get<TodoTaskService>(TodoTaskService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });

  // TODO: Add more cases to cover more cases
});
