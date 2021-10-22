import { AuthGuard } from '@nestjs/passport';
import {
  ExecutionContext,
  Injectable,
  HttpException,
  HttpStatus,
} from '@nestjs/common';

@Injectable()
export class JwtAuthGuard extends AuthGuard('jwt') {
  canActivate(context: ExecutionContext) {
    // Add your custom authentication logic here
    // for example, call super.logIn(request) to establish a session.
    return super.canActivate(context);
  }
  handleRequest(err: any, user: any) {
    // You can throw an exception based on either "info" or "err" arguments
    if (err) {
      throw err;
    }
    if (!user) {
      throw new HttpException(
        {
          statusCode: HttpStatus.UNAUTHORIZED,
          message: 'UNAUTHORIZED',
        },
        HttpStatus.UNAUTHORIZED,
      );
    }
    return user;
  }
}
