import { Test, TestingModule } from '@nestjs/testing';
import { INestApplication, ValidationPipe } from '@nestjs/common';
import supertest from 'supertest';
import { AppModule } from '../src/app.module';
import { TodoService } from '../src/todo/todo.service';
import { getRepositoryToken } from '@nestjs/typeorm';
import { User } from '../src/users/users.entity';
import { Todo, TodoStatus } from '../src/todo/todo.entity';
import bcrypt from 'bcrypt';

describe('AppController (e2e)', () => {
  let app: INestApplication;

  const mockUserDto = {
    username: 'john',
    password: 'john',
    limitPerDay: 5,
  };
  const mockUser = {
    username: 'john',
    password: 'hashed password',
    limitPerDay: 5,
    todos: [
      {
        id: 'a',
        content: 'a',
        status: TodoStatus.ACTIVE,
      },
      {
        id: 'b',
        content: 'b',
        status: TodoStatus.COMPLETED,
      },
    ],
    createAt: new Date(),
    updateAt: new Date(),
  };
  let jwtToken;

  const todosRepository = {
    find: jest.fn(),
    count: jest.fn(),
    update: jest.fn(),
    delete: jest.fn(),
  };
  const usersRepository = {
    findOne: jest.fn(),
    save: jest.fn(),
    create: jest.fn(),
  };
  const todoService = {
    getTodos: jest.fn((username, status) =>
      mockUser.todos.filter((todo) => (status ? todo.status === status : true)),
    ),
    setTodoStatus: jest.fn(),
    setManyTodoStatus: jest.fn(),
    deleteTodoById: jest.fn(),
    deleteAllTodos: jest.fn(),
  };

  beforeAll(async () => {
    const moduleFixture: TestingModule = await Test.createTestingModule({
      imports: [AppModule],
    })
      .overrideProvider(getRepositoryToken(Todo))
      .useValue(todosRepository)
      .overrideProvider(getRepositoryToken(User))
      .useValue(usersRepository)
      .overrideProvider(TodoService)
      .useValue(todoService)
      .compile();

    app = moduleFixture.createNestApplication();
    app.useGlobalPipes(new ValidationPipe());
    await app.init();
  });

  afterAll(async () => {
    await Promise.all([app.close()]);
  });

  describe('/user (POST)', () => {
    it('should not create user with incorrect value', () => {
      return supertest(app.getHttpServer())
        .post('/user')
        .send({
          ...mockUserDto,
          limitPerDay: 'abc',
        })
        .expect(400);
    });

    it('should create new user', () => {
      return supertest(app.getHttpServer())
        .post('/user')
        .send(mockUserDto)
        .expect(201);
    });
  });

  describe('/auth/login (POST)', () => {
    it('should return JWT token', async () => {
      jest.spyOn(usersRepository, 'findOne').mockReturnValueOnce(mockUser);
      jest.spyOn(bcrypt, 'compare').mockReturnValueOnce(true);

      const response = await supertest(app.getHttpServer())
        .post('/auth/login')
        .send(mockUserDto)
        .set('Accept', 'application/json')
        .expect('Content-Type', /json/)
        .expect(201);

      jwtToken = response.body.access_token;
      expect(jwtToken).toMatch(
        /^[A-Za-z0-9-_=]+\.[A-Za-z0-9-_=]+\.?[A-Za-z0-9-_.+/=]*$/,
      );
    });
  });

  describe('/tasks (GET)', () => {
    it('should return all todos', async () => {
      jest.spyOn(usersRepository, 'findOne').mockReturnValueOnce(mockUser);

      const response = await supertest(app.getHttpServer())
        .get('/tasks')
        .set('Authorization', 'Bearer ' + jwtToken)
        .set('Accept', 'application/json')
        .expect('Content-Type', /json/)
        .expect(200);
      expect(response.body.data).toStrictEqual(mockUser.todos);
    });

    it('should return active todos', async () => {
      jest.spyOn(usersRepository, 'findOne').mockReturnValueOnce(mockUser);

      const response = await supertest(app.getHttpServer())
        .get('/tasks?status=ACTIVE')
        .set('Authorization', 'Bearer ' + jwtToken)
        .set('Accept', 'application/json')
        .expect('Content-Type', /json/)
        .expect(200);
      expect(response.body.data).toStrictEqual([mockUser.todos[0]]);
    });
  });

  describe('/tasks (PUT)', () => {
    it('should update todo status', async () => {
      await supertest(app.getHttpServer())
        .put('/tasks')
        .set('Authorization', 'Bearer ' + jwtToken)
        .send({ id: mockUser.todos[0].id, status: TodoStatus.COMPLETED })
        .expect(200);
      expect(todoService.setTodoStatus).toHaveBeenCalledWith(
        mockUser.todos[0].id,
        TodoStatus.COMPLETED,
      );
    });

    it('should not update todo status with incorrect value', () => {
      return supertest(app.getHttpServer())
        .put('/tasks')
        .set('Authorization', 'Bearer ' + jwtToken)
        .send({ id: mockUser.todos[0].id, status: 'status' })
        .expect(400);
    });
  });

  describe('/many-tasks (PUT) ', () => {
    it('should update many todos status', async () => {
      await supertest(app.getHttpServer())
        .put('/many-tasks')
        .set('Authorization', 'Bearer ' + jwtToken)
        .send({
          ids: mockUser.todos.map(({ id }) => id),
          status: TodoStatus.COMPLETED,
        })
        .expect(200);

      expect(todoService.setManyTodoStatus).toHaveBeenCalledWith(
        mockUser.todos.map(({ id }) => id),
        TodoStatus.COMPLETED,
      );
    });

    it('should not update many todos status with incorrect value', () => {
      return supertest(app.getHttpServer())
        .put('/many-tasks')
        .set('Authorization', 'Bearer ' + jwtToken)
        .send({
          ids: mockUser.todos.map(({ id }) => id),
          status: 'status',
        })
        .expect(400);
    });
  });

  describe('/tasks (DELETE)', () => {
    it('should delete todo with id', async () => {
      await supertest(app.getHttpServer())
        .delete(`/tasks/${mockUser.todos[0].id}`)
        .set('Authorization', 'Bearer ' + jwtToken)
        .expect(200);

      expect(todoService.deleteTodoById).toHaveBeenCalledWith(
        mockUser.todos[0].id,
      );
    });

    it('should delete all todos', async () => {
      await supertest(app.getHttpServer())
        .delete(`/tasks`)
        .set('Authorization', 'Bearer ' + jwtToken)
        .expect(200);

      expect(todoService.deleteAllTodos).toHaveBeenCalledWith(
        mockUser.username,
      );
    });
  });
});
