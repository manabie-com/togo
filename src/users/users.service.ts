import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Connection, Repository } from 'typeorm';
import { User } from './users.entity.js';
@Injectable()
export class UserService {
  constructor(
    private connection: Connection,
    @InjectRepository(User)
    private usersRepository: Repository<User>,
  ) {}

  async createUser(user: User) {
    if (await this.usersRepository.findOne(user.username))
      throw new HttpException({
        status: HttpStatus.BAD_REQUEST,
        error: 'Username already existed',
      }, HttpStatus.BAD_REQUEST);
    else
      await this.usersRepository.save(user);
  }

  async findOne(username: string): Promise<User> {
    return this.usersRepository.findOne({
      where: { username: username },
      relations: ['todos'],
    });
  }
}
