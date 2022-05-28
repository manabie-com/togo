import { createParamDecorator, ExecutionContext } from '@nestjs/common';
import { JWTPayload } from '@modules/auth/dto/jwt-payload.model';
import { Request } from 'express';
import { TokenInvalidOrExpiredBadRequestException } from '@modules/common/exceptions';

export const Payload = createParamDecorator((data: unknown, ctx: ExecutionContext): any => {
  const req: Request = ctx.switchToHttp().getRequest();

  if (req.user instanceof JWTPayload) {
    return req.user;
  } else {
    throw new TokenInvalidOrExpiredBadRequestException();
  }
});
