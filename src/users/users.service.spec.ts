import { MongooseModule } from '@nestjs/mongoose';
import { Test, TestingModule } from '@nestjs/testing';
import { constants } from '../constants';
import { CreateUserDto } from './dto/create-user.dto';
import { User, UserSchema } from './schemas/user.schema';
import { UsersService } from './users.service';

describe('UsersService', () => {
  let service: UsersService;

  const mockUser: CreateUserDto = {
    username: 'user1234',
    password: 'user1234',
    displayName: 'user',
  };

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      imports: [
        MongooseModule.forRoot(constants.sandbox),
        MongooseModule.forFeature([{ name: User.name, schema: UserSchema }]),
      ],
      providers: [UsersService],
    }).compile();

    service = module.get<UsersService>(UsersService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });

  it('should create a user', async () => {
    const newUser = await service.create(mockUser);
    expect(newUser).toEqual(
      expect.objectContaining({
        ...mockUser,
        limitedTaskPerDay: 3,
      }),
    );
  });

  it('should find a user', async () => {
    const newUser = await service.findOne(mockUser.username);
    expect(newUser).toEqual(
      expect.objectContaining({
        ...mockUser,
        limitedTaskPerDay: 3,
      }),
    );
  });
});
