const request = require("supertest");
const httpStatus = require("http-status");
const mongoose = require("mongoose");
const jwt = require("jsonwebtoken");

const { User } = require("../../../apis/models");
const mongooseLoader = require("../../../loaders/mongooseLoader");
const expressLoader = require("../../../loaders/expressLoader");
const { authService } = require("../../../apis/services");

let server;

describe("auth middleware", () => {
  beforeAll(async () => {
    await mongooseLoader();
  });
  beforeEach(() => {
    server = expressLoader.server;
  });
  afterEach(async () => {
    await server.close();
    await User.remove({});
  });
  afterAll(async () => {
    mongoose.connection.close();
    await server.close();
  });

  let token;

  const exec = async () => {
    return await request(server)
      .get("/api/v1/tasks")
      .set("authorization", `Bearer ${token}`);
  };

  beforeEach(() => {
    token = new User({
      _id: mongoose.Types.ObjectId().toHexString(),
      maxTask: 1,
    }).generateAuthToken();
  });

  it("should return 401 if no token is provided", async () => {
    token = "";

    const res = await exec();

    expect(res.status).toBe(httpStatus.UNAUTHORIZED);
  });

  it("should return 400 if token is invalid", async () => {
    token = "a";

    const res = await exec();

    expect(res.status).toBe(httpStatus.BAD_REQUEST);
  });

  it("should return 400 if token not a bearer token", async () => {
    token = "Bearer a";

    const res = await exec();

    expect(res.status).toBe(httpStatus.BAD_REQUEST);
  });

  it("should return 401 if token not a bearer token", async () => {
    token = "Bearer a";

    authService.getTokenFromHeaders = jest.fn().mockReturnValue("a");
    jwt.verify = jest.fn().mockReturnValue(new Error("a"));
    const res = exec();

    expect(res).rejects.toThrow();
  });

  it("should return 200 if token is valid", async () => {
    const res = await exec();

    expect(res.status).toBe(httpStatus.OK);
  });
});
