import { createMock } from '@golevelup/ts-jest';
import { ExecutionContext } from '@nestjs/common';
import { Reflector } from '@nestjs/core';
import { Test, TestingModule } from '@nestjs/testing';
import { AuthService } from '../features/auth/auth.service';
import { AuthGuard } from './auth.guard';

describe('AuthGuard', () => {
  let guard: AuthGuard;
  let reflector: Reflector;
  let authService: AuthService;

  beforeEach(async () => {
    const mockAuthService = {
        verifyToken: jest.fn(),
    };

    const mockReflector = {
        get: jest.fn(),
    };

    const module: TestingModule = await Test.createTestingModule({
      providers: [
          AuthGuard,
        {
          provide: AuthService,
          useValue: mockAuthService
        },
        {
            provide: Reflector,
            useValue: mockReflector
        }
      ],
    }).compile();

    guard = module.get<AuthGuard>(AuthGuard);
    reflector = module.get<Reflector>(Reflector)
    authService = module.get<AuthService>(AuthService);
  });

  it('should be defined', () => {
    expect(guard).toBeDefined();
  });

    it('should return true when context is public', async () => {
        const mockContext = createMock<ExecutionContext>();
        jest.spyOn(reflector, 'get').mockReturnValue(true);

        const received = await guard.canActivate(mockContext);
        expect(received).toBe(true);
    });

    it('should call getRequest function', async () => {
        const mockContext = createMock<ExecutionContext>();
        expect(mockContext.switchToHttp().getRequest()).toBeDefined();
    });

    it('should return false when token is invalid', async () => {
        const mockContext = createMock<ExecutionContext>();

        jest.spyOn(authService, 'verifyToken').mockReturnValue(false);
        const received = await guard.canActivate(mockContext);
        expect(received).toBe(false);
    });

    it('should return true when token is valid', async () => {
        const mockContext = createMock<ExecutionContext>();

        jest.spyOn(authService, 'verifyToken').mockReturnValue(true);
        const received = await guard.canActivate(mockContext);
        expect(received).toBe(true);
    });
});
