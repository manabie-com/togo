import { Test, TestingModule } from '@nestjs/testing';
import { TodoService } from './todo.service';
import { getRepositoryToken } from '@nestjs/typeorm';
import { In, Between, Repository } from 'typeorm';
import { startOfToday, endOfToday } from 'date-fns';
import { Todo, TodoStatus } from './todo.entity';

describe('TodoService', () => {
  let service: TodoService;
  let todosRepository: Repository<Todo>;

  const mockUserDto = {
    username: 'john',
    password: 'john',
    limitPerDay: 5,
  };
  const mockTodoDto = {
    id: 'todo',
    status: TodoStatus.ACTIVE,
  };

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        TodoService,
        {
          provide: getRepositoryToken(Todo),
          useValue: {
            find: jest.fn(),
            count: jest.fn(),
            update: jest.fn(),
            delete: jest.fn(),
          },
        },
      ],
    }).compile();

    service = module.get<TodoService>(TodoService);
    todosRepository = module.get<Repository<Todo>>(getRepositoryToken(Todo));
  });

  it('service should be defined', () => {
    expect(service).toBeDefined();
  });

  it('userRepository should be defined', () => {
    expect(todosRepository).toBeDefined();
  });

  describe('getTodos', () => {
    it('should get all todos', async () => {
      await service.getTodos(mockUserDto.username);

      expect(todosRepository.find).toBeCalledWith({
        where: {
          user: {
            username: mockUserDto.username,
          },
        },
        order: {
          createAt: 'DESC',
        },
      });
    });

    it('should get todos with status', async () => {
      await service.getTodos(mockUserDto.username, mockTodoDto.status);

      expect(todosRepository.find).toBeCalledWith({
        where: {
          user: {
            username: mockUserDto.username,
          },
          status: mockTodoDto.status,
        },
        order: {
          createAt: 'DESC',
        },
      });
    });
  });

  describe('countTodoDaily', () => {
    it('should get todos daily count', async () => {
      await service.countTodoDaily(mockUserDto.username);

      expect(todosRepository.count).toBeCalledWith({
        where: {
          user: {
            username: mockUserDto.username,
          },
          createAt: Between(startOfToday(), endOfToday()),
        },
      });
    });
  });

  describe('setTodoStatus', () => {
    it('should update todo status', async () => {
      await service.setTodoStatus(mockTodoDto.id, mockTodoDto.status);

      expect(todosRepository.update).toBeCalledWith(
        { id: mockTodoDto.id },
        { status: mockTodoDto.status },
      );
    });
  });

  describe('setManyTodoStatus', () => {
    it('should update many todos status', async () => {
      await service.setManyTodoStatus([mockTodoDto.id], mockTodoDto.status);

      expect(todosRepository.update).toBeCalledWith(
        { id: In([mockTodoDto.id]) },
        { status: mockTodoDto.status },
      );
    });
  });

  describe('deleteTodoById', () => {
    it('should delete todo', async () => {
      await service.deleteTodoById(mockTodoDto.id);

      expect(todosRepository.delete).toBeCalledWith(mockTodoDto.id);
    });
  });

  describe('deleteAllTodos', () => {
    it('should delete todo', async () => {
      await service.deleteAllTodos(mockUserDto.username);

      expect(todosRepository.delete).toBeCalledWith({
        user: {
          username: mockUserDto.username,
        },
      });
    });
  });
});
