const mongoose = require("mongoose");

const { User } = require("../../../apis/models");
const auth = require("../../../middlewares/auth");

describe("authMiddleware", () => {
  it("should populate req.user with the payload of a valid JWT", () => {
    const user = {
      _id: mongoose.Types.ObjectId().toHexString(),
      maxTask: 10,
    };
    const token = new User(user).generateAuthToken();

    const req = {
      headers: {
        authorization: `Bearer ${token}`,
      },
    };
    const res = {};
    const next = jest.fn();

    auth(req, res, next);

    expect(req.user).toMatchObject(user);
  });
});
