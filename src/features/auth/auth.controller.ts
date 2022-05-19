import { Body, Controller, Injectable, Post } from '@nestjs/common';
import { Public } from '../../decorators/public.decorator';
import { AuthService } from './auth.service';
import { AuthSigninDto } from './dto/auth-signin.dto';

@Injectable()
@Controller()
export class AuthController {
  constructor(private readonly authService: AuthService){}

  @Public()
  @Post('signin')
  async signin(@Body() loginUserDto: AuthSigninDto) {
    const token = await this.authService.signin(loginUserDto);
    return {
      token: token
    };
  }
}