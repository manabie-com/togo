import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/sequelize';
import { User } from 'src/modules/common/entities/user';

@Injectable()
export class UserService {
  constructor(
    @InjectModel(User)
    private userModel: typeof User
  ) {}

  async getUserByUserName(username: string): Promise<User> {
    const user = await this.userModel.findOne({ where: { username } });
    return user;
  }
}
