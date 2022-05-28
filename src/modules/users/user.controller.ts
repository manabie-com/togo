import {
  Controller,
  Get,
  Post,
  Body,
  Patch,
  Param,
  Delete,
  ClassSerializerInterceptor,
  UseInterceptors,
  UseGuards,
} from '@nestjs/common';
import { ApiBearerAuth, ApiTags } from '@nestjs/swagger';

import { UserService } from './user.service';
import { CreateUserDto } from './dto/create-user.dto';
import { UpdateUserDto } from './dto/update-user.dto';
import { PermissionAction } from '@modules/permissions/permission.action.enum';
import { PermissionPossession } from '@modules/permissions/permission.possession.enum';
import { PermissionResource } from '@modules/permissions/permission.resource.enum';
import { UseRoles } from '@modules/common/decorators/role.decorator';
import { AuthGuard } from '@nestjs/passport';
import { RolesGuard } from '@modules/common/guards/role.guards';

@ApiTags('Users')
@UseInterceptors(ClassSerializerInterceptor)
@UseGuards(AuthGuard('jwt'), RolesGuard)
@ApiBearerAuth()
@Controller('users')
export class UserController {
  constructor(private readonly usersService: UserService) {}

  @Post()
  create(@Body() createUserDto: CreateUserDto) {
    return this.usersService.create(createUserDto);
  }

  @Get()
  findAll() {
    return this.usersService.findAll();
  }

  @Get(':id')
  @UseRoles({
    resource: PermissionResource.Users,
    action: PermissionAction.Detail,
    possession: PermissionPossession.Own,
  })
  async findOne(@Param('id') id: string) {
    return await this.usersService.findById(id);
  }

  @Patch(':id')
  update(@Param('id') id: string, @Body() updateUserDto: UpdateUserDto) {
    return this.usersService.update(+id, updateUserDto);
  }

  @Delete(':id')
  remove(@Param('id') id: string) {
    return this.usersService.remove(+id);
  }
}
