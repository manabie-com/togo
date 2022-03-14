
module.exports = class BaseController {
  /**
 * Represents a base controller.
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
   * Redirect another page
   * @param {string} path - this is path
   */
  redirect(path) {
    this._res.redirect(path);
    this._res.end();
  }

  /**
  * renderJson will response json data
  * @param {Object} data - this is data
  */
  renderJson(data, message = '') {
    const response = {
      error: false,
      message: message,
      data: data
    }
    this._res.json(response);
    this._res.end();
  }

  /**
  * Render Error
  * @param {Erroe} error - this is error
  */
  error(error) {
    this._next(error);
  }

  /**
  * render error with message
  * @param {int} status
  * @param {string} error
  */
  errorWithMessage(status, error) {
    if (typeof error != 'object') {
      error = new Error(error);
    }
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
    error.status = 400;
    error.errors = errors;

    this._next(error);
  }
};
