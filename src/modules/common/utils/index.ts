import * as crypto from 'crypto';

export const createPasswordHash = (password: string): string => {
  const sha256 = crypto.createHash('sha256');

  return sha256.update(password, 'utf8').digest('hex');
};

export const checkSignInType = (username: string): 'email' | 'username' => {
  // eslint-disable-next-line security/detect-unsafe-regex
  const emailRegex = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;

  if (emailRegex.test(String(username).toLowerCase())) {
    return 'email';
  } else {
    return 'username';
  }
};

export const generateID = (count: number) => {
  const sym = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890';
  let str = '';

  for (let i = 0; i < count; i++) {
    const idx = Math.random() * sym.length;

    str += sym.charAt(idx);
  }

  return str;
};

// use safeKey = for easy tinkering in the console.
export const safeKey = (() => {
  // Safely allocate plainObject's inside iife
  // Since this function may get called very frequently -
  // I think it's important to have plainObject's
  // statically defined
  const obj = {};
  const arr = [];

  // ...if for some reason you ever use square brackets on these types...
  // const fun = function() {}
  // const bol = true;
  // const num = 0;
  // const str = '';
  return (key) => {
    // eslint-disable-next-line security/detect-object-injection
    if (obj[key] !== undefined || arr[key] !== undefined) {
      return `SAFE_${key}`;
    } else {
      return key;
    }
  };
})();
