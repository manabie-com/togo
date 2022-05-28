/* eslint-disable @typescript-eslint/no-unused-vars */
import { BaseService } from '@modules/common/services/base.service';
import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { EntityManager, FindManyOptions, Repository } from 'typeorm';

import { CreateTodoTaskDto } from './dto/create-todo-task.dto';
import { UpdateTodoTaskDto } from './dto/update-todo-task.dto';
import { TodoTask } from './entities/todo-task.entity';

@Injectable()
export class TodoTaskService extends BaseService {
  constructor(@InjectRepository(TodoTask) private readonly repo: Repository<TodoTask>) {
    super([]);
  }

  getRepository(manager?: EntityManager): Repository<TodoTask> {
    return manager ? manager.getRepository(TodoTask) : this.repo;
  }

  create(createTodoTaskDto: CreateTodoTaskDto) {
    return 'This action adds a new todoTask';
  }

  async findAll(filter?: FindManyOptions, manager?: EntityManager): Promise<TodoTask[]> {
    return await this.getRepository(manager).find(filter);
  }

  async findOne(id: string, manager?: EntityManager): Promise<TodoTask> {
    return await this.getRepository(manager).findOne(id);
  }

  update(id: number, updateTodoTaskDto: UpdateTodoTaskDto) {
    return `This action updates a #${id} todoTask`;
  }

  remove(id: number) {
    return `This action removes a #${id} todoTask`;
  }
}
