import { Test, TestingModule } from '@nestjs/testing';
import { LimitTaskModule } from '../../../model/limit-task/limitTask.module';
import { TaskModule } from '../../../model/task/task.module';
import { TaskController } from './task.controller';

describe('TaskController', () => {
  let controller: TaskController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [TaskController],
      imports: [TaskModule, LimitTaskModule],
    }).compile();

    controller = module.get<TaskController>(TaskController);
  });

  it('create task', async () => {
    expect(controller).toBeDefined();
    expect(
      await controller.createTask({
        userId: '1',
        title: 'lam bai tap',
        desc: 'lam rat rat nhieu bai tap',
      }),
    ).toBeTruthy();
  });
});
