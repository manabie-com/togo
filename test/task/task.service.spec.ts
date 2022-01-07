import { Test, TestingModule } from '@nestjs/testing';
import { TaskRepository } from '../../src/task/task.repository';
import { TaskService } from '../../src/task/task.service';
import { mockCreateTaskDto, mockTask } from './task';

describe('TaskService', () => {
  let service: TaskService;
  let repository: TaskRepository;

  let mockRepository = {
    create: jest.fn(),
    findAll: jest.fn(),
  };

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        TaskService,
        {
          provide: TaskRepository,
          useValue: mockRepository
        }
      ],
    }).compile();

    service = module.get<TaskService>(TaskService);
    repository = module.get<TaskRepository>(TaskRepository);
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
