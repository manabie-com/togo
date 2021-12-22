const express = require("express");
const request = require("supertest");
const mongoose = require("mongoose");
const bcrypt = require("bcryptjs");
const {
  MongoMemoryServer
} = require("mongodb-memory-server");
const route = require("../../src/routers");
const userModel = require('../../src/models/user');

describe("[INTEGRATION TEST]: USER", () => {
  const app = express();
  app.use(express.urlencoded({
    extended: false
  }));
  app.use(express.json());
  route(app);

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

  describe("[REGISTER] - POST /user", () => {
    it("Register Successfull: should return message [Successfull]", (done) => {
      const registerBody = {
        email: "tuyennguyen@gmail.com",
        password: "1234567",
        name: "Tuyen Nguyen"
      };

      request(app)
        .post("/user/register")
        .send(registerBody)
        .expect(201)
        .end((err, res) => {
          if (err) return done(err);
          expect(res.body.user.email).toStrictEqual(registerBody.email);
          done();
        });
    });
    it("Register Failed wrong email: should return message [Please do not leave the email blank]", (done) => {
      const registerBody = {
        email: "",
        password: "1234567",
        name: "Tuyen Nguyen"
      };

      request(app)
        .post("/user/register")
        .send(registerBody)
        .expect(400)
        .end((err, res) => {
          if (err) return done(err);
          expect(res.body.error).toStrictEqual(
            "Please do not leave the email blank"
          );
          done();
        });
    });
    
    it("Register Failed blank password: should return message [Please do not leave the password blank]", (done) => {
      const registerBody = {
        email: "tuyennguyen@gmail.com",
        password: "",
        name: "Tuyen Nguyen"
      };

      request(app)
        .post("/user/register")
        .send(registerBody)
        .expect(400)
        .end((err, res) => {
          if (err) return done(err);
          expect(res.body.error).toStrictEqual(
            "Please do not leave the password blank"
          );
          done();
        });
    });
    it("Register Failed blank name: should return message [Please do not leave the name blank]", (done) => {
      const registerBody = {
        email: "tuyennguyen@gmail.com",
        password: "1234567",
        name: ""
      };

      request(app)
        .post("/user/register")
        .send(registerBody)
        .expect(400)
        .end((err, res) => {
          if (err) return done(err);
          expect(res.body.error).toStrictEqual(
            "User validation failed: name: Path `name` is required."
          );
          done();
        });
    });
    it("Register Failed too short lenght of password: should return message [Password must be at least 7 characters]", (done) => {
      const registerBody = {
        email: "tuyennguyen@gmail.com",
        password: "1234",
        name: "tuyet"
      };

      request(app)
        .post("/user/register")
        .send(registerBody)
        .expect(400)
        .end((err, res) => {
          if (err) return done(err);
          expect(res.body.error).toStrictEqual(
            "Password must be at least 7 characters"
          );
          done();
        });
    });
  })
  describe("[LOGIN] - POST /user/login", () => {
    it("Login success", async () => {
      const body = {
        email: "hieunhan2000@gmail.com",
        password: "1234567"
      };

      await userModel.create({
        email: body.email,
        password: body.password,
        name: "Nhan Phan"
      });

      const res = await request(app).post("/user/login").send(body);
      expect(res.status).toBe(200);
    });

    it("Login Failed wrong email: should return message [User is not exist]", async () => {
      const body = {
        email: "hieunhan123@gmail.com",
        password: "1234567"
      };

      await userModel.create({
        email: "hieunhan2000@gmail.com",
        password: body.password,
        name: "Nhan Phan"
      });

      const res = await request(app).post("/user/login").send(body);
      expect(res.status).toBe(400);
      expect(res.body.error).toBe("User is not exist");
    });

    it("Login Failed wrong password: should return message [Password is wrong]", async () => {
      const body = {
        email: "hieunhan2000@gmail.com",
        password: "1234567789"
      };

      await userModel.create({
        email: "hieunhan2000@gmail.com",
        password: "1234567",
        name: "Nhan Phan"
      });

      const res = await request(app).post("/user/login").send(body);
      expect(res.status).toBe(400);
      expect(res.body.error).toBe("Password is wrong");
    });

    it("Login Failed with missing email: should return message [Please do not leave the email blank]", async () => {
      const body = {
        email: "",
        password: "1234567789"
      };

      const res = await request(app).post("/user/login").send(body);
      expect(res.status).toBe(400);
      expect(res.body.error).toBe("Please do not leave the email blank");
    });

    it("Login Failed with missing password: should return message [Please do not leave the password blank]", async () => {
      const body = {
        email: "hieunhan2000@gmail.com",
        password: ""
      };

      const res = await request(app).post("/user/login").send(body);
      expect(res.status).toBe(400);
      expect(res.body.error).toBe("Please do not leave the password blank");
    });
    it("Login Failed with missing password: should return message [Password must be at least 7 characters]", async () => {
      const body = {
        email: "hieunhan2000@gmail.com",
        password: "123456"
      };

      const res = await request(app).post("/user/login").send(body);
      expect(res.status).toBe(400);
      expect(res.body.error).toBe("Password must be at least 7 characters");
    });
  })
})