const httpStatus = require("http-status");
const mongoose = require("mongoose");

const ApiError = require("../../../utils/api-error");
const {
  errorConverter,
  errorHandler,
  logger,
} = require("../../../middlewares/error");
const env = require("../../../configs/env");

describe("errorConverter", () => {
  it("should convert error to ApiError with status code is 500", () => {
    const req = {};
    const res = {};
    const next = jest.fn();

    const err = new Error("a");

    errorConverter(err, req, res, next);

    expect(next).toHaveBeenCalled();
    expect(next).toBeCalledWith(
      new ApiError(httpStatus.INTERNAL_SERVER_ERROR, "a")
    );
  });

  it("should convert error to ApiError with status code is 400 when error is instance of mongoose error", () => {
    const req = {};
    const res = {};
    const next = jest.fn();

    const err = new mongoose.Error();

    errorConverter(err, req, res, next);

    expect(next).toHaveBeenCalled();
    expect(next).toBeCalledWith(new ApiError(httpStatus.BAD_REQUEST, httpStatus[httpStatus.BAD_REQUEST]));
  });
});

describe("errorHandler", () => {
  it("should return 400 when error throw 400", () => {
    const req = {};
    const res = {
      status: jest.fn().mockReturnValue({
        json: jest.fn(),
      }),
    };

    const next = jest.fn();

    const err = new ApiError(httpStatus.NOT_FOUND, "a");

    errorHandler(err, req, res, next);
    expect(res.status).toHaveBeenCalledWith(httpStatus.NOT_FOUND);
  });

  it("should return 500 when error throw from production environment and no throw from operation", () => {
    const req = {};
    const res = {
      status: jest.fn().mockReturnValue({
        json: jest.fn(),
      }),
    };

    const next = jest.fn();
    env.isProduction = jest.fn().mockReturnValue(true);
    const err = new Error("a");

    errorHandler(err, req, res, next);
    expect(res.status).toHaveBeenCalledWith(httpStatus.INTERNAL_SERVER_ERROR);
  });

  it("should return log error to terminal when app is running in development environment", () => {
    const req = {};
    const res = {
      status: jest.fn().mockReturnValue({
        json: jest.fn(),
      }),
    };

    const next = jest.fn();
    env.isDevelopment = jest.fn().mockReturnValue(true);
    logger.error = jest.fn();
    const err = new Error("a");

    errorHandler(err, req, res, next);
    expect(logger.error).toHaveBeenCalled();
  });
});
