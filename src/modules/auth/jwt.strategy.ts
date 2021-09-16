import { Injectable } from '@nestjs/common';
import { PassportStrategy } from '@nestjs/passport';
import { ExtractJwt, Strategy } from 'passport-jwt';

import { getConfig } from '../../config';
import { ICurrentUser } from '../../decorators/user.decorator';

const appSettings = getConfig<IAppSettings>('AppSettings');

@Injectable()
export class JwtStrategy extends PassportStrategy(Strategy) {
  constructor() {
    super({
      jwtFromRequest: ExtractJwt.fromAuthHeaderAsBearerToken(),
      ignoreExpiration: false,
      secretOrKey: appSettings.jwtSecret
    });
  }

  validate(payload): ICurrentUser {
    return { id: payload.id, email: payload.email };
  }
}
