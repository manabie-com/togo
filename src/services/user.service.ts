import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { CreateUserDto, GetUserListDto, UpdateUserDto } from 'src/dto';
import { User } from 'src/entities/user.entity';
import { Repository } from 'typeorm';

@Injectable()
export class UserService {
  constructor(@InjectRepository(User) private userRepo: Repository<User>) {}

  findOne(id: number) {
    return this.userRepo.findOne(id);
  }

  async find({ limit, page }: GetUserListDto) {
    const [items, total] = await this.userRepo.findAndCount({
      skip: page * limit,
      take: limit,
    });
    return {
      items,
      total,
    };
  }

  async updateUser(id: number, { name, dailyMaxTasks }: UpdateUserDto) {
    const currUser = await this.findOne(id);
    if (!currUser) throw new NotFoundException('Không tìm thấy tài khoản');

    currUser.name = name;
    currUser.dailyMaxTasks = dailyMaxTasks;
    return this.userRepo.save(currUser);
  }

  createUser({ name, dailyMaxTasks }: CreateUserDto) {
    const user = new User({ name, dailyMaxTasks });
    return this.userRepo.save(user);
  }
}
