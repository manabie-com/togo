import { Test, TestingModule } from '@nestjs/testing';
import { Repository } from 'typeorm';
import { UserEntity } from './entities/user.entity';
import { UsersService } from './users.service';

describe('UsersService', () => {
  let service: UsersService;
  let usersRepository: Repository<UserEntity>;

  beforeEach(async () => {
    usersRepository = jest.mock as any;

    const module: TestingModule = await Test.createTestingModule({
      providers: [
        {
          provide: 'UserEntityRepository',
          useValue: usersRepository
        },
        UsersService],
    }).compile();

    service = module.get<UsersService>(UsersService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });
});
