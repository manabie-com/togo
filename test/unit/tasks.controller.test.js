const { expect } = require("chai");
const sinon = require("sinon");
const taskServices = require("../../src/services/tasks.service");
const userServices = require("../../src/services/users.service");
const { create, get, remove, update } = require("../../src/controllers/tasks.controller");

describe("Test tasks.controller.js", () => {
  let req, res, next;

  beforeEach(() => {
    req = {
      query: {}
    };

  });

  afterEach(() => {
    sinon.restore();
  });

  describe("Test function get", () => {
    it("Should return all tasks", async () => {
      res = {
        json: sinon.stub().returnsThis()
      };

      const fakeTasks = [
        {
          _id: "62977ce3ef67a59df94feeb0",
          name: "Test 2",
          user: "62976842fdaba279ef64966a",
          createdAt: "2022-06-01T14:51:15.861Z",
          updatedAt: "2022-06-01T14:51:15.861Z",
        },
        {
          _id: "629859c52d79b6f2e4d5df28",
          name: "Test 3",
          user: "62976842fdaba279ef64966a",
          createdAt: "2022-06-02T06:33:41.837Z",
          updatedAt: "2022-06-02T06:33:41.837Z",
        }
      ];
      sinon.stub(taskServices, "getAll").returns(fakeTasks);
      await get(req, res, next);
      expect(res.json.calledWith(fakeTasks)).to.equal(true);
    });
    it("Should return empty array when do not have any task", async () => {
      res = {
        json: sinon.stub().returnsThis()
      };

      const fakeTasks = [];

      sinon.stub(taskServices, "getAll").returns(fakeTasks);
      await get(req, res, next);
      expect(res.json.calledWith(fakeTasks)).to.equal(true);
    });
  });
  describe("Test function create", () => {
    it("Should return info when created", async () => {
      res = {
        json: sinon.stub().returnsThis()
      };

      req.body = {
        _id: 1,
        username: "admin",
        password: "admin",
      };

      const createdTask = {
        _id: "629859c52d79b6f2e4d5df28",
        name: "Task Test 3",
        user: "62976842fdaba279ef64966a",
        createdAt: "2022-06-02T06:33:41.837Z",
        updatedAt: "2022-06-02T06:33:41.837Z",
      };
      sinon.stub(taskServices, "create").returns(createdTask);
      sinon.stub(taskServices, "getAllTasksTodayByUser").returns([]);
      sinon.stub(userServices, "getUserByName").returns({ _id: 1, limit: 1 });
      await create(req, res, next);
      expect(res.json.calledWith(createdTask)).to.equal(true);
    });
    it("Should throw error when reached limit of tasks per day", async () => {
      next = (err) => {
        expect(err.message).to.equal("You have reached your limit of tasks per day");
      };


      req.body = {
        _id: 1,
        username: "admin",
        password: "admin",
      };

      const createdTasks = [{
        _id: "629859c52d79b6f2e4d5df28",
        name: "Task Test 3",
        user: "62976842fdaba279ef64966a",
        createdAt: "2022-06-02T06:33:41.837Z",
        updatedAt: "2022-06-02T06:33:41.837Z",
      }];
      sinon.stub(taskServices, "getAllTasksTodayByUser").returns([createdTasks]);
      sinon.stub(userServices, "getUserByName").returns({ _id: 1, limit: 1 });
      await create(req, res, next);
    });
  });
  describe("Test function update", () => {
    it("Should return task info when updated", async () => {
      res = {
        json: sinon.stub().returnsThis()
      };

      req = {
        params: {
          id: 1
        },
        body: {
          name: "Test 3",
        }
      };

      const updatedTask = {
        _id: "1",
        name: "Test 3",
        user: "62976842fdaba279ef64966a",
        createdAt: "2022-06-02T06:33:41.837Z",
        updatedAt: "2022-06-02T06:33:41.837Z",
      };
      sinon.stub(taskServices, "update").returns(updatedTask);
      await update(req, res, next);
      expect(res.json.calledWith(updatedTask)).to.equal(true);
    });
  });
  describe("Test function delete", () => {
    it("Should return user info when deleted", async () => {
      res = {
        json: sinon.stub().returnsThis()
      };

      req = {
        params: {
          id: 1
        },
      };

      const deletedTask = {
        _id: "1",
        name: "Test 3",
        user: "62976842fdaba279ef64966a",
        createdAt: "2022-06-02T06:33:41.837Z",
        updatedAt: "2022-06-02T06:33:41.837Z",
      };
      sinon.stub(taskServices, "remove").returns(deletedTask);
      await remove(req, res, next);
      expect(res.json.calledWith(deletedTask)).to.equal(true);
    });
  });
});