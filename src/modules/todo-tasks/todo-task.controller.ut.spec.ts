import { Test, TestingModule } from '@nestjs/testing';
import { getConnectionToken, getRepositoryToken, TypeOrmModule } from '@nestjs/typeorm';
import { ConnectionOptions } from 'typeorm';

import { User } from '@modules/users/entities/user.entity';
import { UserService } from '@modules/users/user.service';
import { TodoTaskController } from './todo-task.controller';
import { TodoTaskService } from './todo-task.service';
import { TodoTask } from './entities/todo-task.entity';
import { faker } from '@faker-js/faker';
import { JWTPayload } from '@modules/auth/dto';
import { USER_ROLE_ID } from '@modules/common/constant';
import { AssignTaskDto } from './dto/assign-task.dto';
import { UserTaskConfig } from '@modules/users/entities/user-task-config.entity';
import { Role } from '@modules/roles/entities/role.entity';
import { Permission } from '@modules/permissions/permission.entity';
import {
  ReachedMaximumTaskTodayBadRequestException,
  TaskIsAssignedToYouBadRequestException,
  TaskIsNotFoundException,
} from '@modules/common/exceptions';

describe('TodoTaskController', () => {
  let module: TestingModule;
  let controller: TodoTaskController;

  const userId = faker.datatype.uuid();

  const userTaskConfigs = [new UserTaskConfig({ userId: userId, numberOfTaskPerDay: 5, date: new Date() })];

  const user = new User({
    id: userId,
    username: faker.datatype.uuid(),
    roleId: USER_ROLE_ID,
    userTaskConfigs: userTaskConfigs,
  });

  const myTasks = [
    new TodoTask({
      id: faker.datatype.uuid(),
      summary: faker.commerce.product(),
      description: faker.lorem.paragraph(),
      assigneeId: userId,
    }),
  ];

  const task = new TodoTask({
    id: faker.datatype.uuid(),
    summary: faker.commerce.product(),
    description: faker.lorem.paragraph(),
    assigneeId: null,
  });

  const mockTodoTaskService = {
    findAll: jest.fn().mockResolvedValue([task]),

    findById: jest.fn().mockResolvedValue(task),
    findOne: jest.fn().mockResolvedValue(myTasks[0]),
  };

  const mockUserService = {
    findById: jest.fn().mockResolvedValue(user),
    userTaskConfigs: jest.fn().mockResolvedValue(userTaskConfigs),
  };

  const mockGetConnectionToken = {
    mockGetConnectionToken: jest.fn().mockImplementation(() => mockGetConnectionToken),
    select: jest.fn().mockImplementation(() => mockGetConnectionToken),
    andWhere: jest.fn().mockImplementation(() => mockGetConnectionToken),
    addSelect: jest.fn().mockImplementation(() => mockGetConnectionToken),
    leftJoinAndSelect: jest.fn().mockImplementation(() => mockGetConnectionToken),
    innerJoin: jest.fn().mockImplementation(() => mockGetConnectionToken),
    leftJoin: jest.fn().mockImplementation(() => mockGetConnectionToken),
    groupBy: jest.fn().mockImplementation(() => mockGetConnectionToken),
    where: jest.fn().mockImplementation(() => mockGetConnectionToken),
    findOne: jest.fn().mockImplementation(() => null),
    save: jest.fn().mockImplementation(() => null),
    getOne: jest.fn().mockImplementation(() => null),
    getMany: jest.fn().mockImplementation(() => null),
    getRawMany: jest.fn().mockImplementation(() => null),
  };

  const mockUserRepository = {};

  beforeEach(async () => {
    module = await Test.createTestingModule({
      imports: [
        TypeOrmModule.forRoot({
          type: 'sqlite',
          database: ':memory:',
          dropSchema: true,
          synchronize: true,
          logging: false,
          entities: [TodoTask, User, UserTaskConfig, Role, Permission],
          name: 'default',
          keepConnectionAlive: true,
        } as ConnectionOptions),
      ],
      controllers: [TodoTaskController],
      providers: [
        TodoTaskService,
        UserService,
        {
          provide: getConnectionToken(),
          useValue: mockGetConnectionToken,
        },
        {
          provide: getRepositoryToken(User),
          useValue: mockUserRepository,
        },
      ],
    })
      .overrideProvider(TodoTaskService)
      .useValue(mockTodoTaskService)
      .overrideProvider(UserService)
      .useValue(mockUserService)
      .compile();

    controller = module.get<TodoTaskController>(TodoTaskController);
  });

  afterEach(async () => {
    await module.close();
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });

  describe('findAll', () => {
    it('should be return a list task', async () => {
      const result = await controller.findAll();

      expect(result).toMatchObject([task]);
    });

    it('should return my tasks list', async () => {
      const jwtPayload: JWTPayload = {
        userId: user.id,
        username: user.username,
      };

      jest.spyOn(mockTodoTaskService, 'findAll').mockResolvedValue(myTasks);

      const result = await controller.findMyTasks(jwtPayload);

      expect(result).toMatchObject(myTasks);
    });
  });

  describe('pick', () => {
    it('should throw TaskIsAssignedToYouBadRequestException', async () => {
      const jwtPayload: JWTPayload = {
        userId: user.id,
        username: user.username,
      };

      const dto: AssignTaskDto = {
        taskId: myTasks[0].id,
      };

      await expect(controller.pick(dto, jwtPayload)).rejects.toThrowError(new TaskIsAssignedToYouBadRequestException());
    });

    it('should throw TaskIsNotFoundException', async () => {
      const jwtPayload: JWTPayload = {
        userId: user.id,
        username: user.username,
      };

      const dto: AssignTaskDto = {
        taskId: 'fakeTaskId',
      };

      jest.spyOn(mockTodoTaskService, 'findOne').mockResolvedValue(undefined);

      await expect(controller.pick(dto, jwtPayload)).rejects.toThrowError(new TaskIsNotFoundException());
    });

    it('should return Pick task successfully message', async () => {
      const jwtPayload: JWTPayload = {
        userId: userId,
        username: user.username,
      };

      const dto: AssignTaskDto = {
        taskId: task.id,
      };

      const expectedResult: TodoTask = new TodoTask({ ...task, assigneeId: userId });

      jest.spyOn(mockTodoTaskService, 'findOne').mockResolvedValue(task);
      jest.spyOn(TodoTask.prototype, 'save').mockImplementation(async () => Promise.resolve(expectedResult));

      const result = await controller.pick(dto, jwtPayload);

      expect(result).toEqual({ message: 'Pick task successfully' });
    });

    it('should throw ReachedMaximumTaskTodayBadRequestException', async () => {
      const jwtPayload: JWTPayload = {
        userId: userId,
        username: user.username,
      };

      const dto: AssignTaskDto = {
        taskId: task.id,
      };

      jest.spyOn(User.prototype, 'myTotalTask').mockResolvedValue(6);

      await expect(controller.pick(dto, jwtPayload)).rejects.toThrowError(
        new ReachedMaximumTaskTodayBadRequestException(),
      );
    });
  });
});
