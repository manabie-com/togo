import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { CreateToDoDto, UpdateToDoDto } from 'src/dto';
import { Task } from 'src/entities/task.entity';
import { ToDoList } from 'src/entities/toDoList.entity';
import { Repository } from 'typeorm';
import { TaskService } from './task.service';

@Injectable()
export class ToDoService {
  constructor(
    @InjectRepository(ToDoList) private toDoListRepo: Repository<ToDoList>,
  ) {}

  find(taskId: number) {
    return this.toDoListRepo.find({
      where: {
        task: {
          id: taskId,
        },
      },
    });
  }

  findOne(id: number) {
    return this.toDoListRepo.findOne(id);
  }

  doneToDoListByTaskId(taskId: number) {
    return this.toDoListRepo.update(
      {
        task: {
          id: taskId,
        },
      },
      { isDone: true },
    );
  }

  async update(id: number, { title, desc, isDone }: UpdateToDoDto) {
    const toDo = await this.findOne(id);
    if (!toDo) throw new NotFoundException('Todo not found!');

    toDo.title = title;
    toDo.desc = desc;
    if (isDone !== null || isDone !== undefined) {
      toDo.isDone = isDone;
    }
    return this.toDoListRepo.save(toDo);
  }

  async create(task: Task, { title, desc }: CreateToDoDto) {
    const todo = new ToDoList({ title, desc, task });
    return this.toDoListRepo.save(todo);
  }

  delete(id: number) {
    return this.toDoListRepo.delete(id);
  }
}
