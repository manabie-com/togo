import {
  Controller,
  Post,
  HttpStatus,
  UseInterceptors,
  ClassSerializerInterceptor,
  UseGuards,
  Body,
  BadRequestException,
  Delete,
  Request,
  Logger,
} from '@nestjs/common';
import { ApiBody, ApiOperation, ApiResponse, ApiTags, ApiBearerAuth } from '@nestjs/swagger';
import { Request as ExpressRequest } from 'express';
import { AuthGuard } from '@nestjs/passport';
import { Connection } from 'typeorm';

import { AuthService } from './auth.service';
import { Token, CreateTokenDto, JWTPayload } from './dto';
import { Payload } from './decorators';
import { getJWTToken } from './common/functions';

@ApiTags('Auth')
@UseInterceptors(ClassSerializerInterceptor)
@Controller('auth')
export class AuthController {
  private logger = new Logger('AuthController', true);

  constructor(private readonly connection: Connection, private readonly authService: AuthService) {}

  @Post('token')
  @ApiBody({ type: CreateTokenDto })
  @ApiOperation({ description: 'Generate a new valid JWT.' })
  @ApiResponse({
    description: 'JWT successfully created.',
    status: HttpStatus.CREATED,
    type: Token,
  })
  @ApiResponse({
    description: 'An authentication error.',
    status: HttpStatus.UNAUTHORIZED,
  })
  async token(@Body() body: CreateTokenDto) {
    let user = null;

    if (body.grant_type === 'password' && body.username && body.password) {
      user = await this.authService.validateByUsernameAndPassword(body.username, body.password);
    } else {
      throw new BadRequestException('You have to specify valid data');
    }

    return await this.authService.token(user, body?.rememberMe);
  }

  @Delete('token')
  @UseGuards(AuthGuard('jwt'))
  @ApiBearerAuth()
  @ApiOperation({ description: 'Remove user jwt' })
  @ApiResponse({
    description: 'You have already logged into the system.',
    status: HttpStatus.OK,
  })
  @ApiResponse({
    description: 'An authentication error.',
    status: HttpStatus.UNAUTHORIZED,
  })
  async removeToken(@Request() req: ExpressRequest, @Payload() payload: JWTPayload) {
    const token = getJWTToken(req);

    return await this.authService.removeToken(payload.userId, token);
  }
}
