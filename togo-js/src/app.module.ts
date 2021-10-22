import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { TypeOrmModule } from '@nestjs/typeorm';
import { getConnectionOptions } from 'typeorm';
import { JwtModule } from '@nestjs/jwt';
import { Task } from './models/task.entity';
import { User } from './models/user.entity';
import { ConfigModule } from '@nestjs/config';
import { JwtStrategy } from './services/jwt/strategy';
import { PassportModule } from '@nestjs/passport';

@Module({
  imports: [
    PassportModule,
    ConfigModule.forRoot(),
    TypeOrmModule.forFeature([User, Task]),
    TypeOrmModule.forRootAsync({
      useFactory: async () => {
        return Object.assign(await getConnectionOptions(), {
          autoLoadEntities: true,
          type: 'postgres',
          host: 'db',
          port: 5432,
          username: 'postgres',
          password: 'postgres',
          database: 'togo',
          entities: ['./dist/**/**.entity{.ts,.js}'],
          subscribers: ['./dist/**/**.subscriber{.ts,.js}'],
          synchronize: false,
          logging: true,
          migrationsRun: true,
          migrations: ['./dist/migrations/*.js'],
          cli: {
            migrationsDir: 'src/migrations',
          },
        });
      },
    }),
    JwtModule.register({
      secret: process.env.JWT_SECRET,
      signOptions: { expiresIn: '1h' },
    }),
  ],
  controllers: [AppController],
  providers: [AppService, JwtStrategy],
  exports: [JwtStrategy],
})
export class AppModule {}
