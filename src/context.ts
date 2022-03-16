import { ApolloError } from 'apollo-server-express';
import { verify } from 'jsonwebtoken';

interface IVerify {
  userId?: number;
}

export function getUserId(ctx) {
  const Authorization = ctx.req.get('Authorization') ?? null;
  if (!Authorization) {
    throw new Error('Not Authorization');
  }
  const verifiedToken = verify(Authorization, process.env.APP_SECRET);
  return (verifiedToken as IVerify)?.userId;
}
