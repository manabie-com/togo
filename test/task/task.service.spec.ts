import { Test, TestingModule } from '@nestjs/testing';
import { UserService } from '../../src/user/user.service';
import { TaskRepository } from '../../src/task/task.repository';
import { TaskService } from '../../src/task/task.service';
import { mockCreateTaskDto, mockTask } from './task';

describe('TaskService', () => {
  let service: TaskService;
  let userService: UserService;
  let repository: TaskRepository;

  let mockRepository = {
    create: jest.fn(),
    findAll: jest.fn(),
  };

  let mockUserService = {
    incrementTask: jest.fn()
  };

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        TaskService,
        {
          provide: TaskRepository,
          useValue: mockRepository
        },
        {
          provide: UserService,
          useValue: mockUserService
        }
      ],
    }).compile();

    service = module.get<TaskService>(TaskService);
    repository = module.get<TaskRepository>(TaskRepository);
    userService = module.get<UserService>(UserService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });

  describe('create()', () => {
    it('should call TaskRepository create with correct values', async () => {
      const createSpy = jest.spyOn(repository, 'create');
      const mockParam = mockCreateTaskDto();
      await service.create(mockParam);
      expect(createSpy).toHaveBeenCalledWith(mockParam);
    })

    it('should throw if TaskRepository create throws', async () => {
      jest.spyOn(repository, 'create').mockRejectedValueOnce(new Error());
      await expect(service.create(mockCreateTaskDto())).rejects.toThrow(new Error());
    })

    it('should return a task on success', async () => {
      const mockReturn = mockTask();
      jest.spyOn(repository, 'create').mockResolvedValueOnce(mockReturn);
      const response = await service.create(mockCreateTaskDto());
      expect(response).toEqual(mockReturn);
    })

    it('should call UserService incrementTask with correct values', async () => {
      const createSpy = jest.spyOn(userService, 'incrementTask');
      const mockParam = mockCreateTaskDto();
      await service.create(mockParam);
      expect(createSpy).toHaveBeenCalled();
    })

    it('should throw if UserService incrementTask throws UserNotFound', async () => {
      jest.spyOn(userService, 'incrementTask').mockRejectedValueOnce(new Error("User not found"));
      await expect(service.create(mockCreateTaskDto())).rejects.toThrow(new Error("User not found"));
    })

    it('should throw if UserService incrementTask throws DailyTaskLimitExceeded', async () => {
      jest.spyOn(userService, 'incrementTask').mockRejectedValueOnce(new Error("Daily task limit exceeded"));
      await expect(service.create(mockCreateTaskDto())).rejects.toThrow(new Error("Daily task limit exceeded"));
    })

    it('should return a task on success', async () => {
      const mockReturn = mockTask();
      jest.spyOn(repository, 'create').mockResolvedValueOnce(mockReturn);
      jest.spyOn(userService, 'incrementTask').mockResolvedValueOnce(true);
      const response = await service.create(mockCreateTaskDto());
      expect(response).toEqual(mockReturn);
    })
  })

  describe('findAll()', () => {
    it('should call TaskRepository find all', async () => {
      const findSpy = jest.spyOn(repository, 'findAll');
      await service.findAll();
      expect(findSpy).toHaveBeenCalled();
    })

    it('should throw if TaskRepository find all throws', async () => {
      jest.spyOn(repository, 'findAll').mockRejectedValueOnce(new Error());
      await expect(service.findAll()).rejects.toThrow(new Error())
    })

    it('should return tasks on success', async () => {
      const mockReturn = [mockTask()]
      jest.spyOn(repository, 'findAll').mockResolvedValueOnce(mockReturn);
      const response = await service.findAll();
      expect(response).toEqual(mockReturn);
    })
  })
});
