import { MongooseModule } from '@nestjs/mongoose';
import { Test, TestingModule } from '@nestjs/testing';
import { constants } from '../constants';
import { CreateUserDto } from './dto/create-user.dto';
import { User, UserSchema } from './schemas/user.schema';
import { UsersController } from './users.controller';
import { UsersService } from './users.service';

describe('UsersController', () => {
  let controller: UsersController;

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
      controllers: [UsersController],
      providers: [UsersService],
    }).compile();

    controller = module.get<UsersController>(UsersController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });

  it('should create a user', async () => {
    const newUser = await controller.create(mockUser);
    expect(newUser).toEqual(
      expect.objectContaining({
        ...mockUser,
        limitedTaskPerDay: 3,
      }),
    );
  });
});
