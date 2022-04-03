import { INestApplication } from '@nestjs/common';
import { Test } from '@nestjs/testing';
import { getRepositoryToken } from '@nestjs/typeorm';
import { ValidationPipe } from 'src/common/pipes/validation.pipe';
import {
  CreateTaskDto,
  CreateToDoDto,
  GetTasksDto,
  UpdateTaskDto,
  UpdateToDoDto,
} from 'src/dto';
import { Task } from 'src/entities/task.entity';
import { ToDoList } from 'src/entities/toDoList.entity';
import { User } from 'src/entities/user.entity';
import { TaskService } from 'src/services/task.service';
import { ToDoService } from 'src/services/todo.service';
import { UserService } from 'src/services/user.service';
import * as request from 'supertest';
import { DeleteResult } from 'typeorm';
import { TaskController } from '../task.controller';

const mockedUser = new User({ id: 1, name: 'NVA', dailyMaxTasks: 2 });
const mockedTask = new Task({
  id: 1,
  title: 'task 1',
  desc: 'desc 1',
  deadlineAt: null,
});
const mockedTodo = new ToDoList({
  id: 1,
  title: 'toto title',
  desc: 'todo desc',
  isDone: false,
});
describe('The TaskController', () => {
  let app: INestApplication;
  let userData: User;
  let taskData: Task;
  let todoData: ToDoList;
  const usersRepository = {
    create: jest.fn().mockResolvedValue(userData),
    save: jest.fn().mockReturnValue(Promise.resolve(userData)),
    findOne: jest.fn().mockResolvedValue(Promise.resolve(userData)),
    findAndCount: jest.fn().mockResolvedValue(Promise.resolve([[userData], 1])),
  };
  const todoRepository = {
    findOne: jest.fn(),
    save: jest.fn(),
    update: jest.fn(),
    count: jest.fn(),
    delete: jest.fn(),
  };
  const taskRepository = {
    findAndCount: jest.fn().mockResolvedValue(Promise.resolve([[], 0])),
    findOne: jest.fn(),
    save: jest.fn(),
    update: jest.fn(),
    count: jest.fn(),
    delete: jest.fn(),
  };
  beforeEach(async () => {
    userData = {
      ...mockedUser,
    };
    taskData = {
      ...mockedTask,
    };
    todoData = {
      ...mockedTodo,
      task: taskData,
    };
    const module = await Test.createTestingModule({
      controllers: [TaskController],
      providers: [
        UserService,
        TaskService,
        ToDoService,
        {
          provide: getRepositoryToken(User),
          useValue: usersRepository,
        },
        {
          provide: getRepositoryToken(ToDoList),
          useValue: todoRepository,
        },
        {
          provide: getRepositoryToken(Task),
          useValue: taskRepository,
        },
      ],
    }).compile();
    app = module.createNestApplication();
    app.useGlobalPipes(new ValidationPipe());
    await app.init();
  });
  describe('get tasks', () => {
    beforeEach(() => {
      taskRepository.findAndCount.mockReturnValue(
        Promise.resolve([[taskData], 1]),
      );
    });
    it('should return data of tasks', () => {
      const expectData = {
        items: [taskData],
        total: 1,
      };
      const queryParameters = new GetTasksDto();
      return request(app.getHttpServer())
        .get('/tasks')
        .query(queryParameters)
        .expect(200)
        .expect(expectData);
    });
  });
  describe('get a task by id', () => {
    describe('task id is not matched', () => {
      beforeEach(() => {
        taskRepository.findOne.mockReturnValue(Promise.resolve(null));
      });
      it('throw error', () => {
        const taskId = 1;
        return request(app.getHttpServer()).get(`/tasks/${taskId}`).expect(404);
      });
    });
    describe('task id is matched', () => {
      beforeEach(() => {
        taskRepository.findOne.mockReturnValue(Promise.resolve(taskData));
      });
      it('should return data of task', () => {
        const taskId = 1;
        return request(app.getHttpServer())
          .get(`/tasks/${taskId}`)
          .expect(200)
          .expect(taskData);
      });
    });
  });
  describe('create a task', () => {
    describe('with invalid params', () => {
      beforeEach(() => {
        taskRepository.findOne.mockReturnValue(Promise.resolve(null));
      });
      it('throw error', () => {
        const createDto = new CreateTaskDto();
        return request(app.getHttpServer())
          .post(`/tasks`)
          .send(createDto)
          .expect(400);
      });
    });
    describe('with valid params', () => {
      beforeEach(() => {
        taskRepository.save.mockReturnValue(Promise.resolve(taskData));
      });
      it('should return data of task', () => {
        const createDto = new CreateTaskDto();
        createDto.deadlineAt = taskData.deadlineAt
          ? taskData.deadlineAt.toISOString()
          : null;
        createDto.title = taskData.title;
        createDto.desc = taskData.desc;
        console.log(createDto);
        return request(app.getHttpServer())
          .post(`/tasks`)
          .send(createDto)
          .expect(201)
          .expect(taskData);
      });
    });
  });
  describe('update task', () => {
    describe('update task with valid parameters', () => {
      beforeEach(() => {
        taskRepository.count.mockReturnValue(Promise.resolve(1));
        taskRepository.findOne.mockReturnValue(Promise.resolve(taskData));
        taskRepository.save.mockReturnValue(Promise.resolve(taskData));
        usersRepository.findOne.mockReturnValue(Promise.resolve(userData));
      });
      it('should response data of task', () => {
        const expectData = {
          ...taskData,
          user: {
            ...userData,
          },
        };
        const taskId = 1;
        const updateDto = new UpdateTaskDto();
        updateDto.userId = userData.id;
        updateDto.title = taskData.title;
        updateDto.desc = taskData.desc;
        return request(app.getHttpServer())
          .put(`/tasks/${taskId}`)
          .send(updateDto)
          .expect(200)
          .expect(expectData);
      });
    });
  });
  describe('delete task', () => {
    describe('task id is matched', () => {
      const deleteResult: DeleteResult = {
        raw: {},
        affected: 1,
      };
      beforeEach(() => {
        taskRepository.findOne.mockReturnValue(Promise.resolve(taskData));
        taskRepository.delete.mockReturnValue(Promise.resolve(deleteResult));
      });
      it('should response delete result', () => {
        const taskId = 1;
        return request(app.getHttpServer())
          .delete(`/tasks/${taskId}`)
          .expect(200)
          .expect(deleteResult);
      });
    });
    describe('task id is not matched', () => {
      beforeEach(() => {
        taskRepository.findOne.mockReturnValue(Promise.resolve(null));
      });
      it('throw error', () => {
        const taskId = 1;
        return request(app.getHttpServer())
          .delete(`/tasks/${taskId}`)
          .expect(404);
      });
    });
  });
  describe('create todo of the task', () => {
    describe('with invalid parameters', () => {
      it('throw error', () => {
        const taskId = 1;
        const createTodoDto = new CreateToDoDto();
        return request(app.getHttpServer())
          .post(`/tasks/${taskId}/to-do`)
          .send(createTodoDto)
          .expect(400);
      });
    });
    describe('task id is not matched', () => {
      beforeEach(() => {
        taskRepository.findOne.mockReturnValue(Promise.resolve(null));
      });
      it('throw error', () => {
        const taskId = 1;
        const createTodoDto = new CreateToDoDto();
        createTodoDto.title = 'title';
        createTodoDto.desc = 'desc';
        return request(app.getHttpServer())
          .post(`/tasks/${taskId}/to-do`)
          .send(createTodoDto)
          .expect(404);
      });
    });
    describe('with valid parameters', () => {
      beforeEach(() => {
        taskRepository.findOne.mockReturnValue(Promise.resolve(taskData));
        todoRepository.save.mockReturnValue(Promise.resolve(todoData));
      });
      it('should return data of todo', () => {
        const expectData = {
          ...todoData,
        };
        const taskId = 1;
        const createTodoDto = new CreateToDoDto();
        createTodoDto.title = todoData.title;
        createTodoDto.desc = todoData.desc;
        return request(app.getHttpServer())
          .post(`/tasks/${taskId}/to-do`)
          .send(createTodoDto)
          .expect(201)
          .expect(expectData);
      });
    });
  });
  describe('update todo of task', () => {
    describe('task id is not matched', () => {
      beforeEach(() => {
        taskRepository.findOne.mockReturnValue(Promise.resolve(null));
      });
      it('throw error', () => {
        const taskId = 1;
        const todoId = 1;
        const updateToDoDto = new UpdateToDoDto();
        updateToDoDto.title = 'title';
        updateToDoDto.desc = 'desc';
        return request(app.getHttpServer())
          .put(`/tasks/${taskId}/to-do/${todoId}`)
          .send(updateToDoDto)
          .expect(404);
      });
    });
    describe('with invalid parameters', () => {
      beforeEach(() => {
        taskRepository.findOne.mockReturnValue(Promise.resolve(taskData));
      });
      it('throw error', () => {
        const taskId = 1;
        const todoId = 1;
        const updateToDoDto = new UpdateToDoDto();
        // updateToDoDto.title = 'title';
        // updateToDoDto.desc = 'desc';
        return request(app.getHttpServer())
          .put(`/tasks/${taskId}/to-do/${todoId}`)
          .send(updateToDoDto)
          .expect(400);
      });
    });
    describe('with valid parameters', () => {
      beforeEach(() => {
        taskRepository.findOne.mockReturnValue(Promise.resolve({ taskData }));
        todoRepository.save.mockReturnValue(Promise.resolve(todoData));
        todoRepository.findOne.mockReturnValue(Promise.resolve(todoData));
      });
      it('throw error', () => {
        const taskId = 1;
        const todoId = 1;
        const updateToDoDto = new UpdateToDoDto();
        updateToDoDto.title = 'title';
        updateToDoDto.desc = 'desc';
        updateToDoDto.isDone = true;
        const expectData = {
          ...todoData,
          title: updateToDoDto.title,
          desc: updateToDoDto.desc,
          isDone: updateToDoDto.isDone,
        };
        return request(app.getHttpServer())
          .put(`/tasks/${taskId}/to-do/${todoId}`)
          .send(updateToDoDto)
          .expect(200)
          .expect(expectData);
      });
    });
  });
});
