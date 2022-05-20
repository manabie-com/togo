import { Test, TestingModule } from '@nestjs/testing';
import { CryptoUtil } from '../../utils/crypto.util';
import { UserEntity } from './entities/user.entity';
import { UserRepository } from './user.repository';
import { UserService } from './user.service';

describe('UserService', () => {
  let service: UserService;

  beforeEach(async () => {
    const mockUserRepository = {
      find: jest.fn(),
      save: jest.fn(x => x)
    };

    const module: TestingModule = await Test.createTestingModule({
      providers: [
        UserService,
        {
          provide: UserRepository,
          useValue: mockUserRepository
        }
      ],
    }).compile();

    service = module.get<UserService>(UserService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });

  describe('create', () => {
    it('should call findOneByEmail function', async () => {
      const payload = {
        email: 'steven@gmail.com',
        password: '123',
        maxTasks: 5
      };

      const spyService = jest.spyOn(service, 'findOneByEmail').mockResolvedValue(null)

      await service.create(payload);

      expect(spyService).toHaveBeenCalled();
      expect(spyService).toHaveBeenCalledWith(payload.email);
    });

    it('should throw not found user exception', async () => {
      const payload = {
        email: 'steven@gmail.com',
        password: '123',
        maxTasks: 5
      };
      let errorMessage = '';

      jest.spyOn(service, 'findOneByEmail').mockResolvedValue({
        email: 'steven@gmail.com'
      } as UserEntity)

      try {
        await service.create(payload);
      } catch(err) {
        errorMessage = err.message;
      }

      expect(errorMessage).toBe(`User with email ${payload.email} already existed`);
    });

    it('should create and return a new user record', async () => {
      const payload = {
        email: 'steven@gmail.com',
        password: '123',
        maxTasks: 5
      };

      const expectedResult = {
        email: 'steven@gmail.com',
        password: 'hash_password_value',
        maxTasks: 5
      };

      jest.spyOn(service, 'findOneByEmail').mockResolvedValue(null);
      jest.spyOn(CryptoUtil, 'hash').mockResolvedValue('hash_password_value')

      const received = await service.create(payload);
      expect(received).toEqual(expectedResult);
    });
  })

});
