import { Injectable, NotFoundException } from '@nestjs/common';
import { CryptoUtil } from '../../utils/crypto.util';
import { CreateUserDto } from './dto/create-user.dto';
import { UpdateUserDto } from './dto/update-user.dto';
import { UserEntity } from './entities/user.entity';
import { UserRepository } from './user.repository';

@Injectable()
export class UserService {
  constructor(
    private readonly userRepository: UserRepository,
  ){}
  
  async create(createUserDto: CreateUserDto): Promise<UserEntity> {
    const { email, password, maxTasks } = createUserDto;

		const isUserExist = await this.findOneByEmail(email);
		if(isUserExist) {
			throw new NotFoundException(`User with email ${email} already existed`);
		}

		const hashPassword = await CryptoUtil.hash(password);

		return await this.userRepository.save({ email, password: hashPassword, maxTasks });
  }

  async findOneByEmail(email: string): Promise<UserEntity | null> {
    const user = await this.userRepository.find({
      email
    });

    return Array.isArray(user) ? user[0] : null;
  }

  async findOneById(id: number) {
    return await this.userRepository.findOne(id);
  }

  async update(id: number, updateUserDto: UpdateUserDto) {
    return await this.userRepository.update(id, updateUserDto);
  }
}
