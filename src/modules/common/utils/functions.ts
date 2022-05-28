import { safeKey } from './';
import { formatDate } from './date-time.helper';

/**
 * Pick object with none empty value
 * @param obj - The object you want to filter
 * @param isStrict - Exclude string of empty
 */
export const pickNotEmpty = (
  obj: Record<string, any> | string | number | boolean | Date,
  isStrict: boolean = false,
) => {
  const newObj = {};

  if (typeof obj === 'string' || typeof obj === 'number' || typeof obj === 'boolean') {
    return obj;
  }

  Object.keys(obj).forEach((key) => {
    if (obj[safeKey(key)] && Array.isArray(obj[safeKey(key)])) {
      // isArray
      const arr = obj[safeKey(key)];

      newObj[safeKey(key)] = [];
      arr.forEach((d: string | number | boolean | Record<string, any>) => {
        newObj[safeKey(key)].push(pickNotEmpty(d, isStrict));
      });
    } else if (obj[safeKey(key)] && typeof obj[safeKey(key)] === 'object') {
      // isObject

      if (Object.prototype.toString.call(obj[safeKey(key)]) === '[object Date]') {
        newObj[safeKey(key)] = formatDate(obj[safeKey(key)], 'yyyy-MM-dd');
      } else {
        newObj[safeKey(key)] = pickNotEmpty(obj[safeKey(key)], isStrict);
      }
    } else if (isStrict && obj[safeKey(key)] === '') {
      // pass
    } else if (obj[safeKey(key)] !== undefined && obj[safeKey(key)] !== null) {
      newObj[safeKey(key)] = obj[safeKey(key)];
    }
  });

  return newObj;
};
