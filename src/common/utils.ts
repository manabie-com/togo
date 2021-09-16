import crypto from 'crypto';
import { getUnixTime } from 'date-fns';

export const unixTime = (): number => getUnixTime(new Date());

export const getSalt = (): string => crypto.randomBytes(16).toString('hex');
export const getHash = (password: string, salt: string): string =>
  crypto.pbkdf2Sync(password, salt, 10_000, 512, 'sha512').toString('hex');
