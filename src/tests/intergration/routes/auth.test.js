const request = require("supertest");
const httpStatus = require("http-status");
const mongoose = require("mongoose");

const { User } = require("../../../apis/models");
const mongooseLoader = require("../../../loaders/mongooseLoader");
const expressLoader = require("../../../loaders/expressLoader");

let server;

describe("/api/v1/auth", () => {
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

  describe("POST /register", () => {
    let userBody;

    const exec = async () => {
      return await request(server).post("/api/v1/auth/register").send(userBody);
    };

    beforeEach(() => {
      userBody = {
        name: "name1",
        email: "a@manabie.com",
        password: "password1",
        maxTask: 10,
      };
    });

    afterEach(async () => {
      await User.remove({});
    });

    it("should return 400 if user don't have a name", async () => {
      delete userBody.name;

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if user's name is empty", async () => {
      userBody.name = "";

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if user's name is less than 5 characters", async () => {
      userBody.name = "a";

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if user's name is greater than 50 characters", async () => {
      userBody.name = new Array(52).join("a");

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if user don't have an email", async () => {
      delete userBody.email;

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if user's email is empty", async () => {
      userBody.email = "";

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if user's email is less than 5 characters", async () => {
      userBody.email = "a";

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if user's email is greater than 255 characters", async () => {
      userBody.email = new Array(257).join("a");

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if user's email is incorrect format", async () => {
      userBody.email = "a";

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if user don't have a password", async () => {
      delete userBody.password;

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if user's password is empty", async () => {
      userBody.password = "";

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if user's password is less than 5 characters", async () => {
      userBody.password = "a";

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if user's password is greater than 1024 characters", async () => {
      userBody.password = new Array(1026).join("a");

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if user's password do not contain at least 1 letter and 1 number", async () => {
      userBody.password = "abcdefgh";

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if user don't have maxTask", async () => {
      delete userBody.maxTask;

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if maxTask is not a number", async () => {
      userBody.maxTask = "a";

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if maxTask equal 0", async () => {
      userBody.maxTask = 0;

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if maxTask less than 0", async () => {
      userBody.maxTask = -1;

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should save the user if it is valid", async () => {
      const res = await exec();
      const user = await User.find({ _id: res.body._id });

      expect(user).not.toBeNull();
      expect(res.status).toBe(httpStatus.CREATED);
    });

    it("should return the user if it is valid", async () => {
      const res = await exec();

      expect(res.body).toHaveProperty("user");
      expect(res.body.user.password).not.toBe(userBody.password);
      expect(res.body).toHaveProperty("token");
      expect(res.status).toBe(httpStatus.CREATED);
    });
  });

  describe("POST /login", () => {
    let loginBody;

    const exec = async () => {
      return await request(server)
        .post("/api/v1/auth/login")
        .send(loginBody);
    };

    beforeEach(() => {
      loginBody = {
        email: "a@manabie.com",
        password: "password1",
      };
    });

    afterEach(async () => {
      await User.remove({});
    });

    it("should return 400 if login info don't have email", async () => {
      delete loginBody.email;

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if login email is empty", async () => {
      loginBody.email = "";

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if login email is less than 5 characters", async () => {
      loginBody.email = "a";

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if login email is greater than 255 characters", async () => {
      loginBody.email = new Array(257).join("a");

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if login email is incorrect format", async () => {
      loginBody.email = "a";

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if login info don't have password", async () => {
      delete loginBody.password;

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if login password is empty", async () => {
      loginBody.password = "";

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if login password is less than 5 characters", async () => {
      loginBody.password = "a";

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if login password is greater than 1024 characters", async () => {
      loginBody.password = new Array(1026).join("a");

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if login password do not contain at least 1 letter and 1 number", async () => {
      loginBody.password = "abcdefgh";

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 401 if email or password is incorrect", async () => {
      const res = await exec();

      expect(res.status).toBe(httpStatus.UNAUTHORIZED);
    });

    it("should return user info and token if email and password are valid", async () => {
      const user = new User({
        name: "name1",
        email: "a@manabie.com",
        password: "password1",
        maxTask: 10,
      })
      await user.save();
      
      const res = await exec();

      expect(res.body).toHaveProperty("user");
      expect(res.body).toHaveProperty("token");
      expect(res.status).toBe(httpStatus.OK);
    });
  });
});
