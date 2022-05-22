const request = require("supertest");
const httpStatus = require("http-status");
const mongoose = require("mongoose");

const { Task } = require("../../../apis/models");
const { User } = require("../../../apis/models");
const mongooseLoader = require("../../../loaders/mongooseLoader");
const expressLoader = require("../../../loaders/expressLoader");
const { priorityTypes } = require("../../../configs/priority");

describe("/api/v1/tasks", () => {
  beforeAll(async () => {
    await mongooseLoader();
  });
  beforeEach(() => {
    server = expressLoader.server;
  });
  afterEach(async () => {
    await server.close();
    await Task.remove({});
  });
  afterAll(async () => {
    mongoose.connection.close();
    await server.close();
  });

  describe("POST /", () => {
    let token;
    let taskBody;

    const exec = async () => {
      return await request(server)
        .post("/api/v1/tasks")
        .set("authorization", `Bearer ${token}`)
        .send(taskBody);
    };

    beforeEach(() => {
      token = new User({
        _id: mongoose.Types.ObjectId().toHexString(),
        maxTask: 1,
      }).generateAuthToken();
      taskBody = {
        title: "title1",
        description: "description1",
        priority: priorityTypes.HIGH,
        completed: false,
      };
    });

    afterEach(async () => {
      await Task.remove({});
    });

    it("should return 401 if client is not logged in", async () => {
      token = "";

      const res = await exec();

      expect(res.status).toBe(httpStatus.UNAUTHORIZED);
    });

    it("should return 400 if task don't have a title", async () => {
      delete taskBody.title;

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if task's title is empty", async () => {
      taskBody.title = "";

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if task's title is not a string", async () => {
      taskBody.title = true;

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if task's description is not a string", async () => {
      taskBody.description = 1;

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if task's priority is not in (high, medium, low)", async () => {
      taskBody.priority = "priority1";

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if task's priority is not a string", async () => {
      taskBody.priority = 1;

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if task's completed is not a boolean type", async () => {
      taskBody.completed = "completed1";

      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 400 if the number of tasks has reached the limit for the day", async () => {
      await exec();
      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should save the task if it is valid", async () => {
      const res = await exec();
      const task = await Task.find({ _id: res.body._id });

      expect(task).not.toBeNull();
      expect(res.status).toBe(httpStatus.CREATED);
    });

    it("should return the task if it is valid", async () => {
      const res = await exec();

      expect(res.body).toHaveProperty("_id");
      expect(res.body).toMatchObject(taskBody);
      expect(res.status).toBe(httpStatus.CREATED);
    });
  });

  describe("GET /", () => {
    let token;
    let _id;

    const exec = async () => {
      return await request(server)
        .get("/api/v1/tasks")
        .set("authorization", `Bearer ${token}`);
    };

    beforeEach(() => {
      _id = mongoose.Types.ObjectId();

      token = new User({
        _id,
        maxTask: 10,
      }).generateAuthToken();
    });

    it("should return 401 if client is not logged in", async () => {
      token = "";

      const res = await exec();

      expect(res.status).toBe(httpStatus.UNAUTHORIZED);
    });

    it("should return all tasks", async () => {
      const tasks = [
        {
          title: "title1",
          description: "description1",
          priority: priorityTypes.HIGH,
          completed: false,
          createdBy: _id,
        },
        {
          title: "title2",
          description: "description2",
          priority: priorityTypes.MEDIUM,
          completed: true,
          createdBy: _id,
        },
      ];
      await Task.collection.insertMany(tasks);

      const res = await exec();

      expect(res.status).toBe(httpStatus.OK);
      expect(res.body.length).toBe(2);
      expect(res.body.some((g) => g.title === "title1")).toBeTruthy();
      expect(res.body.some((g) => g.title === "title2")).toBeTruthy();
    });
  });

  describe("GET /:id", () => {
    let token;
    let _id;

    const exec = async () => {
      return await request(server)
        .get(`/api/v1/tasks/${_id}`)
        .set("authorization", `Bearer ${token}`);
    };

    beforeEach(() => {
      _id = mongoose.Types.ObjectId();

      token = new User({
        _id,
        maxTask: 10,
      }).generateAuthToken();
    });

    it("should return 401 if client is not logged in", async () => {
      token = "";

      const res = await exec();

      expect(res.status).toBe(httpStatus.UNAUTHORIZED);
    });

    it("should return 400 if invalid id is passed", async () => {
      _id = 1;
      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 404 if no task with the given id exists", async () => {
      const res = await exec();

      expect(res.status).toBe(httpStatus.NOT_FOUND);
    });

    it("should return a task if valid id is passed", async () => {
      const task = new Task({
        title: "title1",
        description: "description1",
        priority: priorityTypes.HIGH,
        completed: false,
        createdBy: _id,
      });
      await task.save();
      _id = task._id;

      const res = await exec();

      expect(res.status).toBe(httpStatus.OK);
      expect(res.body).toHaveProperty("title", "title1");
    });
  });

  describe("PATCH /:id", () => {
    let token;
    let _id;
    let taskBody;

    const exec = async () => {
      return await request(server)
        .patch(`/api/v1/tasks/${_id}`)
        .set("authorization", `Bearer ${token}`)
        .send(taskBody);
    };

    beforeEach(() => {
      _id = mongoose.Types.ObjectId();

      token = new User({
        _id,
        maxTask: 10,
      }).generateAuthToken();

      taskBody = {
        title: "title1 updating",
        description: "description1 updating",
        priority: priorityTypes.LOW,
        completed: true,
      };
    });

    it("should return 401 if client is not logged in", async () => {
      token = "";

      const res = await exec();

      expect(res.status).toBe(httpStatus.UNAUTHORIZED);
    });

    it("should return 400 if invalid id is passed", async () => {
      _id = 1;
      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 404 if no task with the given id exists", async () => {
      const res = await exec();

      expect(res.status).toBe(httpStatus.NOT_FOUND);
    });

    it("should return 403 if client access to a task that don't have permission", async () => {
      const task = new Task({
        title: "title1",
        description: "description1",
        priority: priorityTypes.HIGH,
        completed: false,
        createdBy: mongoose.Types.ObjectId(),
      });
      await task.save();
      _id = task._id;

      const res = await exec();

      expect(res.status).toBe(httpStatus.FORBIDDEN);
    });

    it("should save the task if it is valid", async () => {
      const task = new Task({
        title: "title1",
        description: "description1",
        priority: priorityTypes.HIGH,
        completed: false,
        createdBy: _id,
      });
      await task.save();
      _id = task._id;

      const res = await exec();
      const updated = await Task.find({ _id });

      expect(updated).not.toBeNull();
      expect(res.status).toBe(httpStatus.OK);
    });

    it("should return an updated task if valid id is passed", async () => {
      const task = new Task({
        title: "title1",
        description: "description1",
        priority: priorityTypes.HIGH,
        completed: false,
        createdBy: _id,
      });
      await task.save();
      _id = task._id;

      const res = await exec();

      expect(res.status).toBe(httpStatus.OK);
      expect(res.body).toMatchObject(taskBody);
    });
  });

  describe("DELETE /:id", () => {
    let token;
    let _id;
    let taskBody;

    const exec = async () => {
      return await request(server)
        .delete(`/api/v1/tasks/${_id}`)
        .set("authorization", `Bearer ${token}`)
        .send();
    };

    beforeEach(() => {
      _id = mongoose.Types.ObjectId();

      token = new User({
        _id,
        maxTask: 10,
      }).generateAuthToken();
    });

    it("should return 401 if client is not logged in", async () => {
      token = "";

      const res = await exec();

      expect(res.status).toBe(httpStatus.UNAUTHORIZED);
    });

    it("should return 400 if invalid id is passed", async () => {
      _id = 1;
      const res = await exec();

      expect(res.status).toBe(httpStatus.BAD_REQUEST);
    });

    it("should return 404 if no task with the given id exists", async () => {
      const res = await exec();

      expect(res.status).toBe(httpStatus.NOT_FOUND);
    });

    it("should return 403 if client access to a task that don't have permission", async () => {
      const task = new Task({
        title: "title1",
        description: "description1",
        priority: priorityTypes.HIGH,
        completed: false,
        createdBy: mongoose.Types.ObjectId(),
      });
      await task.save();
      _id = task._id;

      const res = await exec();

      expect(res.status).toBe(httpStatus.FORBIDDEN);
    });

    it("should return 200 if valid id is passed", async () => {
      const task = new Task({
        title: "title1",
        description: "description1",
        priority: priorityTypes.HIGH,
        completed: false,
        createdBy: _id,
      });
      await task.save();
      _id = task._id;

      const res = await exec();
      const deletedTask = await Task.findById(_id);

      expect(res.status).toBe(httpStatus.NO_CONTENT);
      expect(deletedTask).toBeNull();
    });
  });
});
