import { Module } from '@nestjs/common';
import { DatabaseModule, TaskModule, UserModule } from '@modules/index';

@Module({
  imports: [DatabaseModule, TaskModule, UserModule],
  controllers: [],
  providers: [],
})
export class AppModule {}
