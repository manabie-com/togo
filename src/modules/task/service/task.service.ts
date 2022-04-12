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

    const task = await this.taskRepository.save({
      ...request,
      userId,
    });

    return task;
  }
}
