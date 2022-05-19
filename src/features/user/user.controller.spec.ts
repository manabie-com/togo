import { Test, TestingModule } from '@nestjs/testing';
import { CoreModule } from '../../globalModule/core.module';
import { UserController } from './user.controller';
import { UserService } from './user.service';

describe('UserController', () => {
  let controller: UserController;
  let spyUserService: UserService;

  beforeEach(async () => {
    const mockUserService = {
      create: jest.fn()
    };

    const module: TestingModule = await Test.createTestingModule({
      imports: [
        CoreModule
      ],
      controllers: [UserController],
      providers: [
        {
          provide: UserService,
          useValue: mockUserService
        }
      ],
    }).compile();

    controller = module.get<UserController>(UserController);
    spyUserService = module.get<UserService>(UserService);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });

  describe('create function', () => {
    it('should call create function of UserService', async () => {
      const payload = {
        email: 'steven@gmail.com',
        password: '123456'
      };

      await controller.create(payload);
      expect(spyUserService.create).toHaveBeenCalled();
      expect(spyUserService.create).toHaveBeenCalledWith(payload);
    });
  });
});
