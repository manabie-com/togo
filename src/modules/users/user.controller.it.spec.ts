import { Test, TestingModule } from '@nestjs/testing';
import { UserController } from './user.controller';
import { TestUtils } from '@test/test.utils';
import { DatabaseModule } from '@test/database';
import { DatabaseService } from '@test/database/database.service';
import { JWTPayload } from '@modules/auth/dto';
import { User } from './entities/user.entity';

describe('Test User Controller', () => {
  let controller: UserController;
  let module: TestingModule;
  let testUtils: TestUtils;

  beforeEach(async () => {
    const testModule = await Test.createTestingModule({
      imports: [DatabaseModule],
      providers: [DatabaseService, TestUtils],
    }).compile();

    testUtils = testModule.get<TestUtils>(TestUtils);

    await testUtils.reloadFixtures();

    module = await Test.createTestingModule({
      // imports: [...testUtils.getRootImportGroup()],
      controllers: [UserController],
      providers: [...testUtils.getUserServiceGroup(), ...testUtils.getConnectionServiceGroup()],
    }).compile();

    controller = module.get<UserController>(UserController);
  });

  afterEach(async () => {
    await testUtils.closeDbConnection();
  });

  afterAll(async () => {
    module.close();
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });

  describe('findOne', () => {
    it('should return user', async () => {
      const user = (await (await testUtils.databaseService.getRepository(User)).findOne()) as User;

      const result = await controller.findOne(user.id);

      expect(result.id).toEqual(user.id);
      expect(result.displayName).toEqual(user.displayName);
    });
  });
});
