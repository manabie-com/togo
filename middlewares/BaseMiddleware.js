const createError = require('http-errors');

module.exports = class BaseMiddleware {
  /**
 * Represents a base middleware.
 * @param {object} req - this is request.
 * @param {object} res - this is response the request
 * @param {object} next - this is next of the request
 */
  constructor(req, res, next) {
    this._req = req;
    this._res = res;
    this._next = next;
  }

  /**
  * render error with message
  * @param {int} status
  * @param {string} message
  */
  errorWithMessage(status, message) {
    const error = new Error(message);
    error.status = status;
    this._next(error);
  }

  /**
  * render bad request
  * @param {array} errors
  */
  badRequest(errors) {
    const message = 'The request was invalid or cannot be otherwise served.';
    const error = new Error(message);
    error.status = createError.BadRequest;
    error.errors = errors;
    this._next(error);
  }

  /**
   * agent info
   */
  agentInfo() {
    const useragent = this._req.useragent;

    this._req.body.agent = {
      mobile: useragent.isMobile,
      browser_name: useragent.browser,
      browser_version: useragent.version,
      os: useragent.os,
      platform: useragent.platform,
      source: useragent.source,
    };

    this._req.body.user_ip = (this._req.header('x-forwarded-for')
      || this._req.connection.remoteAddress).replace(/^.*:/, '');
    this._next();
  }
};
