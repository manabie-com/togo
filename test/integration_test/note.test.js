const bcrypt = require("bcrypt");
const express = require("express");
const request = require("supertest");
const mongoose = require("mongoose");
const { MongoMemoryServer } = require("mongodb-memory-server");

const route = require("../../src/routes");
const userModel = require("../../src/models/user.model");
const noteModel = require("../../src/models/note.model");

describe("[INTEGRATION TEST]: NOTE", () => {
  const app = express();
  app.use(express.urlencoded({ extended: false }));
  app.use(express.json());
  route(app);
  var token
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
    const hashedPassword = await bcrypt.hash("123456", 10);
    await userModel.create({
      userName: "phanducanh",
      password: hashedPassword,
    });
    const loginBody = {
      userName: "phanducanh",
      password: "123456",
    };
    
    const res = await request(app).post("/api/auth/login").send(loginBody);
    token = res.body.data
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

  describe("[CREATE] - POST /note/create", () => {
    it("Create Successfull: should return message [Successfull]", async() => {
      const dataBody  = {
        content: "test add task"
      }
      const bearerToken = `Bearer ${token}`;
      console.log(bearerToken);
      const res = await request(app).post("/api/todo/create",{
        'auth': {
          'bearer': token,
        }
      }).send(dataBody);

      expect(res.status).toBe(200);

    })
  })
})