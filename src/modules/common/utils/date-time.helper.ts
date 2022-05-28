import { environment } from '@env/environment';
import { format } from 'date-fns-tz';

export const formatDate = (date: Date | string, f: string = 'yyyy-MM-dd HH:mm:ss.SSS'): string => {
  return date ? format(new Date(date), f, { timeZone: environment.timeZone }) : null;
};
