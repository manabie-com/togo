const express = require("express");
const request = require("supertest");
const mongoose = require("mongoose");
const {
  MongoMemoryServer
} = require("mongodb-memory-server");

const route = require("../../src/routers");
const taskModel = require("../../src/models/task");
const userModel = require("../../src/models/user");
const userTaskLimitModel = require("../../src/models/taskLimit");
const moment = require("moment")

describe("[INTEGRATION TEST]: TASK", () => {
  const app = express();
  app.use(express.urlencoded({
    extended: false
  }));
  app.use(express.json());
  route(app);
  var token, user;
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
    const data = {
      email: "hieunhan2000@gmail.com",
      password: "1234567"
    }
    user = await userModel.create({
      ...data,
      name: "Nhan Phan"
    });

    const res = await request(app).post("/user/login").send(data);
    token = res.body.token;
  });

  describe("[CREATE] - POST /task/create", () => {
    it("Create Successfull: should return the created task detail", async () => {
      const dataBody = {
        title: "task 1",
        description: "Proin eget tortor risus. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque in ipsum id orci porta dapibus."
      }

      const res = await request(app)
        .post("/task")
        .set("Authorization", `Bearer ${token}`)
        .send(dataBody);

      expect(res.status).toBe(201);
    })

    it("Create Failed: Please do not leave the title blank", async () => {
      const dataBody = {
        title: "",
        description: "Proin eget tortor risus. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque in ipsum id orci porta dapibus."
      }

      const res = await request(app)
        .post("/task")
        .set("Authorization", `Bearer ${token}`)
        .send(dataBody);

      expect(res.status).toBe(400);
      expect(res.body.error).toBe("Please do not leave the title blank");
    })

    it("Create Failed: Please do not leave the description blank", async () => {
      const dataBody = {
        title: "task 1",
        description: ""
      }

      const res = await request(app)
        .post("/task")
        .set("Authorization", `Bearer ${token}`)
        .send(dataBody);

      expect(res.status).toBe(400);
      expect(res.body.error).toBe("Please do not leave the description blank");
    })

    it("Create Failed: should return message [Limit task can add in day]", async () => {
      const dataBody = {
        title: "task 3",
        description: "Vivamus magna justo, lacinia eget consectetur sed, convallis at tellus. Vivamus suscipit tortor eget felis porttitor volutpat."
      };
      const userId = user._id;

      await userTaskLimitModel.create({
        quantity: 2,
        atDate: moment().toDate(),
        createdDate: moment().toDate(),
        userId: userId
      })

      await taskModel.create([{
          title: "task 1",
          description: "Vivamus magna justo, lacinia eget consectetur sed, convallis at tellus. Vivamus suscipit tortor eget felis porttitor volutpat.",
          createdById: userId,
          createdDate: moment().toDate()
        },
        {
          title: "task 2",
          description: "Vivamus magna justo, lacinia eget consectetur sed, convallis at tellus. Vivamus suscipit tortor eget felis porttitor volutpat.",
          createdById: userId,
          createdDate: moment().toDate()
        }
      ]);
      const res = await request(app)
        .post("/task")
        .set("Authorization", `Bearer ${token}`)
        .send(dataBody);

      expect(res.status).toBe(400);
      expect(res.body.error).toBe("The quantity task of the user is over the limit");
    });  
  })
  describe("[GET TASK LIST] - GET /task?day=", () => {
    it("GET task Successfull: should return message [Successfull]", async () => {
      const day = moment().toString();
      const data = [{
          title: "task 1",
          description: "Vivamus magna justo, lacinia eget consectetur sed, convallis at tellus. Vivamus suscipit tortor eget felis porttitor volutpat.",
          createdById: user._id,
          createdDate: moment().add(1, 'h')
        },
        {
          title: "task 2",
          description: "Vivamus magna justo, lacinia eget consectetur sed, convallis at tellus. Vivamus suscipit tortor eget felis porttitor volutpat.",
          createdById: user._id,
          createdDate: moment().subtract(1, 'd')
        },
        {
          title: "task 3",
          description: "Vivamus magna justo, lacinia eget consectetur sed, convallis at tellus. Vivamus suscipit tortor eget felis porttitor volutpat.",
          createdById: user._id,
          createdDate: moment().add(1, 'h')
        },
        {
          title: "task 4",
          description: "Vivamus magna justo, lacinia eget consectetur sed, convallis at tellus. Vivamus suscipit tortor eget felis porttitor volutpat.",
          createdById: user._id,
          createdDate: moment().add(1, 'd')
        },
      ]
      
      const newTasks = await taskModel.create(data);
      const expectResult = [data[0], data[2]]
        
      const res = await request(app)
        .get(`/task?day=${day}`)
        .set("Authorization", `Bearer ${token}`);

      
      expect(res.status).toBe(200);
      expect(res.body.tasks[0].title).toEqual(expectResult[0].title);
      expect(res.body.tasks[1].title).toEqual(expectResult[1].title);
    });
  });

  describe("[DELETE TASK] - DELETE /task/:id", () => {
    it("Delete task Successfull: should return message [Successfull]", async () => {

      const newTask = await taskModel.create({
        title: "task 1",
        description: "Vivamus magna justo, lacinia eget consectetur sed, convallis at tellus. Vivamus suscipit tortor eget felis porttitor volutpat. edited",
        createdById: user._id,
        createdDate: moment().toDate()
      });

      const res = await request(app)
        .delete(`/task/${newTask._id}`)
        .set("Authorization", `Bearer ${token}`);
      expect(res.status).toBe(200);
    });

    it("Delete task Failed: should return message [Task is not exist]", async () => {

      await taskModel.create({
        title: "task 1",
        description: "Vivamus magna justo, lacinia eget consectetur sed, convallis at tellus. Vivamus suscipit tortor eget felis porttitor volutpat. edited",
        createdById: user._id,
        createdDate: moment().toDate()
      });

      const res = await request(app)
        .delete(`/task/61c0cff0d1c8cdf1cda94153`)
        .set("Authorization", `Bearer ${token}`)

      expect(res.status).toBe(400);
      expect(res.body.error).toBe("Task is not exist");
    });

  })
})