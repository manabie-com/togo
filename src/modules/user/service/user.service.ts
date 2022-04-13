import { Repository } from 'typeorm';
import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { EncryptionUtil } from '../../../common/util/encryption.util';
import { UserAuthDto } from '../../../modules/auth/dto/user-auth-request.dto';
import { UserEntity } from '../entity/user.entity';

@Injectable()
export class UserService {
  constructor(
    @InjectRepository(UserEntity)
    private readonly userRepository: Repository<UserEntity>,
  ) {}

  public async getUserById(id: string): Promise<UserEntity> {
    return await this.userRepository.findOne(id);
  }

  public async getUserByUsername(username: string): Promise<UserEntity> {
    return await this.userRepository.findOne({ username });
  }

  public async createUser(request: UserAuthDto): Promise<UserEntity> {
    const encryptedPassword = await EncryptionUtil.encryptPassword(
      request.password,
    );

    return await this.userRepository.save({
      username: request.username,
      password: encryptedPassword,
    });
  }
}
