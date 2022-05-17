import { Body, Controller, Injectable, Post } from '@nestjs/common';
import { Public } from 'src/decorators/public.decorator';
import { AuthService } from './auth.service';
import { AuthSigninDto } from './dto/auth-signin.dto';
import { AuthSignupDto } from './dto/auth-signup.dto';

@Injectable()
@Controller('auth')
export class AuthController {
  constructor(private readonly authService: AuthService){}
  
  @Public()
  @Post('signup')
  async signup(@Body() registerUserDto: AuthSignupDto) {
    return await this.authService.signup(registerUserDto);
  }

  @Public()
  @Post('signin')
  async signin(@Body() loginUserDto: AuthSigninDto) {
    const token = await this.authService.signin(loginUserDto);
    return {
      token: token
    };
  }
}