import { Injectable, Logger } from '@nestjs/common';
import { AuthGuard } from '@nestjs/passport';

@Injectable()
export class JwtAuthGuard extends AuthGuard('jwt') {
  getRequest(context) {
    const req = context.switchToHttp().getRequest();
    const { raw } = req;
    if (raw.cookies && raw.cookies.jwt) {
      Logger.log('raw cookies');
      context.switchToHttp().getRequest().headers[
        'authorization'
      ] = `Bearer ${raw.cookies.jwt}`;
      context.switchToHttp().getRequest().raw.headers[
        'Authorization'
      ] = `Bearer ${raw.cookies.jwt}`;
      context
        .switchToHttp()
        .getRequest().raw.rawHeaders[1] = `Bearer ${raw.cookies.jwt}`;
    }
    return context.switchToHttp().getRequest();
  }
}
