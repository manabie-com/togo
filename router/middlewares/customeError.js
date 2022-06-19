/**
 * @author Nguyen Minh Tam / ngmitamit@gmail.com
 */

 const constants = require('../constants');

 const { errorConstants } = constants;
 
 class CustomError extends Error {
     constructor(code, data) {
         super(errorConstants[code] || code || 'Unknown Error!');
 
         // Maintains proper stack trace for where our error was thrown (only available on V8)
         if (Error.captureStackTrace) {
             Error.captureStackTrace(this, CustomError);
         }
 
         this.error_data = data;
         this.error_code = code;
     }
 }
 module.exports = CustomError;
 