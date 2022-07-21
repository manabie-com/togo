import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { UsersModule } from 'src/users/users.module';
import { TodoEntity } from './entities/todo.entity';
import { TodosController } from './todos.controller';
import { TodosService } from './todos.service';

export const entities = [TodoEntity]

@Module({
  imports: [
    TypeOrmModule.forFeature(entities),
    UsersModule
  ],
  controllers: [TodosController],
  providers: [TodosService]
})
export class TodosModule {}
