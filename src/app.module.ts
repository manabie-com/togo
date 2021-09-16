import { Module } from '@nestjs/common';
import { ThrottlerModule } from '@nestjs/throttler';

import { AuthModule } from './modules/auth/auth.module';
import { JwtStrategy } from './modules/auth/jwt.strategy';
import { SharedModule } from './shared/shared.module';
import { TaskModule } from './modules/task/task.module';

@Module({
  imports: [
    AuthModule,
    TaskModule,
    SharedModule,
    ThrottlerModule.forRoot({
      ttl: 60,
      limit: 10
    })
  ],
  providers: [JwtStrategy]
})
export class AppModule {}
