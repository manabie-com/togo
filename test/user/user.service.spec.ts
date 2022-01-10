import { Test, TestingModule } from '@nestjs/testing';
import { UserRepository } from '../../src/user/user.repository';
import { UserService } from '../../src/user/user.service';
import { mockCreateUserDto, mockUser } from './user';

describe('UserService', () => {
  let service: UserService;
  let repository: UserRepository;

  let mockRepository = {
    create: jest.fn(),
    findAll: jest.fn(),
    incrementTask: jest.fn()
  };

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        UserService,
        {
          provide: UserRepository,
          useValue: mockRepository
        }
      ],
    }).compile();

    service = module.get<UserService>(UserService);
    repository = module.get<UserRepository>(UserRepository);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });

  describe('create()', () => {
    it('should call UserRepository create with correct values', async () => {
      const createSpy = jest.spyOn(repository, 'create');
      const mockParam = mockCreateUserDto();
      await service.create(mockParam);
      expect(createSpy).toHaveBeenCalledWith(mockParam);
    })

    it('should throw if UserRepository create throws', async () => {
      jest.spyOn(repository, 'create').mockRejectedValueOnce(new Error());
      await expect(service.create(mockCreateUserDto())).rejects.toThrow(new Error());
    })

    it('should return a User on success', async () => {
      const mockReturn = mockUser();
      jest.spyOn(repository, 'create').mockResolvedValueOnce(mockReturn);
      const response = await service.create(mockCreateUserDto());
      expect(response).toEqual(mockReturn);
    })
  })

  describe('findAll()', () => {
    it('should call UserRepository find all', async () => {
      const findSpy = jest.spyOn(repository, 'findAll');
      await service.findAll();
      expect(findSpy).toHaveBeenCalled();
    })

    it('should throw if UserRepository find all throws', async () => {
      jest.spyOn(repository, 'findAll').mockRejectedValueOnce(new Error());
      await expect(service.findAll()).rejects.toThrow(new Error())
    })

    it('should return Users on success', async () => {
      const mockReturn = [mockUser()]
      jest.spyOn(repository, 'findAll').mockResolvedValueOnce(mockReturn);
      const response = await service.findAll();
      expect(response).toEqual(mockReturn);
    })
  })

  describe('incrementTask()', () => {
    it('should call UserRepository incrementTask', async () => {
      const findSpy = jest.spyOn(repository, 'incrementTask');
      await service.incrementTask('anyuserid');
      expect(findSpy).toHaveBeenCalled();
    })

    it('should throw if UserRepository incrementTask throws', async () => {
      jest.spyOn(repository, 'incrementTask').mockRejectedValueOnce(new Error());
      await expect(service.incrementTask('anyuserid')).rejects.toThrow(new Error())
    })

    it('should return Users on success', async () => {
      const mockReturn = true;
      jest.spyOn(repository, 'incrementTask').mockResolvedValueOnce(mockReturn);
      const response = await service.incrementTask('anyuserid');
      expect(response).toEqual(mockReturn);
    })
  })
});
