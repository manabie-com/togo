import { Module, HttpModule } from '@nestjs/common';
import { PassportModule } from '@nestjs/passport';
import { JwtModule } from '@nestjs/jwt';
import { TypeOrmModule } from '@nestjs/typeorm';

import { AuthController } from './auth.controller';
import { AuthService } from './auth.service';
import { JwtStrategy } from './passport/jwt.strategy';
import { UserModule } from '@modules/users/user.module';
import { environment } from '@env/environment';
import { BlacklistedToken } from './blacklisted-token.entity';

@Module({
  imports: [
    TypeOrmModule.forFeature([BlacklistedToken]),
    HttpModule,
    UserModule,
    PassportModule,
    JwtModule.register({
      secret: environment.jwt.secretKey,
      signOptions: { expiresIn: environment.jwt.expiration },
    }),
  ],
  controllers: [AuthController],
  providers: [AuthService, JwtStrategy],
})
export class AuthModule {}
