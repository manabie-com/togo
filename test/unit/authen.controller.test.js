const { expect } = require("chai");
const sinon = require("sinon");
const User = require("../../src/models/users.model");
const authenConfig = require("../../src/configs/authen.config");
const { encrypt } = require("../../src/utils/helper.util");
const { getToken, isValidToken } = require("../../src/controllers/authen.controller");

describe("Test authen.controller.js", () => {
  afterEach(() => {
    sinon.restore();
  });

  describe("Test function getToken", () => {
    it("Should return token", async () => {
      const username = "admin";
      const password = "admin";
      const fakeUser = {
        _id: 1,
        username: "admin",
        password: encrypt("admin"),
      };
      sinon.stub(User, "findOne").returns({ lean: () => fakeUser });
      const token = await getToken(username, password);
      expect(token).to.be.ok;
      expect(token).to.have.property("user");
      expect(token).to.have.property("token");
      expect(token.token).to.include("JWT");
    });
  });
  describe("Test function isValidToken", () => {
    it("Should return true when token is valid", async () => {
      const username = "admin";
      const password = "admin";
      const fakeUser = {
        _id: 1,
        username: "admin",
        password: encrypt("admin"),
      };
      sinon.stub(User, "findOne").returns({ lean: () => fakeUser });
      const token = await getToken(username, password);
      const headers = {
        authorization: token.token,
      };
      const tokenString = await isValidToken(headers);
      expect(tokenString).to.be.ok;
    });
    it("Should throw an error when token is invalid", async () => {
      const headers = {
        authorization: "JWT-eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiaWF0IjoxNTU0NjIzNjQ5LCJleHAiOjE1NTQ2MjYwNDl9.3y_7Vu-f6-_Xr_Z5f5G5z5_xJHW8lF-_YQ-I-1lQ2c"
      };
      try {
        await isValidToken(headers);
      } catch (error) {
        expect(error.message).to.be.equal("invalid signature");
      }
    });
    it("Should throw an error when token is expired", async () => {
      const clock = sinon.useFakeTimers();
      try {
        const username = "admin";
        const password = "admin";
        const fakeUser = {
          _id: 1,
          username: "admin",
          password: encrypt("admin"),
        };
        sinon.stub(User, "findOne").returns({ lean: () => fakeUser });
        const token = await getToken(username, password);
        const headers = {
          authorization: token.token,
        };
        clock.tick(authenConfig.jwt.expiresInMinutes * 60 * 1000);
        await isValidToken(headers);
      } catch (error) {
        expect(error.message).to.be.equal("jwt expired");
      }
    });
    it("Should throw an error when missing token in header", async () => {
      try {
        const username = "admin";
        const password = "admin";
        const fakeUser = {
          _id: 1,
          username: "admin",
          password: encrypt("admin"),
        };
        sinon.stub(User, "findOne").returns({ lean: () => fakeUser });
        const token = await getToken(username, password);
        const headers = {
          authorization: "",
        };
        await isValidToken(headers);
      } catch (error) {
        expect(error.message).to.be.equal("Missing Authorization");
      }
    });
    it("Should throw an error when wrong password", async () => {
      try {
        const username = "admin";
        const password = "admin1";
        const fakeUser = {
          _id: 1,
          username: "admin",
          password: encrypt("admin"),
        };
        sinon.stub(User, "findOne").returns({ lean: () => fakeUser });
        const token = await getToken(username, password);
        const headers = {
          authorization: token.token,
        };
        await isValidToken(headers);
      } catch (error) {
        expect(error.message).to.be.equal("Invalid username or password");
      }
    });
  });
});