import { Injectable, BadRequestException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { User } from './user.entity';
import { UserQuery, UserFilter, UpsertUserDTO } from './context/index';
import { StatusEnum, OrderByEnum } from '@common/index';
import { Config } from '@common/config';

@Injectable()
export class UserService {
  constructor(
    @InjectRepository(User)
    private userRepository: Repository<User>,
  ) {}

  async findAll(query: UserQuery): Promise<User[]> {
    return this.userRepository.find({
      where: { ...query.filter },
      order: {
        ...query.orderBy,
      },
      skip: query?.pagination?.offset,
      take: query?.pagination?.limit,
    });
  }

  async findOne(filter: UserFilter): Promise<User> {
    return this.userRepository.findOne({ where: { ...filter } });
  }

  async create(context: UpsertUserDTO): Promise<User> {
    const user = Object.assign(new User(), context);
    return this.userRepository.save(user);
  }

  async update(id: number, context: UpsertUserDTO): Promise<User> {
    //check user exists
    const user = await this.userRepository.findOne({
      where: { id },
    });
    if (!user) {
      throw new BadRequestException({
        message: 'User not found',
      });
    }
    Object.assign(user, context);
    return this.userRepository.save(user);
  }

  async delete(id: number): Promise<number> {
    const result = await this.userRepository.delete({ id });
    return result.affected || 0;
  }
}
