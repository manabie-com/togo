const { User } = require("../../../../apis/models");
const { userService } = require("../../../../apis/services");

describe("getUserByEmail", () => {
  it("should return user by email", async () => {
    const email = "a@manabie.com";
    const user = {
      email,
    };
    User.findOne = jest.fn().mockResolvedValue(user);
    const result = userService.getUserByEmail(email);

    await expect(result).resolves.toMatchObject(user);
    expect(User.findOne).toBeCalledTimes(1);
  });
});

describe("createUser", () => {
  it("should throw if email is taken", async () => {
    const email = "a@manabie.com";
    const userBody = {
      email,
    };

    User.isEmailTaken = jest.fn().mockResolvedValue(true);
    const result = userService.createUser(userBody);

    await expect(result).rejects.toThrowError(/taken/);
    expect(User.isEmailTaken).toBeCalledTimes(1);
  });

  it("should create user", async () => {
    const email = "a@manabie.com";
    const userBody = {
      email,
    };

    User.isEmailTaken = jest.fn().mockResolvedValue(false);
    User.create = jest.fn().mockResolvedValue(userBody);
    const result = userService.createUser(userBody);

    await expect(result).resolves.toMatchObject(userBody);
    expect(User.isEmailTaken).toBeCalledTimes(1);
  });
});
