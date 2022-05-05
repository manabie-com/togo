import { Module } from '@nestjs/common';
import { SequelizeModule } from '@nestjs/sequelize';
import { User } from 'src/modules/common/entities';
import { UserService } from './user.service';

@Module({
  imports: [SequelizeModule.forFeature([User])],
  providers: [UserService],
  exports: [UserService],
})
export class UserModule {}
