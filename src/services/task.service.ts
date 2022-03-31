import {
  BadRequestException,
  Injectable,
  NotFoundException,
} from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { CreateTaskDto, GetTasksDto, UpdateTaskDto } from 'src/dto';
import { ETaskStatus, Task } from 'src/entities/task.entity';
import { In, Repository } from 'typeorm';
import { ToDoService } from './todo.service';
import { UserService } from './user.service';

@Injectable()
export class TaskService {
  constructor(
    @InjectRepository(Task) private taskRepo: Repository<Task>,
    private userService: UserService,
    private toDoService: ToDoService,
  ) {}

  countTasks(userId: number) {
    return this.taskRepo.count({
      where: {
        user: {
          id: userId,
        },
        status: In([ETaskStatus.IN_PROGRESS, ETaskStatus.DO_TO]),
      },
    });
  }

  async find({ limit, page }: GetTasksDto) {
    const [items, total] = await this.taskRepo.findAndCount({
      skip: page * limit,
      take: limit,
    });
    return {
      items,
      total,
    };
  }

  findOne(id: number) {
    return this.taskRepo.findOne(id);
  }

  async update(
    id: number,
    { title, desc, userId, deadlineAt, status }: UpdateTaskDto,
  ) {
    const task = await this.findOne(id);
    if (!task) throw new NotFoundException('Không tìm thấy task');

    if (userId) {
      const user = await this.userService.findOne(userId);
      if (user) throw new NotFoundException('Không tìm thấy user');

      const todayTaskCount = await this.countTasks(userId);
      if (todayTaskCount + 1 >= user.dailyMaxTasks) {
        throw new BadRequestException('User đã đạt giới hạng tasks hôm nay');
      }
      task.user = user;
    }
    if (status !== undefined && status !== null) {
      task.status = status;
      if (status == ETaskStatus.COMPLETE) {
        await this.toDoService.doneToDoListByTaskId(id);
      }
    }
    task.title = title;
    task.desc = desc;
    task.deadlineAt = new Date(deadlineAt);
    return this.taskRepo.save(task);
  }

  async create({ title, desc, deadlineAt }: CreateTaskDto) {
    const task = new Task({ title, deadlineAt: new Date(deadlineAt), desc });
    return this.taskRepo.save(task);
  }

  delete(id: number) {
    return this.taskRepo.delete(id);
  }
}
