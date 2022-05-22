const mongoose = require("mongoose");

const common = require("../../../utils/common");

describe("getStartOfDay", () => {
  it("should return start of current day in milliseconds", () => {
    expect(common.getStartOfDay()).toBeGreaterThan(0);
  });
});

describe("objectIdFromDate", () => {
  it("should return object id from date in milliseconds", () => {
    const result = common.objectIdFromDate(1653066000000);

    expect(mongoose.Types.ObjectId.isValid(result)).toBeTruthy();
  });
});
