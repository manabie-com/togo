import { Injectable } from '@nestjs/common';
import { v1 as uuid } from 'uuid'

import { CreateUserDto } from './dto/create-user-dto';
import { User } from './schemas/user.schema';

@Injectable()
export class UserRepository {

  private users: User[] = [];

  async create(createUserDto: CreateUserDto): Promise<User> {
    const { name,
        dailyTaskLimit,
        } = createUserDto;
    const user = {
      id: uuid(),
      name,
      dailyTaskLimit,
      dailyTaskCounter: 0,
      dailyTaskDate: new Date(),
      createdAt: new Date(),
      updatedAt: new Date(),
    }
    this.users.push(user)
    return user;
  }

  async findAll(): Promise<User[]> {
    return this.users;
  }

}
