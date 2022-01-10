import { Test, TestingModule } from '@nestjs/testing';
import { UserController } from '../../src/user/user.controller';
import { UserService } from '../../src/user/user.service';
import { mockCreateUserDto, mockUser } from './user';

describe('UserController', () => {
  let controller: UserController;
  let service: UserService;

  const mockService = {
    create: jest.fn(),
    findAll: jest.fn()
  }

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [UserController],
      providers: [
        {
          provide: UserService,
          useValue: mockService
        }
      ]
    }).compile();

    controller = module.get<UserController>(UserController);
    service = module.get<UserService>(UserService);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });

  describe('create()', () => {

    it('should have create function', () => {
      expect(controller.create).toBeDefined();
    });

    it('should call UserService create with correct values', async () => {
      const createSpy = jest.spyOn(service, 'create');
      const mockParam = mockCreateUserDto();
      await controller.create(mockParam);
      expect(createSpy).toHaveBeenCalledWith(mockParam);
    })

    it('should throw if UserService create throws', async () => {
      jest.spyOn(service, 'create').mockRejectedValueOnce(new Error());
      await expect(controller.create(mockCreateUserDto())).rejects.toThrow(new Error());
    })

    it('should return a User on success', async () => {
      const mockReturn = mockUser();
      jest.spyOn(service, 'create').mockResolvedValueOnce(mockReturn);
      const response = await controller.create(mockCreateUserDto())
      expect(response).toEqual(mockReturn);
    })
  })

  describe('findAll()', () => {
    it('should call UserService find all', async () => {
      const findSpy = jest.spyOn(service, 'findAll');
      await controller.findAll();
      expect(findSpy).toHaveBeenCalled()
    })

    it('should throw if UserService find all throws', async () => {
      jest.spyOn(service, 'findAll').mockRejectedValueOnce(new Error());
      await expect(controller.findAll()).rejects.toThrow(new Error());
    })

    it('should return Users on success', async () => {
      const mockReturn = [
        mockUser()
      ]
      jest.spyOn(service, 'findAll').mockResolvedValueOnce(mockReturn);
      const response = await controller.findAll();
      expect(response).toEqual(mockReturn)
    })
  })
  
});
