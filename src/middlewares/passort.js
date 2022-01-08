'use strict!';

module.exports.passport = () => {
  return async (ctx, next) => {
    return await next();
  }
}