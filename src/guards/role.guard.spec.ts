import { createMock } from '@golevelup/ts-jest';
import { ExecutionContext } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { Test, TestingModule } from '@nestjs/testing';
import { UserRole } from '../features/user/enum/role.enum';
import { AdminRoleGuard } from './role.guard';

describe('AdminRoleGuard', () => {
  let guard: AdminRoleGuard;
  let jwtService: JwtService;

  beforeEach(async () => {
    const mockJwtService = {
        decode: jest.fn()
    };

    const module: TestingModule = await Test.createTestingModule({
      providers: [
        AdminRoleGuard,
        {
          provide: JwtService,
          useValue: mockJwtService
        }
      ],
    }).compile();

    guard = module.get<AdminRoleGuard>(AdminRoleGuard);
    jwtService = module.get<JwtService>(JwtService);
  });

  it('should be defined', () => {
    expect(guard).toBeDefined();
  });

    it('should call getRequest function', async () => {
        const mockContext = createMock<ExecutionContext>();
        expect(mockContext.switchToHttp().getRequest()).toBeDefined();
    });

    it('should return false when user has a member role', async () => {
        const mockContext = createMock<ExecutionContext>();

        jest.spyOn(jwtService, 'decode').mockReturnValue({
            role: UserRole.MEMBER
        });
        const received = await guard.canActivate(mockContext);
        expect(received).toBe(false);
    });

    it('should return true when user has admin role', async () => {
        const mockContext = createMock<ExecutionContext>();

        jest.spyOn(jwtService, 'decode').mockReturnValue({
            role: UserRole.ADMIN
        });
        const received = await guard.canActivate(mockContext);
        expect(received).toBe(true);
    });
});
