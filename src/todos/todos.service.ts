import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { CreateTodoDto } from './dto/create-todo.dto';
import { UpdateTodoDto } from './dto/update-todo.dto';
import { TodoEntity } from './entities/todo.entity';

@Injectable()
export class TodosService {
  constructor(
    @InjectRepository(TodoEntity) private todosRepository: Repository<TodoEntity>,
  ) { }

  async create(createTodoDto: CreateTodoDto & { user: { id: string } }) {
    const todo = this.todosRepository.create(createTodoDto);

    await this.todosRepository.save(todo);

    return todo;
  }

  findByUserId(userId: string) {
    return this.todosRepository.find({ where: { user: { id: userId } }, relations: ['user'] })
  }

  async findOne(id: string) {
    const todo = await this.todosRepository.findOne({ where: { id }, relations: ['user'] });

    if (!todo) {
      throw new NotFoundException(`Todo not found`);
    }

    return todo;
  }

  async update(id: string, updateTodoDto: UpdateTodoDto) {
    const todo = await this.todosRepository.findOne({ where: { id } });

    if (!todo) {
      throw new NotFoundException(`Todo not found`);
    }

    await this.todosRepository.update({ id }, updateTodoDto);

    return this.todosRepository.findOne({ where: { id }, relations: ['user'] });
  }

  async remove(id: string) {
    const todo = await this.todosRepository.findOne({ where: { id }, relations: ['user'] });

    if (!todo) {
      throw new NotFoundException(`Todo not found`);
    }

    await this.todosRepository.delete({ id });

    return todo;
  }
}
