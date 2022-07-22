import { INestApplication } from '@nestjs/common';
import { Test, TestingModule } from '@nestjs/testing';
import * as moment from 'moment';
import { SettingsService } from 'src/settings/settings.service';
import { TodoDto } from 'src/todos/dto/todo.dto';
import { TodoEntity } from 'src/todos/entities/todo.entity';
import { TodosService } from 'src/todos/todos.service';
import { UserEntity } from 'src/users/entities/user.entity';
import { UsersService } from 'src/users/users.service';
import * as request from 'supertest';
import { Repository } from 'typeorm';
import { AppModule } from '../src/app.module';

describe('TodoController (e2e)', () => {
  let app: INestApplication;
  let user: UserEntity;
  let todoRepository: Repository<TodoEntity>;
  let settingsService: SettingsService;
  let todosService: TodosService;

  const createTodoDto = {
    title: "todo1",
    date: moment().format('YYYY-MM-DD'),
  }

  beforeAll(async () => {
    const moduleFixture: TestingModule = await Test.createTestingModule({
      imports: [AppModule],
    }).compile();

    app = moduleFixture.createNestApplication();
    await app.init();

    todoRepository = app.get<Repository<TodoEntity>>('TodoEntityRepository');
    settingsService = app.get<SettingsService>(SettingsService);
    todosService = app.get<TodosService>(TodosService);

    const usersService = app.get<UsersService>(UsersService);
    user = await usersService.create({
      username: 'username1',
      firstName: 'firstName',
      lastName: 'lastName',
    })
  });

  beforeEach(async () => {
    await todoRepository.clear()
  });

  describe('TodoController: create todo (e2e)', () => {
    it('/users/:userId/todos (POST): should create todo success', async () => {
      const todo: TodoDto = await request(app.getHttpServer())
        .post(`/users/${user.id}/todos`)
        .send(createTodoDto)
        .expect(201)
        .then(response => response.body)

      expect(todo.title).toBe(createTodoDto.title)
      expect(todo.date).toBe(createTodoDto.date)
      expect(todo.user.id).toBe(user.id)
    });

    it('/users/:userId/todos (POST): should throw error for unknown property', async () => {
      return request(app.getHttpServer())
        .post(`/users/${user.id}/todos`)
        .send({ ...createTodoDto, unknown: "any" })
        .expect(400)
        .then(response => response.body)
    });

    it('/users/:userId/todos (POST): should throw error when missing required property', async () => {
      return request(app.getHttpServer())
        .post(`/users/${user.id}/todos`)
        .send({ title: createTodoDto.title })
        .expect(400)
        .then(response => response.body)
    });

    it('/users/:userId/todos (POST): should not create todos greater than in setting', async () => {
      const setting = await settingsService.findByUserId(user.id);
      await settingsService.update(setting.id, { todoPerday: 1 });

      await request(app.getHttpServer())
        .post(`/users/${user.id}/todos`)
        .send(createTodoDto)
        .expect(201)

      const body = await request(app.getHttpServer())
        .post(`/users/${user.id}/todos`)
        .send(createTodoDto)
        .expect(400)
        .then(response => response.body)

      expect(body.message).toBe("Exceed number of todos per day")
    });

    it('/users/:userId/todos (POST): should not create todos greater than in setting when request multiple times', async () => {
      const setting = await settingsService.findByUserId(user.id);
      await settingsService.update(setting.id, { todoPerday: 1 });

      await Promise.all([0, 1, 2, 3, 4].map(() => request(app.getHttpServer())
        .post(`/users/${user.id}/todos`)
        .send(createTodoDto)))

      const todos = await todosService.findByUserId(user.id)

      expect(todos.length).toBe(1)
    });

  });
});
