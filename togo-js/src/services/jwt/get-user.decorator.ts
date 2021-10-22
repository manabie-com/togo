import { createParamDecorator, ExecutionContext } from '@nestjs/common';

export const GetUser = createParamDecorator((data: string, context: ExecutionContext) => {
  const request = context;
  const user = request.getArgs();
  return user[0].user && user[0].user['id'] ? user[0].user['id'] : null;
});
