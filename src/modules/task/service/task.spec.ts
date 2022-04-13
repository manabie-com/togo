import { Test, TestingModule } from '@nestjs/testing';
import { TaskService } from './task.service';
import { UserService } from '../../user/service/user.service';
import { Repository } from 'typeorm';
import { TaskEntity } from '../entity/task.entity';
import { getRepositoryToken } from '@nestjs/typeorm';
import { UserEntity } from 'src/modules/user/entity/user.entity';
import { HttpException } from '@nestjs/common';

describe('TaskService', () => {
  let taskService: TaskService;
  let userService: UserService;
  let repositoryMock: MockType<Repository<TaskEntity>>;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        TaskService,
        UserService,
        {
          provide: getRepositoryToken(TaskEntity),
          useFactory: repositoryMockFactory,
        },
        {
          provide: getRepositoryToken(UserEntity),
          useFactory: repositoryMockFactory,
        },
      ],
    }).compile();
    taskService = module.get<TaskService>(TaskService);
    userService = module.get<UserService>(UserService);
    repositoryMock = module.get(getRepositoryToken(TaskEntity));
  });

  it('TaskService - should be defined', () => {
    expect(taskService).toBeDefined();
  });

  it('#Create task should be failed by reach max user task per day', async () => {
    const taskMockValues = {
      title: 'title',
      content: 'content',
      startDate: new Date(),
    };

    jest.spyOn(userService, 'getUserById').mockImplementation(() =>
      Promise.resolve({
        id: '61f320f3-6b7a-432f-bb1c-80a48ec6ffa7',
        username: 'togo',
        password: '123456',
        maxTask: 5,
      }),
    );
    jest
      .spyOn(taskService, 'countTaskOnDate')
      .mockImplementation(() => Promise.resolve(5));
    repositoryMock.save.mockReturnValue(taskMockValues);

    try {
      await taskService.create(
        taskMockValues,
        '61f320f3-6b7a-432f-bb1c-80a48ec6ffa7',
      );
    } catch (error) {
      expect(error).toBeInstanceOf(HttpException);
    }
  });

  it('#Create task should work', async () => {
    const taskMockValues = {
      title: 'title',
      content: 'content',
      startDate: new Date(),
    };

    jest.spyOn(userService, 'getUserById').mockImplementation(() =>
      Promise.resolve({
        id: '61f320f3-6b7a-432f-bb1c-80a48ec6ffa7',
        username: 'togo',
        password: '123456',
        maxTask: 5,
      }),
    );
    jest
      .spyOn(taskService, 'countTaskOnDate')
      .mockImplementation(() => Promise.resolve(1));
    repositoryMock.save.mockReturnValue(taskMockValues);

    const task = await taskService.create(
      taskMockValues,
      '61f320f3-6b7a-432f-bb1c-80a48ec6ffa7',
    );

    expect(task.title).toEqual(taskMockValues.title);
    expect(repositoryMock.save).toHaveBeenCalled();
    expect(taskService.countTaskOnDate).toHaveBeenCalled();
    expect(userService.getUserById).toHaveBeenCalled();
  });
});

export type MockType<T> = {
  // eslint-disable-next-line @typescript-eslint/ban-types
  [P in keyof T]?: jest.Mock<{}>;
};

export const repositoryMockFactory: () => MockType<Repository<any>> = jest.fn(
  () => ({
    findOne: jest.fn((entity) => entity),
    save: jest.fn((entity) => entity),
  }),
);
