const jwt = require("jsonwebtoken");
const mongoose = require("mongoose");
const bcrypt = require("bcryptjs");

const { User, preSaveFunc } = require("../../../../apis/models");
const env = require("../../../../configs/env");

describe("user.generateAuthToken", () => {
  it("should return a valid JWT", () => {
    const payload = {
      _id: mongoose.Types.ObjectId().toHexString(),
      maxTask: 10,
    };
    const user = new User(payload);
    const token = user.generateAuthToken();
    const decoded = jwt.verify(token, env.jwt.secret);
    expect(decoded).toMatchObject(payload);
  });
});

describe("user.isPasswordMatch", () => {
  it("should return true if the password is match", async () => {
    const user = new User({ password: await bcrypt.hash("password1", 10) });

    const matchResult = await user.isPasswordMatch("password1");
    expect(matchResult).toBe(true);
  });

  it("should return false if the password is match", async () => {
    const user = new User({ password: await bcrypt.hash("password1", 10) });

    const nonMatchResult = await user.isPasswordMatch("password2");
    expect(nonMatchResult).toBe(false);
  });
});

describe("User.isEmailTaken", () => {
  it("should return true if the email is taken", async () => {
    User.findOne = jest.fn().mockReturnValue({ email: "a" });
    const takenResult = await User.isEmailTaken("a");
    expect(takenResult).toBe(true);
  });

  it("should return false if the email is not taken", async () => {
    User.findOne = jest.fn().mockReturnValue(null);
    const notTakenResult = await User.isEmailTaken("a");
    expect(notTakenResult).toBe(false);
  });
});

describe("preSaveFunc", () => {
  it("should not change password when password is not modified", async () => {
    const next = jest.fn();
    const user = {
      isModified: jest.fn().mockReturnValueOnce(false),
      password: "password1",
    };

    await preSaveFunc(next, user);
    expect(user.isModified).toBeCalledWith("password");
    expect(next).toBeCalledTimes(1);
    expect(user.password).toBe("password1");
  });

  it("should hash password when password is modified", async () => {
    const next = jest.fn();
    const user = {
      isModified: jest.fn().mockReturnValueOnce(true),
      password: "password1",
    };

    await preSaveFunc(next, user);
    expect(user.isModified).toBeCalledWith("password");
    expect(next).toBeCalledTimes(1);
    expect(user.password).not.toBe("password1");
  });
});
