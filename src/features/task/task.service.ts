import { Injectable, InternalServerErrorException, NotFoundException } from '@nestjs/common';
import { UserService } from '../user/user.service';
import { CreateTaskDto } from './dto/create-task.dto';
import { TaskEntity } from './entities/task.entity';
import { TaskRepository } from './task.repository';

@Injectable()
export class TaskService {
  constructor(
    private readonly userService: UserService,
    private readonly taskRepository: TaskRepository
  ){}

  async create(createTaskDto: CreateTaskDto): Promise<TaskEntity> {
    const { userId } = createTaskDto;

    const user = await this.userService.findOneById(userId);
    if(!user) {
      throw new NotFoundException(`User with id ${userId} does not exist`);
    }

    const currentNumberOfTasks = await this.taskRepository.getNumberOfTaskByUserId(userId) || 0;
    if(currentNumberOfTasks >= user.maxTasks) {
      throw new InternalServerErrorException(`The number of tasks of user has reached limit.`);
    }

    return await this.taskRepository.save(createTaskDto);
  }
}
