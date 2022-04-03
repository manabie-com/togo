import { Test } from '@nestjs/testing';
import { getRepositoryToken } from '@nestjs/typeorm';
import { CreateUserDto, GetUserListDto, UpdateUserDto } from 'src/dto';
import { User } from 'src/entities/user.entity';
import { UserService } from '../user.service';

describe('The UserService', () => {
  let userService: UserService;
  let findOne: jest.Mock;
  let findAndCount: jest.Mock;
  let save: jest.Mock;
  let create: jest.Mock;
  beforeEach(async () => {
    findOne = jest.fn();
    findAndCount = jest.fn();
    save = jest.fn();
    create = jest.fn();
    const module = await Test.createTestingModule({
      providers: [
        UserService,
        {
          provide: getRepositoryToken(User),
          useValue: {
            findOne,
            findAndCount,
            save,
            create,
          },
        },
      ],
    }).compile();
    userService = await module.get(UserService);
  });
  describe('when getting a user by id', () => {
    describe('and the user is matched', () => {
      let user: User;
      beforeEach(() => {
        user = new User({});
        findOne.mockReturnValue(Promise.resolve(user));
      });
      it('should return the user', async () => {
        const userId = 1;
        const fetchedUser = await userService.findOne(userId);
        expect(fetchedUser).toEqual(user);
      });
    });
    describe('and the user is not matched', () => {
      beforeEach(() => {
        findOne.mockReturnValue(Promise.resolve(null));
      });
      it('should return null', async () => {
        const fetchedUser = await userService.findOne(100);
        expect(fetchedUser).toEqual(null);
      });
    });
  });
  describe('when getting user list', () => {
    const user = new User({});
    const result = {
      items: [user],
      total: 0,
    };
    beforeEach(() => {
      findAndCount.mockReturnValue(Promise.resolve([[user], 0]));
    });
    it('should return the user pagination', async () => {
      const queryParameters = new GetUserListDto();
      const usersFetched = await userService.find(queryParameters);
      expect(usersFetched).toEqual(result);
    });
  });
  describe('update user info', () => {
    describe('user is not matched', () => {
      beforeEach(() => {
        findOne.mockReturnValue(Promise.resolve(null));
      });
      it('should be throw an error', async () => {
        const userId = 3;
        const updateUserDto = new UpdateUserDto();
        await expect(
          userService.updateUser(userId, updateUserDto),
        ).rejects.toThrow();
      });
    });
    describe('user is matched', () => {
      const prevUser = new User({ id: 1, name: 'NVA', dailyMaxTasks: 1 });
      const afterUser = new User({ id: 1, name: 'NVB', dailyMaxTasks: 2 });
      beforeEach(() => {
        findOne.mockReturnValue(Promise.resolve(prevUser));
        save.mockReturnValue(Promise.resolve(prevUser));
      });
      it('should return user', async () => {
        const userId = 3;
        const updateUserDto = new UpdateUserDto();
        updateUserDto.dailyMaxTasks = 2;
        updateUserDto.name = 'NVB';
        const userUpdated = await userService.updateUser(userId, updateUserDto);
        expect(userUpdated).toEqual(afterUser);
      });
    });
  });
  describe('create user', () => {
    const user = new User({ name: 'NVA', dailyMaxTasks: 1 });
    beforeEach(() => {
      save.mockReturnValue(Promise.resolve(user));
    });
    it('should return user', async () => {
      const createUserDto = new CreateUserDto();
      createUserDto.name = 'NVA';
      createUserDto.dailyMaxTasks = 1;
      const newUser = await userService.createUser(createUserDto);
      expect(newUser).toEqual(user);
    });
  });
});
