import { Test, TestingModule } from '@nestjs/testing';

import { TestUtils } from '@test/test.utils';
import { DatabaseModule } from '@test/database';
import { DatabaseService } from '@test/database/database.service';
import { TodoTaskController } from './todo-task.controller';
import { JWTPayload } from '@modules/auth/dto';
import { TodoTask } from './entities/todo-task.entity';
import {
  ReachedMaximumTaskTodayBadRequestException,
  TaskIsAssignedToYouBadRequestException,
  TaskIsNotFoundException,
} from '@modules/common/exceptions';
import * as users from '../../../src/test/fixtures/User';

describe('TodoTaskController', () => {
  let controller: TodoTaskController;
  let module: TestingModule;
  let testUtils: TestUtils;

  beforeAll(async () => {
    const testModule = await Test.createTestingModule({
      imports: [DatabaseModule],
      providers: [DatabaseService, TestUtils],
    }).compile();

    testUtils = testModule.get<TestUtils>(TestUtils);

    await testUtils.reloadFixtures();

    module = await Test.createTestingModule({
      controllers: [TodoTaskController],
      providers: [
        ...testUtils.getConnectionServiceGroup(),
        ...testUtils.getTodoTaskServiceGroup(),
        ...testUtils.getUserServiceGroup(),
      ],
    }).compile();

    controller = module.get<TodoTaskController>(TodoTaskController);
  });

  afterAll(async () => {
    await testUtils.closeDbConnection();
    module.close();
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });

  describe('findAll', () => {
    it('should return task list', async () => {
      const result = await controller.findAll();

      expect(result.length).toBeGreaterThan(0);
    });

    it('should return my task list', async () => {
      const jwtPayload: JWTPayload = {
        userId: users[0].id,
        username: users[0].username,
      };
      const result = await controller.findMyTasks(jwtPayload);

      expect(result.length).toBeGreaterThan(0);
    });
  });

  describe('pick', () => {
    it('should throw ReachedMaximumTaskTodayBadRequestException', async () => {
      const task = (await (await testUtils.databaseService.getRepository(TodoTask)).findOne({
        where: { assigneeId: users[1].id },
      })) as TodoTask;

      const jwtPayload: JWTPayload = {
        userId: users[1].id,
        username: users[1].username,
      };

      await expect(controller.pick({ taskId: task.id }, jwtPayload)).rejects.toThrowError(
        new ReachedMaximumTaskTodayBadRequestException(),
      );
    });

    it('should throw TaskIsAssignedToYouBadRequestException', async () => {
      const task = (await (await testUtils.databaseService.getRepository(TodoTask)).findOne({
        where: { assigneeId: users[2].id },
      })) as TodoTask;

      const jwtPayload: JWTPayload = {
        userId: users[2].id,
        username: users[2].username,
      };

      await expect(controller.pick({ taskId: task.id }, jwtPayload)).rejects.toThrowError(
        new TaskIsAssignedToYouBadRequestException(),
      );
    });

    it('should throw TaskIsNotFoundException', async () => {
      const jwtPayload: JWTPayload = {
        userId: users[2].id,
        username: users[2].username,
      };

      await expect(controller.pick({ taskId: 'fakeTaskId' }, jwtPayload)).rejects.toThrowError(
        new TaskIsNotFoundException(),
      );
    });

    it('should throw TaskIsNotFoundException', async () => {
      const task = new TodoTask({
        summary: 'TODO-X: Buy Iphone 13 Pro Max',
        description: 'Try to pass to have money to buy iphone 13 pro Max',
      });

      const createdTask = await (await testUtils.databaseService.getRepository(TodoTask)).save(task);

      const jwtPayload: JWTPayload = {
        userId: users[2].id,
        username: users[2].username,
      };

      const result = await controller.pick({ taskId: createdTask.id }, jwtPayload);

      expect({ message: 'Pick task successfully' }).toEqual(result);
    });
  });
});
