const bcrypt = require("bcrypt");
const express = require("express");
const request = require("supertest");
const mongoose = require("mongoose");
const { MongoMemoryServer } = require("mongodb-memory-server");

const route = require("../../src/routes");
const userModel = require("../../src/models/user.model");

describe("[INTEGRATION TEST]: AUTH", () => {
  const app = express();
  app.use(express.urlencoded({ extended: false }));
  app.use(express.json());
  route(app);
  // const port = process.env.PORT || 4000;
  // const server = app.listen(port, () => {
  //   console.log(`Testing Server Is Running on: http://localhost:${port}`);
  // });
  const mongoMock = new MongoMemoryServer();

  beforeAll(async () => {
    (async () => {
      await mongoMock.start();
      const mongoUri = mongoMock.getUri();

      await mongoose.connect(mongoUri, {
        useNewUrlParser: true,
        useUnifiedTopology: true,
      });
    })();
  });

  afterEach(async () => {
    const collections = mongoose.connection.collections;

    for (const key in collections) {
      const collection = collections[key];
      await collection.deleteMany({});
    }
  });

  afterAll(async () => {
    await mongoose.connection.dropDatabase();
    await mongoose.connection.close();
    await mongoMock.stop();
    // server.close();
  });

  describe("[REGISTER] - POST /auth/register", () => {
    it("Register Successfull: should return message [Successfull]", (done) => {
      const registerBody = {
        userName: "phanducanh",
        password: "123456",
      };

      request(app)
        .post("/api/auth/register")
        .send(registerBody)
        .expect(200)
        .end((err, res) => {
          if (err) return done(err);
          expect(res.body.message).toStrictEqual("Successful");
          done();
        });
    });

    it("Register Failed wrong username: should return message [Please do not leave the username blank]", (done) => {
      const registerBody = {
        userName: "",
        password: "123456",
        repassword: "123456",
      };

      request(app)
        .post("/api/auth/register")
        .send(registerBody)
        .expect(400)
        .end((err, res) => {
          if (err) return done(err);
          expect(res.body.message).toStrictEqual(
            "Please do not leave the username blank"
          );
          done();
        });
    });

    it("Register Failed wrong password: should return message [Password must be at least 6 characters]", (done) => {
      const registerBody = {
        userName: "phanducanh",
        password: "11",
        repassword: "11",
      };

      request(app)
        .post("/api/auth/register")
        .send(registerBody)
        .expect(400)
        .end((err, res) => {
          if (err) return done(err);
          expect(res.body.message).toStrictEqual(
            "Password must be at least 6 characters"
          );
          done();
        });
    });
  });

  describe("[LOGIN] - POST /auth/login", () => {
    it("Login Successful", async () => {
      const loginBody = {
        userName: "phanducanh",
        password: "123456",
      };
      const hashedPassword = await bcrypt.hash("123456", 10);
      await userModel.create({
        userName: "phanducanh",
        password: hashedPassword,
      });

      const res = await request(app).post("/api/auth/login").send(loginBody);
      expect(res.status).toBe(200);
    });

    it("Login Failed wrong password: should return message [Wrong username or password]", async () => {
      const loginBody = {
        userName: "phanducanh",
        password: "12345688",
      };

      const hashedPassword = await bcrypt.hash("123456", 10);
      await userModel.create({
        userName: "phanducanh",
        password: hashedPassword,
      });

      const res = await request(app).post("/api/auth/login").send(loginBody);
      expect(res.status).toBe(400);
      expect(res.body.message).toBe("Wrong username or password");
    });

    it("Login Failed wrong username: should return message [Wrong username or password]", async () => {
      const loginBody = {
        userName: "usernametest",
        password: "12345688",
      };

      const hashedPassword = await bcrypt.hash("123456", 10);
      await userModel.create({
        userName: "phanducanh",
        password: hashedPassword,
      });

      const res = await request(app).post("/api/auth/login").send(loginBody);
      expect(res.status).toBe(400);
      expect(res.body.message).toBe("Wrong username or password");
    });

    it("Login Failed with missing username: should return message [Please do not leave the username blank]", async () => {
      const loginBody = {
        userName: "",
        password: "123456",
      };

      const hashedPassword = await bcrypt.hash("123456", 10);
      await userModel.create({
        userName: "phanducanh",
        password: hashedPassword,
      });

      const res = await request(app).post("/api/auth/login").send(loginBody);
      expect(res.status).toBe(400);
      expect(res.body.message).toBe("Please do not leave the username blank");
    });

    it("Login Failed with missing password: should return message [Please do not leave the password blank]", async () => {
      const loginBody = {
        userName: "phanducanh",
        password: "",
      };

      const hashedPassword = await bcrypt.hash("123456", 10);
      await userModel.create({
        userName: "phanducanh",
        password: hashedPassword,
      });

      const res = await request(app).post("/api/auth/login").send(loginBody);
      expect(res.status).toBe(400);
      expect(res.body.message).toBe("Please do not leave the password blank");
    });

    it("Login Failed with wrong format password: should return message [Password must be at least 6 characters]", async () => {
      const loginBody = {
        userName: "phanducanh",
        password: "123",
      };

      const hashedPassword = await bcrypt.hash("123456", 10);
      await userModel.create({
        userName: "phanducanh",
        password: hashedPassword,
      });

      const res = await request(app).post("/api/auth/login").send(loginBody);
      expect(res.status).toBe(400);
      expect(res.body.message).toBe("Password must be at least 6 characters");
    });
  });
});
