import { Module } from '@nestjs/common';
import { AppController } from './app.controller.js';
import { TodoModule } from './todo/todo.module.js';
import { TypeOrmModule } from '@nestjs/typeorm';
import { getConnectionOptions } from 'typeorm';
import { AuthModule } from './auth/auth.module.js';
import { UserModule } from './users/users.module.js';

TypeOrmModule.forRootAsync({
  useFactory: async () =>
    Object.assign(await getConnectionOptions(), {
      autoLoadEntities: true,
    }),
});

@Module({
  imports: [
    TypeOrmModule.forRoot(),
    TodoModule,
    AuthModule,
    UserModule
  ],
  controllers: [AppController],
})
export class AppModule {}
