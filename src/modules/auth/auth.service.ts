import { Injectable } from '@nestjs/common';
import jwt from 'jsonwebtoken';

import { getHash, getSalt, unixTime } from '../../common/utils';
import { getConfig } from '../../config';
import { UserRepository } from '../../shared/repositories/user.repository';
import { UserLoginDto, UserRegisterDto } from './auth.dto';

@Injectable()
export class AuthService {
  constructor(private _userRepository: UserRepository) {}
  async createNewUser(userDto: UserRegisterDto): Promise<boolean> {
    const existedUser = await this._userRepository.findOne({
      email: userDto.email
    });

    if (existedUser) {
      return false;
    }

    const salt = getSalt();
    const hash = getHash(userDto.password, salt);

    await this._userRepository.save({
      ...userDto,
      salt,
      hash,
      createdAt: unixTime()
    });

    return true;
  }

  async login({ email, password }: UserLoginDto): Promise<unknown> {
    const user = await this._userRepository.findOne({ email });
    if (!user) {
      return false;
    }

    const salt = user.salt;
    const hash = getHash(password, salt);
    if (user.hash !== hash) {
      return false;
    }

    const appSettings = getConfig<IAppSettings>('AppSettings');
    const token = jwt.sign(
      {
        id: user.id,
        email: user.email,
        exp: unixTime() + 3600 * 7
      },
      appSettings.jwtSecret
    );

    return {
      id: user.id,
      name: user.name,
      email: user.email,
      createdAt: user.createdAt,
      token
    };
  }

  getProfile(id: number): Promise<unknown> {
    return this._userRepository.findOne({
      where: { id },
      select: ['id', 'name', 'email', 'createdAt']
    });
  }
}
