import { Test } from '@nestjs/testing';
import { getRepositoryToken } from '@nestjs/typeorm';
import { CreateTaskDto, GetTasksDto, UpdateTaskDto } from 'src/dto';
import { ETaskStatus, Task } from 'src/entities/task.entity';
import { ToDoList } from 'src/entities/toDoList.entity';
import { User } from 'src/entities/user.entity';
import { ToDoService } from 'src/services/todo.service';
import { UserService } from 'src/services/user.service';
import { DeleteResult } from 'typeorm';
import { TaskService } from '../task.service';

describe('The TaskService', () => {
  let taskService: TaskService;
  let userService: UserService;
  let todoService: ToDoService;
  let findOne: jest.Mock;
  let findAndCount: jest.Mock;
  let save: jest.Mock;
  let create: jest.Mock;
  let count: jest.Mock;
  let findUser: jest.Mock;
  let deleteTask: jest.Mock;
  beforeEach(async () => {
    findOne = jest.fn();
    findAndCount = jest.fn();
    save = jest.fn();
    create = jest.fn();
    count = jest.fn();
    findUser = jest.fn();
    deleteTask = jest.fn();
    const module = await Test.createTestingModule({
      providers: [
        UserService,
        ToDoService,
        TaskService,
        {
          provide: getRepositoryToken(User),
          useValue: {
            findOne: findUser,
          },
        },
        {
          provide: getRepositoryToken(ToDoList),
          useValue: {
            update: jest.fn(),
          },
        },
        {
          provide: getRepositoryToken(Task),
          useValue: {
            findOne,
            findAndCount,
            save,
            create,
            count,
            delete: deleteTask,
          },
        },
      ],
    }).compile();
    taskService = await module.get(TaskService);
    userService = await module.get(UserService);
    todoService = await module.get(ToDoService);
  });
  describe("count current user's pending tasks", () => {
    const result = 10;
    beforeEach(() => {
      count.mockReturnValue(Promise.resolve(10));
    });
    it('should return number', async () => {
      const userId = 1;
      const fetchedUser = await taskService.countTasks(userId);
      expect(fetchedUser).toEqual(result);
    });
  });
  describe('find tasks', () => {
    const task = new Task({});
    const result = {
      items: [task],
      total: 1,
    };
    beforeEach(() => {
      findAndCount.mockReturnValue(Promise.resolve([[task], 1]));
    });
    it('should return the task pagination', async () => {
      const queryParameters = new GetTasksDto();
      const tasksFetched = await taskService.find(queryParameters);
      expect(tasksFetched).toEqual(result);
    });
  });
  describe('when getting a task by id', () => {
    const result = new Task({ id: 1 });
    beforeEach(() => {
      findOne.mockReturnValue(Promise.resolve(result));
    });
    it('should return the task', async () => {
      const taskId = 1;
      const fetchedTask = await taskService.findOne(taskId, []);
      expect(fetchedTask).toEqual(result);
    });
  });
  describe('when updating a task', () => {
    const task = new Task({ id: 1 });
    const user = new User({ id: 1, name: 'NVA', dailyMaxTasks: 2 });
    const updateTaskDto = new UpdateTaskDto();
    updateTaskDto.deadlineAt = null;
    describe('task not found', () => {
      beforeEach(() => {
        findOne.mockReturnValue(Promise.resolve(null));
      });
      it('throw error', async () => {
        const taskId = 1;
        await expect(
          taskService.update(taskId, updateTaskDto),
        ).rejects.toThrow();
      });
    });
    describe('user not found', () => {
      beforeEach(() => {
        findOne.mockReturnValue(Promise.resolve(task));
        findUser.mockReturnValue(Promise.resolve(null));
      });
      it('throw error', async () => {
        const taskId = 1;
        updateTaskDto.userId = 5;
        await expect(
          taskService.update(taskId, updateTaskDto),
        ).rejects.toThrow();
      });
    });
    describe('Daily task count is reached', () => {
      beforeEach(() => {
        findOne.mockReturnValue(Promise.resolve(task));
        findUser.mockReturnValue(Promise.resolve(user));
        count.mockReturnValue(Promise.resolve(2));
      });
      it('throw error', async () => {
        const taskId = 1;
        updateTaskDto.userId = user.id;
        await expect(
          taskService.update(taskId, updateTaskDto),
        ).rejects.toThrow();
      });
    });
    describe('update task with status is nil', () => {
      beforeEach(() => {
        findOne.mockReturnValue(Promise.resolve(task));
        findUser.mockReturnValue(Promise.resolve(user));
        count.mockReturnValue(Promise.resolve(1));
        save.mockReturnValue(Promise.resolve(task));
      });
      it('should return a task', async () => {
        const result = task;
        result.user = user;
        updateTaskDto.userId = user.id;
        const taskId = 1;
        const updatedTask = await taskService.update(taskId, updateTaskDto);
        expect(updatedTask).toEqual(result);
      });
    });
    describe('update task with status is COMPLETE', () => {
      beforeEach(() => {
        findOne.mockReturnValue(Promise.resolve(task));
        findUser.mockReturnValue(Promise.resolve(user));
        count.mockReturnValue(Promise.resolve(1));
        save.mockReturnValue(Promise.resolve(task));
      });
      it('should return a task with no deadline', async () => {
        const result = task;
        result.user = user;
        updateTaskDto.userId = user.id;
        updateTaskDto.status = ETaskStatus.COMPLETE;
        updateTaskDto.deadlineAt = null;
        const taskId = 1;
        const updatedTask = await taskService.update(taskId, updateTaskDto);
        expect(updatedTask).toEqual(result);
      });
      it('should return a task with deadlinw', async () => {
        const result = task;
        result.user = user;
        updateTaskDto.userId = user.id;
        updateTaskDto.status = ETaskStatus.COMPLETE;
        updateTaskDto.deadlineAt = new Date().toISOString();
        const taskId = 1;
        const updatedTask = await taskService.update(taskId, updateTaskDto);
        expect(updatedTask).toEqual(result);
      });
    });
  });
  describe('delete task by id', () => {
    const deleteResult: DeleteResult = {
      raw: {},
      affected: 1,
    };
    beforeEach(() => {
      deleteTask.mockReturnValue(Promise.resolve(deleteResult));
    });
    it('should return the delete result', async () => {
      const taskId = 1;
      const fetchedTask = await taskService.delete(taskId);
      expect(fetchedTask).toEqual(deleteResult);
    });
  });
  describe('create a task', () => {
    const task = new Task({ id: 1, title: 'title', desc: 'desc' });
    beforeEach(() => {
      save.mockReturnValue(Promise.resolve(task));
    });
    it('should return todo with no deadline', async () => {
      const createDto = new CreateTaskDto();
      createDto.title = 'title';
      createDto.desc = 'desc';
      createDto.deadlineAt = new Date().toISOString();
      const newTask = await taskService.create(createDto);
      expect(newTask).toEqual(task);
    });
    it('should return todo with deadline', async () => {
      const createDto = new CreateTaskDto();
      createDto.title = 'title';
      createDto.desc = 'desc';
      createDto.deadlineAt = null;
      const newTask = await taskService.create(createDto);
      expect(newTask).toEqual(task);
    });
  });
});
