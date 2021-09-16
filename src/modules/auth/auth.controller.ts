import { Body, Controller, Get, Post, UseGuards } from '@nestjs/common';
import { ApiOperation, ApiTags } from '@nestjs/swagger';

import { UserLoginDto, UserRegisterDto } from './auth.dto';
import { AuthService } from './auth.service';
import { CurrentUser, ICurrentUser } from '../../decorators/user.decorator';
import { AuthGuard } from '../../guards/auth.guard';

@Controller('auth')
@ApiTags('auth')
export class AuthController {
  constructor(private _authService: AuthService) {}

  @Post('register')
  @ApiOperation({ summary: 'User register' })
  register(@Body() userDto: UserRegisterDto): Promise<boolean> {
    return this._authService.createNewUser(userDto);
  }

  @Post('login')
  @ApiOperation({ summary: 'User login' })
  login(@Body() userDto: UserLoginDto): Promise<unknown> {
    return this._authService.login(userDto);
  }

  @Get('profile')
  @ApiOperation({ summary: 'User login' })
  @UseGuards(new AuthGuard('jwt'))
  getProfile(@CurrentUser() user: ICurrentUser): Promise<unknown> {
    return this._authService.getProfile(user.id);
  }
}
