import { Module } from '@nestjs/common';
import { HttpModule } from '@nestjs/axios';
import { PgModule } from './common/database/pg.module';
import { EnvironmentModule } from './common/environment/environment.module';
import { JwtStrategy } from './common/jwt/jwt.strategy';
import { AuthModule } from './modules/auth/auth.module';
import { TaskEntity } from './modules/task/entity/task.entity';
import { TaskModule } from './modules/task/task.module';
import { UserEntity } from './modules/user/entity/user.entity';
import { UserModule } from './modules/user/user.module';

@Module({
  imports: [
    HttpModule,
    EnvironmentModule,
    TaskModule,
    UserModule,
    AuthModule,
    PgModule.register({
      entities: [UserEntity, TaskEntity],
    }),
  ],
  controllers: [],
  providers: [JwtStrategy],
})
export class AppModule {}
