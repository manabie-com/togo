const { expect } = require("chai");
const sinon = require("sinon");
const userServices = require("../../src/services/users.service");
const { create, get, remove, update } = require("../../src/controllers/users.controller");
const req = require("express/lib/request");

describe("Test users.controller.js", () => {
  let req, res, next;

  beforeEach(() => {
    req = {
      query: {}
    };

    next = () => { };
  });

  afterEach(() => {
    sinon.restore();
  });

  describe("Test function get", () => {
    it("Should return all users", async () => {
      res = {
        json: sinon.stub().returnsThis()
      };

      const fakeUsers = [
        {
          _id: "1",
          username: "admin",
          tasks: [],
          limit: 1
        },
        {
          _id: "2",
          username: "admin2",
          tasks: [],
          limit: 1
        }
      ];
      sinon.stub(userServices, "getAll").returns(fakeUsers);
      await get(req, res, next);
      expect(res.json.calledWith(fakeUsers)).to.equal(true);
    });
    it("Should return empty array when do not have any user", async () => {
      res = {
        json: sinon.stub().returnsThis()
      };

      const fakeUsers = [];

      sinon.stub(userServices, "getAll").returns(fakeUsers);
      await get(req, res, next);
      expect(res.json.calledWith(fakeUsers)).to.equal(true);
    });
  });
  describe("Test function create", () => {
    it("Should return user info when created", async () => {
      res = {
        json: sinon.stub().returnsThis()
      };

      req.body = {
        _id: 1,
        username: "admin",
        password: "admin",
      };

      const createdUser = {
        _id: 1,
        username: "admin",
        tasks: [],
        limit: 0
      };
      sinon.stub(userServices, "create").returns(createdUser);
      await create(req, res, next);
      expect(res.json.calledWith(createdUser)).to.equal(true);
    });
  });
  describe("Test function update", () => {
    it("Should return user info when updated", async () => {
      res = {
        json: sinon.stub().returnsThis()
      };

      req = {
        params: {
          id: 1
        },
        body: {
          limit: 1
        },
      };

      const updatedUser = {
        _id: 1,
        username: "admin",
        tasks: [],
        limit: 1
      };
      sinon.stub(userServices, "update").returns(updatedUser);
      await update(req, res, next);
      expect(res.json.calledWith(updatedUser)).to.equal(true);
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

      const deletedUser = {
        _id: 1,
        username: "admin",
        tasks: [],
        limit: 1
      };
      sinon.stub(userServices, "remove").returns(deletedUser);
      await remove(req, res, next);
      expect(res.json.calledWith(deletedUser)).to.equal(true);
    });
  });
});