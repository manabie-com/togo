import { Test, TestingModule } from '@nestjs/testing';
import { InputType, Field, registerEnumType } from '@nestjs/graphql';
import { PrismaService } from './prisma.service';
import { UserResolver } from './resolvers.user';
import { TaskResolver } from './resolvers.task';
import { LoginUserInput } from './dtos/login-user.input';
import { CreateTaskInput } from './dtos/create-task.input';

@InputType()
class TaskOrderByUpdatedAtInput {
  @Field((type) => SortOrder)
  updatedAt: SortOrder;
}

enum SortOrder {
  asc = 'asc',
  desc = 'desc',
}

registerEnumType(SortOrder, {
  name: 'SortOrder',
});

const contextMock = (data) => {
  const ctx = {
    req: {
      get: jest.fn().mockReturnValue(data),
    },
  };

  return ctx;
};
let taskId = 0;
const searchString = 'task';
const skip = 0;
const take = 3;
const orderBy: TaskOrderByUpdatedAtInput = { updatedAt: SortOrder.asc };

describe('TaskResolver', () => {
  let resolver: TaskResolver;
  let userResolver: UserResolver;
  const loginUserInput: LoginUserInput = {
    email: 'john.doe@email.com',
    password: '123456',
  };
  const createTaskInput: CreateTaskInput = {
    title: 'task1',
    content: 'this is content task1',
  };

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [TaskResolver, UserResolver, PrismaService],
    }).compile();

    resolver = module.get<TaskResolver>(TaskResolver);
    userResolver = module.get<UserResolver>(UserResolver);
  });

  it('should be defined', () => {
    expect(resolver).toBeDefined();
  });

  describe('createTask', () => {
    it('should createTask and return data', async () => {
      const user = await userResolver.login(loginUserInput);
      const ctx = contextMock(user.token);
      const task = await resolver.createTask(createTaskInput, ctx);
      taskId = task.id;

      expect(taskId).toEqual(task.id);
    });
  });

  describe('taskById', () => {
    it('should get task by id', async () => {
      const task = await resolver.taskById(taskId);
      expect(task.id).toEqual(taskId);
    });
  });

  describe('tasks', () => {
    it('should get multies tasks', async () => {
      const ctx = contextMock(null);
      const tasks = await resolver.tasks(
        searchString,
        skip,
        take,
        orderBy,
        ctx,
      );
      expect(tasks.length).toBeGreaterThan(0);
    });
  });
});
