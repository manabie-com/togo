import * as morgan from 'morgan';

export const logger = morgan((tokens, req, res) => {
  const forwarded = (req.headers['x-forwarded-for'] || req.ip || 'unknown').replace(' ', '').split(',');
  const uid = (req.headers['x-authenticated-user'] || 'anonymous').trim();

  return [
    tokens.method(req, res),
    tokens.url(req, res),
    tokens.status(req, res),
    tokens.res(req, res, 'content-length'),
    '-',
    tokens['response-time'](req, res),
    'ms',
    forwarded[0],
    uid,
  ].join(' ');
});
