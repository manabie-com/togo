import { Injectable, Inject } from '@nestjs/common';
import { USER } from 'src/constance/variable';
import { User } from '../schema/user.entity';

@Injectable()
export class UserService {
  constructor(
    @Inject(USER)
    private userModel: typeof User
  ) {}

  async findAll(): Promise<User[]> {
    return this.userModel.findAll<User>();
  }
}