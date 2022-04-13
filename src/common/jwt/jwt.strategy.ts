import { ExtractJwt, Strategy } from 'passport-jwt';
import { Injectable } from '@nestjs/common';
import { PassportStrategy } from '@nestjs/passport';
import { CommonConstants } from '../constant/common.constant';
import { UserAccount } from '../guards/user-account.class';
import { EnvironmentService } from '../environment/services/environment.service';

@Injectable()
export class JwtStrategy extends PassportStrategy(Strategy) {
  constructor(private readonly environmentService: EnvironmentService) {
    super({
      jwtFromRequest: ExtractJwt.fromAuthHeaderAsBearerToken(),
      ignoreExpiration: false,
      secretOrKey: environmentService.getKey(CommonConstants.JWT_SECRET_KEY),
    });
  }

  validate(payload: any): UserAccount {
    return { id: payload.id, username: payload.username };
  }
}
