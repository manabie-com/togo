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

  beforeEach(async () => {
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
    token = res.body.data;
  });

  describe("[CREATE] - POST /note/create", () => {
    it("Create Successfull: should return message [Successfull]", async() => {
      const dataBody  = {
        content: "test add task"
      }
      
      const res = await request(app)
        .post("/api/todo/create")
        .set("Authorization", `Bearer ${token}`)
        .send(dataBody);

      expect(res.status).toBe(200);

    })

    it("Create Failed: should return message [Missing content]", async () => {
      const dataBody = {
        content: "",
      };
      const res = await request(app)
        .post("/api/todo/create")
        .set("Authorization", `Bearer ${token}`)
        .send(dataBody);

      expect(res.status).toBe(400);
      expect(res.body.message).toBe("Missing content");
    });

    it("Create Failed: should return message [Limit task can add in day]", async () => {
      const dataBody = {
        content: "task 6",
      };
      const user = await userModel.findOne({ userName: "phanducanh" });
      const userId = user._id;
      await noteModel.create([
        {
          content: "task 1",
          user: userId,
        },
        {
          content: "task 2",
          user: userId,
        },
        {
          content: "task 3",
          user: userId,
        },
        {
          content: "task 4",
          user: userId,
        },
        {
          content: "task 5",
          user: userId,
        },
      ]);
      const res = await request(app)
        .post("/api/todo/create")
        .set("Authorization", `Bearer ${token}`)
        .send(dataBody);

      expect(res.status).toBe(400);
      expect(res.body.message).toBe("Limit task can add in day");
    });
  })

  describe("[TICK COMPLETE TASK] - PUT todo/task/tick/:id", () => {
    it("Tick task Successfull: should return message [Successfull]", async () => {
      const user = await userModel.findOne({ userName: "phanducanh" });
      const userId = user._id;
      
      const newTask  = await noteModel.create({
        user:userId,
        content: "test tick complete task"
      });
      const res = await request(app)
        .put(`/api/todo/task/tick/${newTask._id}`)
        .set("Authorization", `Bearer ${token}`);

      expect(res.status).toBe(200);
    });

    it("Tick task Faild: should return message [Opps, something went wrong]", async () => {
      const res = await request(app)
        .put("/api/todo/task/tick/123456")
        .set("Authorization", `Bearer ${token}`);

      expect(res.status).toBe(400);
      expect(res.body.message).toBe("Opps, something went wrong");
    });
  });

  describe("[GET TASK LIST] - GET todo/tasks?day=", () => {
    it("GET task Successfull: should return message [Successfull]", async () => {
      const user = await userModel.findOne({ userName: "phanducanh" });
      const userId = user._id;
      const day = "2021-12-06";
      await noteModel.create({
        _id: "61adaf5442e631017317c5b9",
        content: "test task day 3-2",
        user: userId,
        createdAt: "2021-12-06T06:36:04.102Z",
        updatedAt: "2021-12-06T06:36:04.102Z",
        __v: 0,
      });
      const res = await request(app)
        .get(`/api/todo/tasks?day=${day}`)
        .set("Authorization", `Bearer ${token}`);

      expect(res.status).toBe(200);
    });
  });

  describe("[UPDATE TASK] - PUT todo/task/update/:id", () => {
    it("update task Successfull: should return message [Successfull]", async () => {
      const user = await userModel.findOne({ userName: "phanducanh" });
      const userId = user._id;
      const dataBody = {
        content: "test update task"
      }
      const newTask  = await noteModel.create({
        user:userId,
        content: "test task"
      });
      const res = await request(app)
        .put(`/api/todo/task/update/${newTask._id}`)
        .set("Authorization", `Bearer ${token}`)
        .send(dataBody);
      
      expect(res.status).toBe(200);
    });

    it("update task Faild: should return message [Missing content]", async () => {
      const user = await userModel.findOne({ userName: "phanducanh" });
      const userId = user._id;
      const dataBody = {
        content: "",
      };
      
      const res = await request(app)
        .put("/api/todo/task/update/5a15s15w15ad15ad4")
        .set("Authorization", `Bearer ${token}`)
        .send(dataBody);

      expect(res.status).toBe(400);
      expect(res.body.message).toBe("Missing content");
    });

    it("update task Faild: should return message [ERROR]", async () => {
      const user = await userModel.findOne({ userName: "phanducanh" });
      const userId = user._id;
      const dataBody = {
        content: "test update task",
      };
      await noteModel.create({
        user: userId,
        content: "test task",
      });
      const res = await request(app)
        .put('/api/todo/task/update/5a15s15w15ad15ad4')
        .set("Authorization", `Bearer ${token}`)
        .send(dataBody);

      expect(res.status).toBe(400);
    });
  });

  describe("[DELETE TASK] - DELETE todo/task/delete/:id", () => {
    it("delete task Successfull: should return message [Successfull]", async () => {
      const user = await userModel.findOne({ userName: "phanducanh" });
      const userId = user._id;
      const newTask  = await noteModel.create({
        user: userId,
        content: "test task"
      });
      const res = await request(app)
        .delete(`/api/todo/task/delete/${newTask._id}`)
        .set("Authorization", `Bearer ${token}`);
      expect(res.status).toBe(200);
    });

    it("delete task Faild: should return message [Successfull]", async () => {
      const user = await userModel.findOne({ userName: "phanducanh" });
      const userId = user._id;
      await noteModel.create({
        user: userId,
        content: "test task",
      });
      const res = await request(app)
        .delete('/api/todo/task/delete/1a5d15aw4d8adad')
        .set("Authorization", `Bearer ${token}`);
      expect(res.status).toBe(400);
    });
  });
})