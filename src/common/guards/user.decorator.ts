import { createParamDecorator, ExecutionContext } from '@nestjs/common';
import { UserAccount } from './user-account.class';

export const User = createParamDecorator(
  (data: unknown, ctx: ExecutionContext) => {
    const arg = ctx.getArgByIndex(0);
    const user = arg.user ? new UserAccount(arg.user) : null;
    return user;
  },
);
