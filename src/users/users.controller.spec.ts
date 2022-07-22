import { Test, TestingModule } from '@nestjs/testing';
import { Repository } from 'typeorm';
import { UserEntity } from './entities/user.entity';
import { UsersController } from './users.controller';
import { UsersService } from './users.service';

describe('UsersController', () => {
  let controller: UsersController;
  let usersRepository: Repository<UserEntity>;

  beforeEach(async () => {
    usersRepository = jest.mock as any;

    const module: TestingModule = await Test.createTestingModule({
      controllers: [UsersController],
      providers: [
        {
          provide: 'UserEntityRepository',
          useValue: usersRepository
        },
        UsersService],
    }).compile();

    controller = module.get<UsersController>(UsersController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
