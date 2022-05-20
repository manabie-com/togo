import { Test, TestingModule } from '@nestjs/testing';
import { CoreModule } from '../../globalModule/core.module';
import { TaskController } from './task.controller';
import { TaskService } from './task.service';

describe('TaskController', () => {
  let controller: TaskController;
  let spyTaskService: TaskService;

  beforeEach(async () => {
    const mockTaskService = {
      create: jest.fn()
    };

    const module: TestingModule = await Test.createTestingModule({
      imports: [CoreModule],
      controllers: [TaskController],
      providers: [
        TaskService,
        {
          provide: TaskService,
          useValue: mockTaskService
        }
      ],
    }).compile();

    controller = module.get<TaskController>(TaskController);
    spyTaskService = module.get<TaskService>(TaskService);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });

  describe('create function', () => {
    it('should call create function of TaskService', async () => {
      const payload = {
        content: 'Task One',
        userId: 1
      };

      await controller.create(payload);
      expect(spyTaskService.create).toHaveBeenCalled();
      expect(spyTaskService.create).toHaveBeenCalledWith(payload);
    });
  });
});
