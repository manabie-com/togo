import { INestApplication } from '@nestjs/common';
import { Test, TestingModule } from '@nestjs/testing';
import * as moment from 'moment';
import { TodoDto } from 'src/todos/dto/todo.dto';
import { UserEntity } from 'src/users/entities/user.entity';
import { UsersService } from 'src/users/users.service';
import * as request from 'supertest';
import { AppModule } from '../src/app.module';

describe('TodoController (e2e)', () => {
  let app: INestApplication;
  let user: UserEntity;

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

    const userService = app.get<UsersService>(UsersService);
    user = await userService.create({
      username: 'username1',
      firstName: 'firstName',
      lastName: 'lastName',
    })
  });

  beforeEach(async () => {

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

  });
});
