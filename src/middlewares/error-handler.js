'use strict';


const moment = require('moment-timezone');

module.exports.errorHandler = () => {
  // Catch all exception error in context handler
  return async (ctx, next) => {
    try {
      await next()
    } catch (err) {
      const logMessage = {
        timestamp: moment().unix(),
        message: err.message,
        raw: JSON.stringify(err)
      }

      return ctx.showError(err.message);
      // TODO: Log this payload to trace log table 
    }
  }
}