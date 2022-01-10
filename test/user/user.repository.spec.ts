import { HttpException, NotFoundException } from '@nestjs/common';
import { Test, TestingModule } from '@nestjs/testing';
import { mockUsersList, userId1, userId2, userId3 } from './user';
import { UserRepository } from '../../src/user/user.repository';

describe('UserRepository', () => {
  let repository: UserRepository;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        UserRepository
      ],
    }).compile();

    repository = module.get<UserRepository>(UserRepository);
    repository.setUsers(mockUsersList());
  });

  it('should be defined', () => {
    expect(repository).toBeDefined();
  });

  describe('implementTask()', () => {
    it('should throw Task daily limit exceeded for user1', async () => {
      await expect(repository.incrementTask(userId1)).rejects.toThrow(new HttpException("Task daily limit exceeded", 405));
    });

    it('should throw NotFoundException for unknown userId', async () => {
      await expect(repository.incrementTask('anyuserid')).rejects.toThrow(new NotFoundException());
    });

    it('should return true if increment task succeeded and throw when daily task limit exceeded', async () => {
      const incrementTaskSpy = jest.spyOn(repository, 'incrementTask');
      expect(await repository.incrementTask(userId2)).toEqual(true);
      expect(incrementTaskSpy).toHaveBeenCalled();
      await expect(repository.incrementTask(userId2)).rejects.toThrow(new HttpException("Task daily limit exceeded", 405));
      expect(incrementTaskSpy).toHaveBeenCalled();
    })

    it('should update dailyTaskDate and dailyTaskCounter', async () => {
      await repository.incrementTask(userId2);
      const users = await repository.findAll();
      const user2 = users.find((user) => user.id === userId2);
      const now = new Date();
      expect(user2.dailyTaskDate.toISOString().split('T')[0]).toEqual(now.toISOString().split('T')[0]);
      expect(user2.dailyTaskCounter).toEqual(1);
    })

    it('should increment dailyTaskCounter', async () => {
      await repository.incrementTask(userId3);
      const users = await repository.findAll();
      const user3 = users.find((user) => user.id === userId3);
      expect(user3.dailyTaskCounter).toEqual(3);
    })
  })
});
