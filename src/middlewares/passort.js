'use strict!';


const { verify } = require('jsonwebtoken');
const { verifyToken } = require('../services/services.user');


module.exports.passport = () => {
  return async (ctx, next) => {
    const url = ctx.request.url || '';
    let isPrivate = !url.startsWith('/api/public');

    if (isPrivate) {
      const token = (ctx.request.headers?.authorization || '').replace('Bearer', '').trim();
      if (!token) {
        ctx.status = 401;
        return ctx.body = {
          message: 'Request token required!'
        }
      }

      const valid = await verifyToken(token);
      if (!valid) {
        ctx.status = 401;
        return ctx.body = {
          message: 'Invalid JWT Token or Token is expired!'
        }
      }
      ctx.user = {
        sub: valid.sub
      }
    }

    return await next();
  }
}