const httpStatus = require("http-status");

const ApiError = require("../../../utils/api-error");
const { errorConverter, errorHandler } = require("../../../middlewares/error");

describe("errorConverter", () => {
  it("should convert error to ApiError", () => {
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
});

describe("errorHandler", () => {
  it("should return http response", () => {
    const req = {};
    const res = {
      status: jest.fn().mockReturnValue({
        json: jest.fn(),
      }),
    };

    const next = jest.fn();

    const err = new ApiError(httpStatus.INTERNAL_SERVER_ERROR, "a");

    errorHandler(err, req, res, next);
    expect(res.status).toHaveBeenCalled();
  });
});
