import { Module } from '@nestjs/common';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { JwtModule } from '@nestjs/jwt';
import { PassportModule } from '@nestjs/passport';
import { UserModule } from '../users/user.module';
import { AuthController } from './auth.contoller';
import { AuthService } from './auth.service';
import { JwtStrategy } from './strategies/jwt.strategy';

@Module({
  imports: [
    PassportModule,
    JwtModule.registerAsync({
      imports: [ConfigModule],
      useFactory: async (configService: ConfigService) => {
        console.log('sadsa', configService.get('JWT_KEY'));
        return {
          secret: configService.get('JWT_KEY'),
          signOptions: {
            expiresIn: +configService.get('EXPIRE_DATE'), //1 day
          },
        };
      },
      inject: [ConfigService],
    }),
    UserModule,
    ConfigModule,
  ],
  controllers: [AuthController],
  providers: [JwtStrategy, AuthService],
  exports: [AuthService],
})
export class AuthModule {}
