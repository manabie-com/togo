import { Test, TestingModule } from '@nestjs/testing';
import { AuthService } from './auth.service';
import { UserService } from '../users/users.service';
import { JwtService } from '@nestjs/jwt';
import { getRepositoryToken } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { User } from '../users/users.entity';
import bcrypt from 'bcrypt';

describe('AuthService', () => {
  let authService: AuthService;
  let userService: UserService;
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
        AuthService,
        {
          provide: JwtService,
          useValue: {
            sign: jest.fn(),
          },
        },
        UserService,
        {
          provide: getRepositoryToken(User),
          useValue: {
            findOne: jest.fn(),
            save: jest.fn(),
            create: jest.fn(),
          },
        },
      ],
    }).compile();

    authService = module.get<AuthService>(AuthService);
    userService = module.get<UserService>(UserService);
    usersRepository = module.get<Repository<User>>(getRepositoryToken(User));
  });

  it('authService should be defined', () => {
    expect(authService).toBeDefined();
  });

  it('userService should be defined', () => {
    expect(userService).toBeDefined();
  });

  it('userRepository should be defined', () => {
    expect(usersRepository).toBeDefined();
  });

  describe('validateUser', () => {
    it('password should be compared with hashed password', async () => {
      jest.spyOn(usersRepository, 'findOne').mockResolvedValueOnce(mockUser);
      jest.spyOn(bcrypt, 'compare').mockReturnValue(false);
      await authService.validateUser(
        mockUserDto.username,
        mockUserDto.password,
      );

      expect(bcrypt.compare).toBeCalledWith(
        mockUserDto.password,
        mockUser.password,
      );
    });

    it('should return null if credentials is incorrect', async () => {
      jest.spyOn(usersRepository, 'findOne').mockResolvedValueOnce(mockUser);
      jest.spyOn(bcrypt, 'compare').mockReturnValue(false);

      expect(
        authService.validateUser(mockUserDto.username, mockUserDto.password),
      ).resolves.toBe(null);
    });

    it('should return user if credentials is correct', async () => {
      jest.spyOn(usersRepository, 'findOne').mockResolvedValueOnce(mockUser);
      jest.spyOn(bcrypt, 'compare').mockReturnValue(true);

      // eslint-disable-next-line @typescript-eslint/no-unused-vars
      const { password, ...user } = mockUser;

      expect(
        authService.validateUser(mockUserDto.username, mockUserDto.password),
      ).resolves.toStrictEqual(user);
    });
  });
});
