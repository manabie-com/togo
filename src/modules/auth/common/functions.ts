import { Request } from 'express';

export const getJWTToken = (req: Request): string | undefined => {
  try {
    if (!req.headers['authorization']) {
      return undefined;
    }

    const re = /(Bearer)\s+(\S+)/;
    const matches = req.headers['authorization'].match(re);

    if (matches.length < 3) {
      return undefined;
    }

    return matches[2];
  } catch {
    return undefined;
  }
};
