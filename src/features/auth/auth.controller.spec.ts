import { Test, TestingModule } from '@nestjs/testing';
import { AuthController } from './auth.controller';
import { AuthService } from './auth.service';

describe('TaskController', () => {
  let controller: AuthController;
  let spyService: AuthService;

  beforeEach(async () => {
    const mockUserResult = {
        id: 1,
        email: 'steven@gmail.com',
        maxTasks: 5
    };

    const mockTokenResult = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InN0ZXZlbkBnbWFpbC5jb20iLCJpYXQiOjE2NTI4Nzk4MDcsImV4cCI6MTY1Mjg4MzQwN30.gk-xt7PvmAHmR7tVM1kMCPLu5kXb8lWBpV20QD1GyLs';

    const mockAuthService = {
        signup: jest.fn(() => mockUserResult),
        signin: jest.fn(() => mockTokenResult)
      };

    const module: TestingModule = await Test.createTestingModule({
      controllers: [AuthController],
      providers: [
        {
          provide: AuthService,
          useValue: mockAuthService
        }
      ],
    }).compile();

    controller = module.get<AuthController>(AuthController);
    spyService = module.get<AuthService>(AuthService);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });

  describe('signin function', () => {
    it('should call signin function of AuthService', async () => {  
        const payload = {
          email: 'steven2@gmail.com',
          password: '123',
        };
    
        await controller.signin(payload);
        expect(spyService.signin).toHaveBeenCalled();
        expect(spyService.signin).toHaveBeenCalledWith(payload);
    });
  });
});
