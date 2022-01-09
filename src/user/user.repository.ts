import { HttpException, Injectable, NotFoundException } from '@nestjs/common';
import { v1 as uuid } from 'uuid'

import { CreateUserDto } from './dto/create-user-dto';
import { User } from './schemas/user.schema';
import { mockUsersList } from './user';

@Injectable()
export class UserRepository {
  private users = mockUsersList();

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

  async update(userId: string, user: User): Promise<User[]> {
    const objIndex = this.users.findIndex((obj => obj.id == userId));
    this.users[objIndex] = user;
    return this.users;
  }

  async incrementTask(userId: string): Promise<Boolean> {
    const foundUser = this.users.find((user) => user.id === userId);
    if (foundUser) {
      const now = new Date();
      if (foundUser.dailyTaskDate.toISOString().split('T')[0] === now.toISOString().split('T')[0]) {
        if(foundUser.dailyTaskCounter < foundUser.dailyTaskLimit) {
          foundUser.dailyTaskCounter++;
          this.update(userId, foundUser);
          return true;
        } 
        throw new HttpException("Task daily limit exceeded", 405);
      } else {
        foundUser.dailyTaskDate = now;
        foundUser.dailyTaskCounter = 1;
        this.update(userId, foundUser);
        return true;
      }
    }
    throw new NotFoundException();
  }
}
