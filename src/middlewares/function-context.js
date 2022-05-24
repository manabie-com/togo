'use strict';


module.exports.functionContext = () => {
  return async (ctx, next) => {
    // In context after passport always return 200 status. Other error or return in body response.
    ctx.showResult = (result, code = 200) => {
      ctx.status = 200;
      ctx.body = {
        success: true,
        code: code,
        data: result
      }
    }

    // On error functuon approve only string input
    ctx.showError = (message, code = 400) => {
      if (typeof (message) !== 'string') {
        message = 'Something went wrong!'
      }
      ctx.status = 200;
      ctx.body = {
        success: false,
        code: code,
        message: message
      }
    }

    return await next();
  }
}