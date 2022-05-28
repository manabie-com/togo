import * as utils from '@modules/common/utils';
import { Injectable, BadRequestException, Logger } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { environment } from '@env/environment';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';

import { User } from '@modules/users/entities/user.entity';
import { UserService } from '@modules/users/user.service';
import { Token } from './dto';
import { BlacklistedToken } from './blacklisted-token.entity';
import { PermissionStatus } from '@modules/permissions/permission.status.enum';
import { LoginFailedException } from '@modules/common/exceptions';
import { MessageResult } from '@modules/common/dto/message-result.model';

@Injectable()
export class AuthService {
  private logger = new Logger('AuthService', true);

  constructor(
    @InjectRepository(BlacklistedToken) private readonly bTokenRepo: Repository<BlacklistedToken>,
    private readonly userService: UserService,
    private readonly jwtService: JwtService,
  ) {}

  async validateByUsernameAndPassword(username: string, password: string): Promise<User> {
    const user = await this.validateUser(username, password);

    if (!user) {
      throw new LoginFailedException();
    }

    return user;
  }

  async token(user: Pick<User, 'username' | 'role' | 'id'>, rememberMe?: boolean): Promise<Token> {
    try {
      const expiresIn = !rememberMe ? environment.jwt.expiration : environment.jwt.rememberMe;

      const permissions = user.role?.permissions
        ?.filter((p) => p.status === PermissionStatus.Active)
        .map((p) => {
          return `${p.resource}_${p.action}`;
        });

      const payload = {
        username: user.username,
        role: user.role?.name,
        permissions,
        userId: user.id,
      };

      const token = this.jwtService.sign(payload, {
        audience: environment.jwt.audience,
        subject: user.id,
        issuer: environment.jwt.issuer,
        jwtid: utils.generateID(8),
        secret: environment.jwt.secretKey,
        expiresIn,
      });
      const refreshToken = utils.generateID(64);
      const accessToken = new Token(token, refreshToken, expiresIn);

      return accessToken;
    } catch (error) {
      this.logger.error('token got error: ', error);
      throw error;
    }
  }

  async removeToken(userId: string, token: string): Promise<MessageResult> {
    const bTokenRepo: Repository<BlacklistedToken> = this.bTokenRepo;

    if (!userId) {
      throw new BadRequestException('User Id is invalid');
    }

    if (!token) {
      throw new BadRequestException('Token is invalid');
    }

    const entity = new BlacklistedToken({
      userId,
      token,
    });

    await bTokenRepo.save(entity);

    return { message: 'You have already logged out the system.' };
  }

  private async validateUser(username: string, password: string): Promise<User | undefined> {
    const signInType = utils.checkSignInType(username);

    const user =
      signInType === 'email'
        ? await this.userService.findByEmail(username)
        : await this.userService.findByUsername(username);
    const passwordHash = utils.createPasswordHash(password);

    if (user && passwordHash == user.password) {
      return user;
    }

    return undefined;
  }
}
