import { Module } from '@nestjs/common';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { JwtModule } from '@nestjs/jwt';
import { TypeOrmModule } from '@nestjs/typeorm';
import { UserEntity } from './entities/user.entity';
import { UserController } from './user.controller';
import { UserRepository } from './user.repository';
import { UserService } from './user.service';
import { UserSubscribe } from './user.subscribe';

@Module({
    imports: [
        TypeOrmModule.forFeature([UserEntity]),
        JwtModule.registerAsync({
            imports: [ConfigModule],
            useFactory: (config: ConfigService) => {
                return {
                    secret: config.get<string>("JWT_SECRET")
                }
            },
            inject: [ConfigService]
        }),
    ],
    controllers: [UserController],
    providers: [UserService, UserRepository, UserSubscribe],
    exports: [UserService]
})
export class UserModule { }
