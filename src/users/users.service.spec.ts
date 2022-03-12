import { HttpException } from '@nestjs/common';
import { Test, TestingModule } from '@nestjs/testing';
import { getRepositoryToken } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { UserService } from './users.service';
import { User } from './users.entity';
import bcrypt from 'bcrypt';
import { saltRounds } from '../auth/constants';

describe('UsersService', () => {
  let service: UserService;
  let usersRepository: Repository<User>;

  const mockUserDto = {
    username: 'john',
    password: 'john',
    limitPerDay: 5,
  };
  const mockUser = {
    username: 'john',
    password: 'hashed password',
    limitPerDay: 5,
    todos: [],
    createAt: new Date(),
    updateAt: new Date(),
  };

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        UserService, {
          provide: getRepositoryToken(User),
          useValue: {
            findOne: jest.fn(),
            save: jest.fn(),
            create: jest.fn(),
          },
        }
      ],
    }).compile();

    service = module.get<UserService>(UserService);
    usersRepository = module.get<Repository<User>>(getRepositoryToken(User));
  });

  it('service should be defined', () => {
    expect(service).toBeDefined();
  });

  it('userRepository should be defined', () => {
    expect(usersRepository).toBeDefined();
  });

  describe('createUser', () => {
    jest.spyOn(bcrypt, 'hash').mockReturnValue('hashed password');

    it('password should be hashed', async () => {
      await service.createUser(mockUserDto);

      expect(bcrypt.hash).toHaveBeenCalledWith('john', saltRounds);
    });

    it('user should be created with correct params', async () => {
      jest.spyOn(usersRepository, 'create').mockReturnValueOnce(mockUser);
      await service.createUser(mockUserDto);

      expect(usersRepository.save).toHaveBeenCalledWith(mockUser);
    });

    it('user should not be created with existing username', async () => {
      jest.spyOn(usersRepository, 'findOne').mockResolvedValueOnce(mockUser);

      expect(service.createUser(mockUserDto)).rejects.toThrow(HttpException);
    });
  });

  describe('findUser', () => {
    it('findUser should return the user', async () => {
      jest.spyOn(usersRepository, 'findOne').mockResolvedValueOnce(mockUser);
      expect(service.findUser(mockUserDto.username)).resolves.toStrictEqual(
        mockUser,
      );
    });
  });

  describe('createTodo', () => {
    it('should throw error if daily limit is exceeded', async () => {
      jest.spyOn(service, 'findUser').mockResolvedValueOnce(mockUser);

      expect(service.createTodo('john', 'todo', 5)).rejects.toThrow(
        HttpException,
      );
    });

    it('todo should be added to user', async () => {
      jest.spyOn(service, 'findUser').mockResolvedValueOnce(mockUser);
      await service.createTodo('john', 'todo', 0);

      expect(usersRepository.save).toBeCalledWith({
        ...mockUser,
        todos: [{ content: 'todo' }],
      });
    });
  });
});
