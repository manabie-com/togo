import { INestApplication } from '@nestjs/common';
import { Test } from '@nestjs/testing';
import { getRepositoryToken } from '@nestjs/typeorm';
import { ValidationPipe } from 'src/common/pipes/validation.pipe';
import { User } from 'src/entities/user.entity';
import { UserService } from 'src/services/user.service';
import { UserController } from '../user.controller';
import * as request from 'supertest';
import { CreateUserDto, GetUserListDto, UpdateUserDto } from 'src/dto';

const mockedUser = new User({ id: 1, name: 'NVA', dailyMaxTasks: 2 });

describe('The AuthenticationController', () => {
  let app: INestApplication;
  let userData: User;
  let usersRepository: Record<string, any>;
  beforeEach(async () => {
    userData = {
      ...mockedUser,
    };
    usersRepository = {
      create: jest.fn().mockResolvedValue(userData),
      save: jest.fn().mockReturnValue(Promise.resolve(userData)),
      findOne: jest.fn().mockResolvedValue(Promise.resolve(userData)),
      findAndCount: jest
        .fn()
        .mockResolvedValue(Promise.resolve([[userData], 1])),
    };

    const module = await Test.createTestingModule({
      controllers: [UserController],
      providers: [
        UserService,
        {
          provide: getRepositoryToken(User),
          useValue: usersRepository,
        },
      ],
    }).compile();
    app = module.createNestApplication();
    app.useGlobalPipes(new ValidationPipe());
    await app.init();
  });
  describe('when creating with valid data', () => {
    it('should response with data of user', () => {
      const expectedData = {
        ...userData,
      };
      const createUserDto = new CreateUserDto();
      createUserDto.dailyMaxTasks = mockedUser.dailyMaxTasks;
      createUserDto.name = mockedUser.name;
      return request(app.getHttpServer())
        .post('/users')
        .send(createUserDto)
        .expect(201)
        .expect(expectedData);
    });
  });
  describe('when creating with invalid data', () => {
    it('throw error', () => {
      const createUserDto = new CreateUserDto();
      createUserDto.dailyMaxTasks = mockedUser.dailyMaxTasks;
      createUserDto.name = null;
      return request(app.getHttpServer())
        .post('/users')
        .send(createUserDto)
        .expect(400);
    });
  });
  describe('update user by id', () => {
    it('should return data of user', () => {
      const updateUserDto = new UpdateUserDto();
      updateUserDto.dailyMaxTasks = mockedUser.dailyMaxTasks;
      updateUserDto.name = 'NVB';
      const userId = 1;
      const result = {
        ...userData,
        name: updateUserDto.name,
        dailyMaxTasks: updateUserDto.dailyMaxTasks,
      };
      return request(app.getHttpServer())
        .put(`/users/${userId}`)
        .send(updateUserDto)
        .expect(200)
        .expect(result);
    });
  });
  describe('get user by id', () => {
    describe('user not found', () => {
      beforeEach(() => {
        usersRepository['findOne'].mockReturnValue(Promise.resolve(null));
      });
      it('throw error', () => {
        const userId = 1;
        return request(app.getHttpServer()).get(`/users/${userId}`).expect(404);
      });
    });
    describe('user found', () => {
      it('should return data of user', () => {
        const userId = 1;
        return request(app.getHttpServer())
          .get(`/users/${userId}`)
          .expect(200)
          .expect(userData);
      });
    });
  });
  describe('get users', () => {
    it('should return data of users', () => {
      const expectData = {
        items: [userData],
        total: 1,
      };
      const queryParameters = new GetUserListDto();
      return request(app.getHttpServer())
        .get('/users')
        .query(queryParameters)
        .expect(200)
        .expect(expectData);
    });
  });
});
