import { Module } from '@nestjs/common';
import { AuthService } from './auth.service.js';
import { UserModule } from '../users/users.module.js';
import { PassportModule } from '@nestjs/passport';
import { LocalStrategy } from './local.strategy.js';
import { JwtStrategy } from './jwt.strategy.js';
import { JwtModule } from '@nestjs/jwt';
import { jwtConstants } from './constants.js';

@Module({
  imports: [
    UserModule,
    PassportModule,
    JwtModule.register({
      secret: jwtConstants.secret,
      signOptions: { expiresIn: '14400s' },
    }),
  ],
  providers: [AuthService, LocalStrategy, JwtStrategy],
  exports: [AuthService],
})
export class AuthModule {}
