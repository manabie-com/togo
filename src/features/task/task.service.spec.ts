import { Test, TestingModule } from '@nestjs/testing';
import { UserEntity } from '../user/entities/user.entity';
import { UserRepository } from '../user/user.repository';
import { UserService } from '../user/user.service';
import { TaskEntity } from './entities/task.entity';
import { TaskRepository } from './task.repository';
import { TaskService } from './task.service';

describe('TaskService', () => {
  let service: TaskService;
  let spyUserService: UserService;
  let spyTaskRepository: TaskRepository;

  beforeEach(async () => {
    const mockUserService = {
      findOneById: jest.fn()
    };

    const mockUserRepository = {
      save: jest.fn()
    };

    const mockTaskRepository = {
      getNumberOfTaskByUserId: jest.fn(),
      save: jest.fn(x => x)
    };

    const module: TestingModule = await Test.createTestingModule({
      providers: [
        TaskService,
        {
          provide: UserService,
          useValue: mockUserService
        },
        {
          provide: UserRepository,
          useValue: mockUserRepository
        },
        {
          provide: TaskRepository,
          useValue: mockTaskRepository
        },
      ],
    }).compile();

    service = module.get<TaskService>(TaskService);
    spyUserService = module.get<UserService>(UserService);
    spyTaskRepository = module.get<TaskRepository>(TaskRepository)
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });

  describe('create function', () => {
    it('should call findOneById function of UserService', async () => {
      const payload = {
        content: 'Task One',
        userId: 1
      };

      jest.spyOn(spyUserService, 'findOneById').mockResolvedValue({
        id: 1,
        email: 'steven@gmail.com',
        maxTasks: 5
      } as UserEntity);

      await service.create(payload);
      expect(spyUserService.findOneById).toHaveBeenCalled();
      expect(spyUserService.findOneById).toHaveBeenCalledWith(payload.userId);
    });

    it('should throw not found user exception', async () => {
      const payload = {
        content: 'Task One',
        userId: 1
      };

      let errorMessage = '';

      jest.spyOn(spyUserService, 'findOneById').mockResolvedValue(null);

      try {
        await service.create(payload)
      } catch(err) {
        errorMessage = err.message;
      }
  
      expect(errorMessage).toBe(`User with id ${payload.userId} does not exist`);
    });

    it('should throw reach limit the number of tasks exception', async () => {
      const payload = {
        content: 'Task One',
        userId: 1
      };

      let errorMessage = '';

      jest.spyOn(spyUserService, 'findOneById').mockResolvedValue({
        id: 1,
        email: 'steven@gmail.com',
        maxTasks: 5
      } as UserEntity);

      jest.spyOn(spyTaskRepository, 'getNumberOfTaskByUserId').mockResolvedValue(5);

      try {
        await service.create(payload)
      } catch(err) {
        errorMessage = err.message;
      }
  
      expect(errorMessage).toBe(`The number of tasks of user has reached limit.`);
    });

    it('should save and return a new task record', async () => {
      const payload = {
        content: 'Task One',
        userId: 1
      };

      const expectedResult = {
        content: 'Task One',
        userId: 1
      };

      jest.spyOn(spyUserService, 'findOneById').mockResolvedValue({
        id: 1,
        email: 'steven@gmail.com',
        maxTasks: 5
      } as UserEntity);

      jest.spyOn(spyTaskRepository, 'getNumberOfTaskByUserId').mockResolvedValue(4);      
 
      expect(await service.create(payload)).toEqual(expectedResult);
    });
  });
});
