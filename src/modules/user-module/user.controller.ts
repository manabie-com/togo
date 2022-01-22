import { Body, Controller, Get, Post, Query, Req } from '@nestjs/common';
import { ApiBearerAuth } from '@nestjs/swagger';
import { Request } from 'express';
import { JSONResponse } from 'src/interfaces/response.interface';
import { Constants } from '../../utils/constants';
import { handleChecker } from '../../utils/helpers';
import { ConfigureMaxTaskDTO } from './dto/configure_max_task.dto';
import { FindUserDTO } from './dto/find.dto';
import { LoginDTO } from './dto/login.dto';
import { RegisterDTO } from './dto/register.dto';
import { TaskDTO } from './dto/task.dto';
import { Roles } from './role.decorator';
import { UserService } from './user.service';

@Controller("/user")
export class UserController {
    constructor(private readonly userService: UserService) { }

    @Post("register")
    async postRegister(@Body() data: RegisterDTO): Promise<JSONResponse> {
        const result = await this.userService.register(data);
        return handleChecker(result);
    }

    @Post("login")
    async postLogin(@Body() data: LoginDTO): Promise<JSONResponse> {
        const result = await this.userService.login(data);
        return handleChecker(result);
    }

    @Get("find")
    @ApiBearerAuth("access-token")
    @Roles(Constants.ADMIN_ROLE)
    async findUser(@Query() query: FindUserDTO): Promise<JSONResponse> {
        const result = await this.userService.find(query);
        return handleChecker(result);
    }

    @Post("add-task")
    @ApiBearerAuth("access-token")
    async addTask(@Req() req: Request, @Body() data: TaskDTO): Promise<JSONResponse> {
        const result = await this.userService.addTask(req.user.id, data);
        return handleChecker(result);
    }

    @Post("configure-max-task")
    @ApiBearerAuth("access-token")
    @Roles(Constants.ADMIN_ROLE)
    async configureMaxTask(@Body() data: ConfigureMaxTaskDTO): Promise<JSONResponse> {
        const result = await this.userService.configureMaxTask(data);
        return handleChecker(result);
    }
}
