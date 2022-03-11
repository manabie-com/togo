import { startOfToday, endOfToday } from 'date-fns';
import { Injectable, Options } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Connection, Repository, Between, In } from 'typeorm';
import { Todo, TodoStatus } from './todo.entity';

@Injectable()
export class TodoService {
  constructor(
    private connection: Connection,
    @InjectRepository(Todo)
    private todosRepository: Repository<Todo>,
  ) {}

  async getTodos(username: string, status?: TodoStatus | undefined): Promise<Todo[]> {
    let options: { user: { username: string }, status?: TodoStatus | undefined } = {
      user: { username: username }
    };
    if (status)
      options.status = status;

    return await this.todosRepository.find({
      where: options,
      order: {
        createAt: 'DESC',
      },
    });
  }

  async countTodoDaily(username: string): Promise<number> {
    return await this.todosRepository.count({
      where: {
        user: { username: username },
        createAt: Between(startOfToday(), endOfToday()),
      }
    });
  }

  async setTodoStatus(id: string, status: TodoStatus) {
    await this.todosRepository.update({ id }, { status });
  }

  async setManyTodoStatus(ids: string[], status: TodoStatus) {
    await this.todosRepository.update(
      { id: In(ids) },
      { status },
    );
  }

  async deleteTodoById(id: string) {
    await this.todosRepository.delete(id);
  }

  async deleteAllTodos(username: string) {
    await this.todosRepository.delete({
      user: {
        username: username,
      }
    });
  }
}
