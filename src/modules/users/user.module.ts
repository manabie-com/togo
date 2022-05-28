import { Module } from '@nestjs/common';
import { UserService } from './user.service';
import { environment } from '@env/environment';
import { JwtModule } from '@nestjs/jwt';
import { TypeOrmModule } from '@nestjs/typeorm';
import { User } from './entities/user.entity';

@Module({
  imports: [
    TypeOrmModule.forFeature([User]),
    JwtModule.register({
      secret: environment.jwt.secretKey,
      signOptions: { expiresIn: environment.jwt.expiration },
    }),
  ],
  providers: [UserService],
  // controllers: [UserController],
  exports: [UserService, TypeOrmModule],
})
export class UserModule {}
