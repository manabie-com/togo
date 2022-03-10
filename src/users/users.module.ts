import { Module } from '@nestjs/common';
import { UserService } from './users.service.js';
import { UserController } from './users.controller.js';
import { TypeOrmModule } from '@nestjs/typeorm';
import { User } from './users.entity.js';

@Module({
  controllers: [UserController],
  imports: [TypeOrmModule.forFeature([User])],
  providers: [UserService],
  exports: [UserService, TypeOrmModule],
})
export class UserModule {}
