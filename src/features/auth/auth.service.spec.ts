import { ConfigModule, ConfigService } from '@nestjs/config';
import { JwtModule, JwtService } from '@nestjs/jwt';
import { Test, TestingModule } from '@nestjs/testing';
import { CryptoUtil } from '../../utils/crypto.util';
import { UserEntity } from '../user/entities/user.entity';
import { UserRepository } from '../user/user.repository';
import { UserService } from '../user/user.service';
import { AuthService } from './auth.service';

describe('AuthService', () => {
  let service: AuthService;
  let spyUserService: UserService;
  let spyJwtService: JwtService;

  beforeEach(async () => {
    const mockUserService = {
      findOneByEmail: jest.fn(),
      create: jest.fn()
    };

    const mockUserRepository = {
      find: jest.fn(),
      save: jest.fn()
    };

    const mockJWTService = {
      sign: jest.fn().mockReturnValue('fakeJWTToken'),
      verify: jest.fn()
    };

    const module: TestingModule = await Test.createTestingModule({
      imports: [
        JwtModule.registerAsync({
          imports: [ConfigModule],
          inject: [ConfigService],
          useFactory: (configService: ConfigService) => ({
              secret: configService.get('AUTH_JWT_SECRET'),
              signOptions: { expiresIn: configService.get('AUTH_JWT_EXPIRED_TIME') }
          }),
      })
      ],
      providers: [
        AuthService,
        {
          provide: UserService,
          useValue: mockUserService 
        },
        {
          provide: UserRepository,
          useValue: mockUserRepository
        },
        {
          provide: JwtService,
          useValue: mockJWTService
        }
      ],
    })
    .compile();

    service = module.get<AuthService>(AuthService);
    spyUserService = module.get<UserService>(UserService);
    spyJwtService = module.get<JwtService>(JwtService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });

  describe('singin function', () =>{
    it('should call findOneByEmail function of UserService', async () => {  
      const payload = {
        email: 'steven@gmail.com',
        password: '123'
      };

      jest.spyOn(spyUserService, 'findOneByEmail').mockResolvedValue({
        id: 1,
        email: 'steven@gmail.com',
        maxTasks: 5
      } as UserEntity);

      jest.spyOn(CryptoUtil, 'compareHashWithPlainText').mockResolvedValue(true);

      await service.signin(payload);
      expect(spyUserService.findOneByEmail).toHaveBeenCalled();
      expect(spyUserService.findOneByEmail).toHaveBeenCalledWith(payload.email);
    });

    it('should throw an invalid email or password exception when receiving a non-exist user', async () => {  
      const payload = {
        email: 'steven@gmail.com',
        password: '123'
      };

      let errorMessage = '';

      try {
        await service.signin(payload)
      } catch(err) {
        errorMessage = err.message;
      }
  
      expect(errorMessage).toBe(`Email or password is invalid. Please try again.`);
    });

    it('should throw an invalid email or password exception when receiving a wrong password', async () => {  
      const payload = {
        email: 'steven@gmail.com',
        password: '123'
      };

      let errorMessage = '';

      jest.spyOn(spyUserService, 'findOneByEmail').mockResolvedValue({
        id: 1,
        email: 'steven@gmail.com',
      } as UserEntity);

      jest.spyOn(CryptoUtil, 'compareHashWithPlainText').mockResolvedValue(false);

      try {
        await service.signin(payload)
      } catch(err) {
        errorMessage = err.message;
      }
  
      expect(errorMessage).toBe(`Email or password is invalid. Please try again.`);
    });

    it('should return a JWT token', async () => {  
      const payload = {
        email: 'steven@gmail.com',
        password: '123'
      };

      const expectedResult = 'fakeJWTToken';

      jest.spyOn(spyUserService, 'findOneByEmail').mockResolvedValue({
        id: 1,
        email: 'steven@gmail.com',
        maxTasks: 5
      } as UserEntity);

      jest.spyOn(CryptoUtil, 'compareHashWithPlainText').mockResolvedValue(true);

      const token = await service.signin(payload);
      expect(token).toBe(expectedResult);
    });

    it('should return false when verifying an invalid JWT token', async () => {  
      const result = service.verifyToken('invalidToken');
      expect(result).toBe(false);
    });
  });
});
