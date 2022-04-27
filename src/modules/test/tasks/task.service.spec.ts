import { getModelToken } from '@nestjs/sequelize';
import { Test, TestingModule } from '@nestjs/testing';
import { Task, UserSettingTask } from '../../common/entities';
import { TaskService } from '../../features/tasks/task.service';

const mockTaskModel = () => ({
  findOne: jest.fn(),
  count: jest.fn(),
  create: jest.fn(),
});

const mockUserSettingTaskModel = () => ({
  findOne: jest.fn(),
});

const dataTaskCreate = {
  tasks: [
    {
      description: 'test',
      excutors: ['1', '2'],
      note: 'test',
      title: '12',
      watchers: ['2'],
    },
    {
      description: 'test',
      excutors: ['1', '2'],
      note: 'test',
      title: '12',
      watchers: ['2'],
    },
  ],
};

describe('TaskService', () => {
  let taskService: TaskService;
  let userSettingModel;
  let taskModel;
  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        TaskService,
        {
          provide: getModelToken(Task),
          useFactory: mockTaskModel,
        },
        {
          provide: getModelToken(UserSettingTask),
          useFactory: mockUserSettingTaskModel,
        },
      ],
    }).compile();
    taskService = module.get<TaskService>(TaskService);
    userSettingModel = module.get(getModelToken(UserSettingTask));
    taskModel = module.get(getModelToken(Task));
  });

  describe('createTask', () => {
    it('should be error when not found user setting', async () => {
      jest.spyOn(userSettingModel, 'findOne').mockResolvedValue(null);
      try {
        await taskService.createTask('userId', dataTaskCreate);
      } catch (err) {
        expect(err.message).toBe('Not found setting of user');
      }
    });

    it('should be error when maximum create task in a day', async () => {
      jest.spyOn(userSettingModel, 'findOne').mockResolvedValue({
        get: () => ({
          maximum: 9,
        }),
      });

      jest.spyOn(taskModel, 'count').mockResolvedValue(8);
      try {
        await taskService.createTask('userId', dataTaskCreate);
      } catch (err) {
        expect(err.message).toBe('Exceed the number of tasks created per day.');
      }
    });

    it('should be save when data valid', async () => {
      jest.spyOn(userSettingModel, 'findOne').mockResolvedValue({
        get: () => ({
          maximum: 10,
        }),
      });

      jest.spyOn(taskModel, 'count').mockResolvedValue(8);
      await taskService.createTask('userId', dataTaskCreate);

      expect(taskModel.create).toBeCalledTimes(dataTaskCreate.tasks.length);
    });
  });
});
