import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { TodoModule } from './todo/todo.module';
import { TypeOrmModule } from '@nestjs/typeorm';
import { config } from './config';
import { AuthModule } from './auth/auth.module';
import { UserModule } from './users/users.module';

TypeOrmModule.forRootAsync({
  useFactory: () => ({
    type: 'mysql',
    host: config.dbHost,
    port: config.dbPort,
    username: config.dbUsername,
    password: config.dbPassword,
    database: config.dbDatabase,
    entities: [__dirname + '/**/*.entity{.ts,.js}'],
    synchronize: true,
  }),
});

@Module({
  imports: [TypeOrmModule.forRoot(), TodoModule, AuthModule, UserModule],
  controllers: [AppController],
})
export class AppModule {}
