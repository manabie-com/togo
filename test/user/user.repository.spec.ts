import { HttpException, NotFoundException } from '@nestjs/common';
import { Test, TestingModule } from '@nestjs/testing';
import { userId1, userId2 } from '../../src/user/user';
import { UserRepository } from '../../src/user/user.repository';
import { UserService } from '../../src/user/user.service';
import { mockCreateUserDto, mockUser } from './user';

describe('UserRepository', () => {
  let repository: UserRepository;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        UserRepository
      ],
    }).compile();

    repository = module.get<UserRepository>(UserRepository);
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
      expect(await repository.incrementTask(userId2)).toEqual(true);
      expect(incrementTaskSpy).toHaveBeenCalled();
      await expect(repository.incrementTask(userId2)).rejects.toThrow(new HttpException("Task daily limit exceeded", 405));
      expect(incrementTaskSpy).toHaveBeenCalled();
    })
  })
});
