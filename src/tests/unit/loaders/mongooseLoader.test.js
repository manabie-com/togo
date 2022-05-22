const mongoose = require("mongoose");

const mongooseLoader = require("../../../loaders/mongooseLoader");

describe("mongooseLoader", () => {
  it("should throw if connect to mongo DB fail", async () => {
    mongoose.connect = jest
      .fn()
      .mockRejectedValue(new Error("Failed to connect to MongoDB"));

    const result = mongooseLoader()
    await expect(result).rejects.toThrowError(/Failed/);
  });
});
