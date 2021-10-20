import { Injectable } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { InjectRepository } from '@nestjs/typeorm';
// import { Brackets, Connection, getManager, IsNull, Repository } from 'typeorm';
import { User } from './models/user.entity';

@Injectable()
export class AppService {
  constructor(
    private jwtService: JwtService, // @InjectRepository(User) // private userRepository: Repository<User>,
  ) {}
  getHello(): string {
    return 'Hello World!';
  }

  async login(userId: string, password: string) {
    const user = await User.findOne({
      where: {
        userId,
        password,
      },
    });
    const payload = this.jwtService.sign({
      id: user.id,
      user_name: user.userId,
    });
    return {
      data: payload,
    };
  }
}
