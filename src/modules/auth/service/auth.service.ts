import { HttpStatus, Injectable } from '@nestjs/common';
import * as jwt from 'jsonwebtoken';
import { UserAuthDto } from '../dto/user-auth-request.dto';
import { CommonConstants } from '../../../common/constant/common.constant';
import { EnvironmentService } from '../../../common/environment/services/environment.service';
import { EncryptionUtil } from '../../../common/util/encryption.util';
import { ErrorCode, handleError } from '../../../common/util/handle-error.util';
import { UserService } from '../../../modules/user/service/user.service';

@Injectable()
export class AuthService {
  constructor(
    private readonly environmentService: EnvironmentService,
    private readonly userService: UserService,
  ) {}

  public async login(
    request: UserAuthDto,
  ): Promise<{ id: string; token: string }> {
    const user = await this.userService.getUserByUsername(request.username);

    if (!user) {
      handleError('User not found', ErrorCode.NOT_FOUND, HttpStatus.NOT_FOUND);
    }

    if (!EncryptionUtil.comparePassword(request.password, user.password)) {
      handleError(
        'Username or password mismatch',
        ErrorCode.BAD_REQUEST,
        HttpStatus.BAD_REQUEST,
      );
    }

    const userData = {
      id: user.id,
      username: user.username,
    };

    const token = jwt.sign(
      userData,
      this.environmentService.getKey(CommonConstants.JWT_SECRET_KEY),
      { expiresIn: CommonConstants.TOKEN_EXPIRE_TIME },
    );

    return {
      id: user.id,
      token,
    };
  }

  public async register(request: UserAuthDto) {
    const user = await this.userService.getUserByUsername(request.username);

    if (user) {
      handleError(
        'User already existed',
        ErrorCode.BAD_REQUEST,
        HttpStatus.BAD_REQUEST,
      );
    }

    await this.userService.createUser(request);
  }
}
