import {
  Controller,
  Get,
  Post,
  Body,
  ClassSerializerInterceptor,
  UseGuards,
  UseInterceptors,
  HttpCode,
  HttpStatus,
  BadRequestException,
} from '@nestjs/common';
import { ApiBearerAuth, ApiTags } from '@nestjs/swagger';

import { TodoTaskService } from './todo-task.service';
import { RolesGuard } from '@modules/common/guards/role.guards';
import { AuthGuard } from '@nestjs/passport';
import { UseRoles } from '@modules/common/decorators/role.decorator';
import { PermissionAction } from '@modules/permissions/permission.action.enum';
import { PermissionResource } from '@modules/permissions/permission.resource.enum';
import { Payload } from '@modules/auth/decorators';
import { JWTPayload } from '@modules/auth/dto';
import { Connection, FindManyOptions, getConnection } from 'typeorm';
import { UserService } from '@modules/users/user.service';
import { AssignTaskDto as PickTaskDto } from './dto/assign-task.dto';
import { MessageResult } from '@modules/common/dto/message-result.model';
import {
  ReachedMaximumTaskTodayBadRequestException,
  TaskIsAssignedToYouBadRequestException,
  TaskIsHandedByOtherBadRequestException,
  TaskIsNotFoundException,
} from '@modules/common/exceptions';

@ApiTags('Tasks')
@Controller('tasks')
@UseInterceptors(ClassSerializerInterceptor)
@UseGuards(AuthGuard('jwt'), RolesGuard)
@ApiBearerAuth()
export class TodoTaskController {
  constructor(private readonly todoTaskService: TodoTaskService, private readonly userService: UserService) {}

  @Get()
  @UseRoles({
    resource: PermissionResource.Tasks,
    action: PermissionAction.List,
  })
  async findAll() {
    return this.todoTaskService.findAll();
  }

  @Post('/pick')
  @HttpCode(HttpStatus.OK)
  @UseRoles({
    resource: PermissionResource.Tasks,
    action: PermissionAction.Pick,
  })
  async pick(@Body() body: PickTaskDto, @Payload() payload: JWTPayload): Promise<MessageResult> {
    try {
      return getConnection().transaction(async (manager) => {
        const user = await this.userService.findById(payload.userId, manager);

        if (!(await user.canPickMoreTask())) {
          throw new ReachedMaximumTaskTodayBadRequestException();
        }

        const task = await this.todoTaskService.findOne(body.taskId, manager);

        if (!task) {
          throw new TaskIsNotFoundException();
        }

        if (task && task.assigneeId === user.id) throw new TaskIsAssignedToYouBadRequestException();

        if (task && task.assigneeId !== null) throw new TaskIsHandedByOtherBadRequestException();

        task.assigneeId = user.id;

        await task.save();

        return { message: 'Pick task successfully' };
      });
    } catch (error) {
      throw error;
    }
  }

  @Get('/my-tasks')
  @UseRoles({
    resource: PermissionResource.Tasks,
    action: PermissionAction.List,
  })
  async findMyTasks(@Payload() payload: JWTPayload) {
    return getConnection().transaction(async (manager) => {
      const user = await this.userService.findById(payload.userId, manager);

      const filter: FindManyOptions = {
        where: { assigneeId: user.id },
      };

      return this.todoTaskService.findAll(filter);
    });
  }
}
