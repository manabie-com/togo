//The error class that contains the error data

class HttpError extends Error {
  constructor(message = undefined, errorCode = undefined) {
    super(message);
    this.code = errorCode;
  }
}

module.exports = HttpError;
