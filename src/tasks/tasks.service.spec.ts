import { MongooseModule } from '@nestjs/mongoose';
import { Test, TestingModule } from '@nestjs/testing';
import { UsersService } from '../users/users.service';
import { constants } from '../constants';
import { Task, TaskSchema } from './schemas/task.schema';
import { TasksService } from './tasks.service';
import { User, UserSchema } from '../users/schemas/user.schema';
import { TaskInterface } from './interfaces/task.interface';

describe('TasksService', () => {
  let service: TasksService;
  const today = new Date();
  const mockTask: TaskInterface = {
    owner: '620b0d967c5565386dab8d8d',
    description: 'solve challenge',
  };

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      imports: [
        MongooseModule.forRoot(constants.sandbox),
        MongooseModule.forFeature([{ name: Task.name, schema: TaskSchema }]),
        MongooseModule.forFeature([{ name: User.name, schema: UserSchema }]),
      ],
      providers: [TasksService, UsersService],
    }).compile();

    service = module.get<TasksService>(TasksService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });

  it('should create a task', async () => {
    const newTask = await service.create(mockTask);
    expect(newTask).toEqual(
      expect.objectContaining({
        ...mockTask,
      }),
    );
  });

  it('should return a number of task which created today', async () => {
    const count = await service.countTasks('11111', today, today);
    expect(count).toEqual(expect.any(Number));
  });
});
