import {NestInterceptor,ExecutionContext,CallHandler,Injectable} from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';

@Injectable()
export class CurrentUserInterceptor implements NestInterceptor {
  constructor(private jwtService: JwtService) {}

  async intercept(context: ExecutionContext, handler: CallHandler) {
    let request = context.switchToHttp().getRequest();

    const authheader = request.header('Authorization');
    const token = authheader && authheader.split(" ")[1];

    const decodedToken = this.jwtService.decode(token);
    request.body = { ...request.body, userId: decodedToken['id'] };

    return handler.handle();
  }
}