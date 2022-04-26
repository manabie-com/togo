import { Injectable, UnauthorizedException } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { compare } from 'bcrypt';
import { UserService } from '../users/user.service';
import { LoginPayload } from './dto/login.input';

@Injectable()
export class AuthService {
  constructor(
    private jwtService: JwtService,
    private userService: UserService
  ) {}

  async validateUser(username: string, pass: string): Promise<any> {
    let user = await this.userService.getUserByUserName(username);
    if (user) {
      const compareResult = await compare(pass, user.password);
      if (compareResult) {
        console.log(user);
        const { password, ...result } = user;

        return { ...result };
      }
    }
    return null;
  }

  async login(user: LoginPayload) {
    const userData = await this.validateUser(user.username, user.password);
    if (!userData) {
      throw new UnauthorizedException();
    }

    const access_token = this.jwtService.sign(userData);
    return {
      access_token,
      ...userData,
    };
  }
}
