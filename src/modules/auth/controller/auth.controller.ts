import { Body, Controller, Post } from '@nestjs/common';
import { UserAuthDto } from '../dto/user-auth-request.dto';
import { AuthService } from '../service/auth.service';

@Controller('v1/auth')
export class AuthController {
  constructor(private readonly authService: AuthService) {}

  @Post('login')
  async login(@Body() request: UserAuthDto) {
    return await this.authService.login(request);
  }

  @Post('register')
  async register(@Body() request: UserAuthDto) {
    return await this.authService.register(request);
  }
}
