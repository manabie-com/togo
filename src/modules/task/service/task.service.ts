import { HttpStatus, Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { CreateTaskRequestDto } from '../dto/create-task-request.dto';
import { TaskEntity } from '../entity/task.entity';
import { ErrorCode, handleError } from '../../../common/util/handle-error.util';
import { UserService } from '../../../modules/user/service/user.service';

@Injectable()
export class TaskService {
  constructor(
    @InjectRepository(TaskEntity)
    private readonly taskRepository: Repository<TaskEntity>,
    private readonly userService: UserService,
  ) {}

  public async create(
    request: CreateTaskRequestDto,
    userId: string,
  ): Promise<TaskEntity> {
    const user = await this.userService.getUserById(userId);

    if (!user) {
      handleError('User not found', ErrorCode.NOT_FOUND, HttpStatus.NOT_FOUND);
    }

    const countUserTask = await this.countTaskOnDate(userId, request.startDate);

    if (countUserTask >= user.maxTask) {
      handleError(
        'Maximum tasks per day reached',
        ErrorCode.BAD_REQUEST,
        HttpStatus.BAD_REQUEST,
      );
    }

    const task = await this.taskRepository.save({
      ...request,
      userId,
    });

    return task;
  }

  public async countTaskOnDate(userId: string, date: Date): Promise<number> {
    return await this.taskRepository
      .createQueryBuilder('task')
      .andWhere('task.userId = :userId', { userId })
      .andWhere('task.startDate = :date', { date })
      .getCount();
  }
}
