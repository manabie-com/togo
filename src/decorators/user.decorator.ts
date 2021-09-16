import { createParamDecorator, ExecutionContext } from '@nestjs/common';

export interface ICurrentUser {
  id: number;
  email: string;
}

export const CurrentUser = createParamDecorator(
  (_data: unknown, ctx: ExecutionContext) => {
    const request = ctx.switchToHttp().getRequest();
    return request.user as ICurrentUser;
  }
);
