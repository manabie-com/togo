import { INestApplication, Logger } from '@nestjs/common';
import { Test } from '@nestjs/testing';
import { getRepositoryToken } from '@nestjs/typeorm';
import * as cookieParser from 'cookie-parser';
import { AllExceptionFilter } from 'src/common/exceptions/exception.filter';
import { TransformInterceptor } from 'src/common/interceptor/transform.interceptor';
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
    app.use(cookieParser());
    app.setGlobalPrefix('api');
    const logger = new Logger();
    app.useGlobalFilters(new AllExceptionFilter(logger));
    app.useGlobalInterceptors(new TransformInterceptor());
    app.enableCors();
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
        .get('/api/tasks')
        .query(queryParameters)
        .then((result) => {
          expect(result.statusCode).toEqual(200);
          expect(result.body.success).toEqual(true);
          expect(result.body.data).toEqual(expectData);
        });
    });
  });
  describe('get a task by id', () => {
    describe('task id is not matched', () => {
      beforeEach(() => {
        taskRepository.findOne.mockReturnValue(Promise.resolve(null));
      });
      it('throw error', () => {
        const taskId = 1;
        return request(app.getHttpServer())
          .get(`/api/tasks/${taskId}`)
          .then((result) => {
            expect(result.statusCode).toEqual(200);
            expect(result.body.success).toEqual(false);
          });
      });
    });
    describe('task id is matched', () => {
      beforeEach(() => {
        taskRepository.findOne.mockReturnValue(Promise.resolve(taskData));
      });
      it('should return data of task', () => {
        const taskId = 1;
        return request(app.getHttpServer())
          .get(`/api/tasks/${taskId}`)
          .then((result) => {
            expect(result.statusCode).toEqual(200);
            expect(result.body.success).toEqual(true);
            expect(result.body.data).toEqual(taskData);
          });
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
          .post(`/api/tasks`)
          .send(createDto)
          .then((result) => {
            expect(result.statusCode).toEqual(200);
            expect(result.body.success).toEqual(false);
          });
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
          .post(`/api/tasks`)
          .send(createDto)
          .expect(201)
          .expect({
            message: null,
            data: taskData,
            success: true,
          });
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
          .put(`/api/tasks/${taskId}`)
          .send(updateDto)
          .then((result) => {
            expect(result.statusCode).toEqual(200);
            expect(result.body.success).toEqual(true);
            expect(result.body.data).toEqual(expectData);
          });
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
          .delete(`/api/tasks/${taskId}`)
          .then((result) => {
            expect(result.statusCode).toEqual(200);
            expect(result.body.success).toEqual(true);
            expect(result.body.data).toEqual(deleteResult);
          });
      });
    });
    describe('task id is not matched', () => {
      beforeEach(() => {
        taskRepository.findOne.mockReturnValue(Promise.resolve(null));
      });
      it('throw error', () => {
        const taskId = 1;
        return request(app.getHttpServer())
          .delete(`/api/tasks/${taskId}`)
          .then((result) => {
            expect(result.statusCode).toEqual(200);
            expect(result.body.success).toEqual(false);
          });
      });
    });
  });
  describe('create todo of the task', () => {
    describe('with invalid parameters', () => {
      it('throw error', () => {
        const taskId = 1;
        const createTodoDto = new CreateToDoDto();
        return request(app.getHttpServer())
          .post(`/api/tasks/${taskId}/to-do`)
          .send(createTodoDto)
          .then((result) => {
            expect(result.statusCode).toEqual(200);
            expect(result.body.success).toEqual(false);
          });
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
          .post(`/api/tasks/${taskId}/to-do`)
          .send(createTodoDto)
          .then((result) => {
            expect(result.statusCode).toEqual(200);
            expect(result.body.success).toEqual(false);
          });
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
          .post(`/api/tasks/${taskId}/to-do`)
          .send(createTodoDto)
          .then((result) => {
            expect(result.statusCode).toEqual(201);
            expect(result.body.success).toEqual(true);
            expect(result.body.data).toEqual(expectData);
          });
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
          .put(`/api/tasks/${taskId}/to-do/${todoId}`)
          .send(updateToDoDto)
          .then((result) => {
            expect(result.statusCode).toEqual(200);
            expect(result.body.success).toEqual(false);
          });
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
        return request(app.getHttpServer())
          .put(`/api/tasks/${taskId}/to-do/${todoId}`)
          .send(updateToDoDto)
          .then((result) => {
            expect(result.statusCode).toEqual(200);
            expect(result.body.success).toEqual(false);
          });
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
          .put(`/api/tasks/${taskId}/to-do/${todoId}`)
          .send(updateToDoDto)
          .then((result) => {
            expect(result.statusCode).toEqual(200);
            expect(result.body.success).toEqual(true);
            expect(result.body.data).toEqual(expectData);
          });
      });
    });
  });
});
