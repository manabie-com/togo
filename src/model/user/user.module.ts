import { Module } from '@nestjs/common';
import { DatabaseModule } from '../../database/database.module';
import { userProviders } from './provider/user.provider';
import { UserService } from './service/limitTask.service';

@Module({
  imports: [DatabaseModule],
  providers: [UserService, ...userProviders],
})
export class UserModule {}
