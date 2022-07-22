import { BadRequestException, HttpException, HttpStatus, Injectable, NotFoundException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { SettingEntity } from 'src/settings/entities/setting.entity';
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
    return this.todosRepository.manager.transaction('SERIALIZABLE', async (trx) => {

      const [, todayTodoCount] = await trx.findAndCount(TodoEntity, {
        where: {
          date: createTodoDto.date,
          user: {
            id: createTodoDto.user.id
          }
        }
      })

      const setting = await trx.findOne(SettingEntity, {
        where: { user: { id: createTodoDto.user.id } }
      });

      if (setting.todoPerday - todayTodoCount <= 0) {
        throw new BadRequestException(`Exceed number of todos per day`)
      }

      const todo = this.todosRepository.create(createTodoDto);

      await trx.save(todo);

      return todo;
    }).catch(e => {
      if (e?.code === 'ER_LOCK_DEADLOCK') {
        throw new HttpException('Too many request', HttpStatus.TOO_MANY_REQUESTS)
      } else {
        throw e
      }

    })
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
