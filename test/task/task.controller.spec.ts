import { Test, TestingModule } from '@nestjs/testing';
import { TaskController } from '../../src/task/task.controller';
import { TaskService } from '../../src/task/task.service';
import { mockCreateTaskDto, mockVideo } from './task';

describe('TaskController', () => {
  let controller: TaskController;
  let service: TaskService;

  const mockService = {
    create: jest.fn()
  }

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [TaskController],
      providers: [
        {
          provide: TaskService,
          useValue: mockService
        }
      ]
    }).compile();

    controller = module.get<TaskController>(TaskController);
    service = module.get<TaskService>(TaskService);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });

  describe('create()', () => {

    it('should have create function', () => {
      expect(controller.create).toBeDefined();
    });

    it('should call TaskService create with correct values', async () => {
      const createSpy = jest.spyOn(service, 'create');
      const mockParam = mockCreateTaskDto();
      await controller.create(mockParam);
      expect(createSpy).toHaveBeenCalledWith(mockParam);
    })

    it('should throw if VideosService create throws', async () => {
      jest.spyOn(service, 'create').mockRejectedValueOnce(new Error());
      await expect(controller.create(mockCreateTaskDto())).rejects.toThrow(new Error());
    })

    it('should return a video on success', async () => {
      const mockReturn = mockVideo();
      jest.spyOn(service, 'create').mockResolvedValueOnce(mockReturn);
      const response = await controller.create(mockCreateTaskDto())
      expect(response).toEqual(mockReturn);
    })
  })
  
});
