import { Module } from '@nestjs/common';

import { AppController } from './app.controller';
import { AppService } from './app.service';
import { UserModule } from './modules/users/user.module';
import { TodoTaskModule } from './modules/todo-tasks/todo-task.module';
import { RolesModule } from './modules/roles/roles.module';
import { AuthModule } from '@modules/auth/auth.module';
import { TypeOrmModule } from '@nestjs/typeorm';

@Module({
  imports: [TypeOrmModule.forRoot(), UserModule, TodoTaskModule, RolesModule, AuthModule],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
