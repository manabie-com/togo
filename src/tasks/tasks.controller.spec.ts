import { MongooseModule } from '@nestjs/mongoose';
import { Test, TestingModule } from '@nestjs/testing';
import { constants } from '../constants';
import { TasksController } from './tasks.controller';
import { TasksService } from './tasks.service';
import { Task, TaskSchema } from './schemas/task.schema';
import { UsersService } from '../users/users.service';
import { User, UserSchema } from '../users/schemas/user.schema';
import { CreateTaskDto } from './dto/create-task.dto';

describe('TasksController', () => {
  let controller: TasksController;

  const mockTaskDto: CreateTaskDto = {
    description: 'solve challenge',
  };
  const mockJwtGuard = {
    user: {
      id: '620b0d967c5565386dab8dff',
    },
  };
  const mockParams = {
    page: 1,
    size: 10,
  };

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      imports: [
        MongooseModule.forRoot(constants.sandbox),
        MongooseModule.forFeature([{ name: Task.name, schema: TaskSchema }]),
        MongooseModule.forFeature([{ name: User.name, schema: UserSchema }]),
      ],
      controllers: [TasksController],
      providers: [TasksService, UsersService],
    }).compile();

    controller = module.get<TasksController>(TasksController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });

  it('should return empty list', async () => {
    const newTask = await controller.findAll(mockJwtGuard, mockParams);
    expect(newTask).toEqual({
      data: [],
      total: 0,
      ...mockParams,
    });
  });
});
