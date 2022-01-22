import { Body, Controller, Post, Put} from '@nestjs/common';
import { ApiBearerAuth } from '@nestjs/swagger';
import { Request } from 'express';
import { JSONResponse } from 'src/interfaces/response.interface';
import { Constants } from '../../utils/constants';
import { handleChecker } from '../../utils/helpers';
import { LoginDTO } from './dto/login.dto';
import { RegisterDTO } from './dto/register.dto';
import { Roles } from './role.decorator';
import { UserService } from './user.service';

@Controller("/user")
export class UserController {
    constructor(private readonly userService: UserService) { }

    @Put("register")
    async postRegister(@Body() data: RegisterDTO): Promise<JSONResponse>{
        const result = await this.userService.register(data);
        return handleChecker(result);
    }

    @Post("login")
    async postLogin(@Body() data: LoginDTO): Promise<JSONResponse>{
        const result = await this.userService.login(data);
        return handleChecker(result);
    }
}
