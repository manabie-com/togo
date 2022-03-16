import { Test, TestingModule } from '@nestjs/testing';
import { PrismaService } from './prisma.service';
import { UserResolver } from './resolvers.user';
import { CreateUserInput } from './dtos/create-user.input';
import { LoginUserInput } from './dtos/login-user.input';

describe('UserResolver', () => {
  let resolver: UserResolver;
  const createUserInput: CreateUserInput = {
    email: 'john.doe@email.com',
    password: '123456',
    name: 'John Doe',
    maxJob: 5,
  };

  const loginUserInput: LoginUserInput = {
    email: 'john.doe@email.com',
    password: '123456',
  };

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [UserResolver, PrismaService],
    }).compile();

    resolver = module.get<UserResolver>(UserResolver);
  });

  it('should be defined', () => {
    expect(resolver).toBeDefined();
  });

  describe('signup', () => {
    it('should signup and return info token', async () => {
      const user = await resolver.signup(createUserInput);
      const returnData = { ...createUserInput, ...user };

      expect(user).toEqual(returnData);
    });
  });

  describe('login', () => {
    it('should login and return info token', async () => {
      const user = await resolver.login(loginUserInput);
      const returnData = { ...loginUserInput, ...user };

      expect(user).toEqual(returnData);
    });
  });
});
