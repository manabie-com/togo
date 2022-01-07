import { Test, TestingModule } from '@nestjs/testing';
import { TaskRepository } from '../../src/task/task.repository';
import { TaskService } from '../../src/task/task.service';
import { mockCreateTaskDto, mockVideo } from './task';

describe('TaskService', () => {
  let service: TaskService;
  let repository: TaskRepository;

  let mockRepository = {
    create: jest.fn()
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
    it('should call VideosRepository create with correct values', async () => {
      const createSpy = jest.spyOn(repository, 'create');
      const mockParam = mockCreateTaskDto();
      await service.create(mockParam);
      expect(createSpy).toHaveBeenCalledWith(mockParam);
    })

    it('should throw if VideosRepository create throws', async () => {
      jest.spyOn(repository, 'create').mockRejectedValueOnce(new Error());
      await expect(service.create(mockCreateTaskDto())).rejects.toThrow(new Error());
    })

    it('should return a video on success', async () => {
      const mockReturn = mockVideo();
      jest.spyOn(repository, 'create').mockResolvedValueOnce(mockReturn);
      const response = await service.create(mockCreateTaskDto());
      expect(response).toEqual(mockReturn);
    })
  })
});
