import { Test, TestingModule } from '@nestjs/testing';
import { UserService } from './user.service';
import { Repository } from 'typeorm';
import { UserEntity } from '../entity/user.entity';
import { getRepositoryToken } from '@nestjs/typeorm';
import { EncryptionUtil } from 'src/common/util/encryption.util';

class EncryptionUtilMock {
  encryptPassword(password: string) {
    return password;
  }
}

describe('UserService', () => {
  let userService: UserService;
  let repositoryMock: MockType<Repository<UserEntity>>;

  beforeEach(async () => {
    const encryptionUtilProvider = {
      provide: EncryptionUtil,
      useClass: EncryptionUtilMock,
    };
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        UserService,
        {
          provide: getRepositoryToken(UserEntity),
          useFactory: repositoryMockFactory,
        },
        encryptionUtilProvider,
      ],
    }).compile();
    userService = module.get<UserService>(UserService);
    repositoryMock = module.get(getRepositoryToken(UserEntity));
  });

  it('UserService - should be defined', () => {
    expect(userService).toBeDefined();
  });

  it('#Get user by id should work', async () => {
    const userMockValues = {
      id: '61f320f3-6b7a-432f-bb1c-80a48ec6ffa7',
      username: 'togo',
      password: '123456',
    };

    repositoryMock.findOne.mockReturnValue(userMockValues);

    const user = await userService.getUserById(
      '61f320f3-6b7a-432f-bb1c-80a48ec6ffa7',
    );

    expect(user.id).toEqual(userMockValues.id);
    expect(repositoryMock.findOne).toHaveBeenCalledWith(user.id);
  });

  it('#Get user by username should work', async () => {
    const userMockValues = {
      id: '61f320f3-6b7a-432f-bb1c-80a48ec6ffa7',
      username: 'togo',
      password: '123456',
    };

    repositoryMock.findOne.mockReturnValue(userMockValues);

    const user = await userService.getUserByUsername('togo');

    expect(user.id).toEqual(userMockValues.id);
    expect(repositoryMock.findOne).toHaveBeenCalledWith({
      username: user.username,
    });
  });

  it('#Create user should work', async () => {
    const userMockValues = {
      id: '61f320f3-6b7a-432f-bb1c-80a48ec6ffa7',
      username: 'togo',
      password: '123456',
    };

    repositoryMock.save.mockReturnValue(userMockValues);

    const user = await userService.createUser(userMockValues);

    expect(user.id).toEqual(userMockValues.id);
    expect(repositoryMock.save).toHaveBeenCalled();
  });
});

export type MockType<T> = {
  // eslint-disable-next-line @typescript-eslint/ban-types
  [P in keyof T]?: jest.Mock<{}>;
};

export const repositoryMockFactory: () => MockType<Repository<any>> = jest.fn(
  () => ({
    findOne: jest.fn((entity) => entity),
    save: jest.fn((entity) => entity),
  }),
);
