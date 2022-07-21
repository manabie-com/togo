import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { SettingEntity } from 'src/settings/entities/setting.entity';
import { Repository } from 'typeorm';
import { CreateUserDto } from './dto/create-user.dto';
import { UpdateUserDto } from './dto/update-user.dto';
import { UserEntity } from './entities/user.entity';

@Injectable()
export class UsersService {
  constructor(
    @InjectRepository(UserEntity) private usersRepository: Repository<UserEntity>,
  ) { }

  async create(createUserDto: CreateUserDto) {
    const user = this.usersRepository.create({ ...createUserDto, setting: {} });
    user.setting.user = user;

    await this.usersRepository.save(user);

    return this.usersRepository.findOne({ where: { id: user.id } })
  }

  findAll() {
    return this.usersRepository.find()
  }

  async findOne(id: string) {
    const user = await this.usersRepository.findOne({ where: { id } });

    if (!user) {
      throw new NotFoundException(`user not found`);
    }

    return user;
  }

  async update(id: string, updateUserDto: UpdateUserDto) {
    const user = await this.usersRepository.findOne({ where: { id } });

    if (!user) {
      throw new NotFoundException(`user not found`);
    }

    await this.usersRepository.update({ id }, updateUserDto);

    return this.usersRepository.findOne({ where: { id } });
  }

  async remove(id: string) {
    const user = await this.usersRepository.findOne({ where: { id } });

    if (!user) {
      throw new NotFoundException(`user not found`);
    }

    await this.usersRepository.delete({ id });

    return user;
  }
}
