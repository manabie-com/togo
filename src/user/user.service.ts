import { Injectable } from '@nestjs/common';
import { CreateUserDto } from './dto/create-user-dto';
import { User } from './schemas/user.schema';
import { UserRepository } from './user.repository';

@Injectable()
export class UserService {

  constructor(private userRepository: UserRepository) { }

  async create(createUserDto: CreateUserDto): Promise<User> {
    return this.userRepository.create(createUserDto);
  }

  async incrementTask(userId: string): Promise<Boolean> {
    return this.userRepository.incrementTask(userId);
  }

  async findAll(): Promise<User[]> {
    return await this.userRepository.findAll();
  }
}
