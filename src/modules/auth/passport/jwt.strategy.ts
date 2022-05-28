import { Strategy, VerifiedCallback } from 'passport-custom';
import { Request } from 'express';
import { PassportStrategy } from '@nestjs/passport';
import { Injectable } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { TokenExpiredError, JsonWebTokenError } from 'jsonwebtoken';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';

import { JWTPayload } from '../dto';
import { getJWTToken } from '../common/functions';
import { environment } from '@env/environment';
import { BlacklistedToken } from '../blacklisted-token.entity';
import {
  AuthenticationFailedException,
  TokenExpiredException,
  InvalidTokenException,
} from '@modules/common/exceptions';

@Injectable()
export class JwtStrategy extends PassportStrategy(Strategy, 'jwt') {
  constructor(
    private readonly jwtService: JwtService,
    @InjectRepository(BlacklistedToken) private readonly bTokenRepo: Repository<BlacklistedToken>,
  ) {
    super(async (req: Request, done: VerifiedCallback) => {
      const token = getJWTToken(req);

      if (!token) {
        return done(new AuthenticationFailedException(), null);
      }

      try {
        done(null, await this.validate(token));
      } catch (err) {
        if (err instanceof TokenExpiredError) {
          done(new TokenExpiredException(), null);
        } else if (err instanceof JsonWebTokenError) {
          done(new InvalidTokenException(), null);
        } else {
          done(new InvalidTokenException(), null);
        }
      }
    });
  }

  async validate(token: string): Promise<JWTPayload> {
    const bTokenRepo: Repository<BlacklistedToken> = this.bTokenRepo;
    const hasBlacklisted = await bTokenRepo.count({ where: { token } });

    if (hasBlacklisted > 0) {
      throw new InvalidTokenException();
    }

    const payload = await this.jwtService.verify(token, { secret: environment.jwt.secretKey });

    return new JWTPayload(payload);
  }
}
