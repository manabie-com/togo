import { Injectable, OnModuleInit } from '@nestjs/common';
import { InjectModel } from '@nestjs/sequelize';
import { hash } from 'bcrypt';
import { User, UserSettingTask } from 'src/modules/common/entities';

@Injectable()
export class InitDataService implements OnModuleInit {
  constructor(
    @InjectModel(User)
    private user: typeof User,
    @InjectModel(UserSettingTask)
    private userSettingTaskModel: typeof UserSettingTask
  ) {}

  async onModuleInit() {
    let user = await this.user.findOne({ where: { username: 'user1' } });
    if (!user) {
      const saltRounds = 10;

      const hashPassword = await hash('Test@123', saltRounds);

      user = await this.user.create({
        name: 'user1',
        username: 'user1',
        password: hashPassword,
      });

      await this.userSettingTaskModel.create({
        userId: user.id,
        maximum: 10,
      });
    }
  }
}
