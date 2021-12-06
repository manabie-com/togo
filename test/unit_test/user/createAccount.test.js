const userModel = require("../../../src/models/user.model");
const { createAccount } = require("../../../src/services/user.service");

describe("note: createAccount", () => {
  it("check create user", async () => {
    const userName = "Anh Phan";
    const password = "1234567";
    const passwordHashed = "ad4w646d8aw48was3232135aw484";
    userModel.create = jest.fn().mockResolvedValue({
      userName: userName,
      password: passwordHashed,
      createdAt: "2021-12-03T16:38:58.158Z",
      updatedAt: "2021-12-03T16:38:58.158Z",
      __v: 0,
    });

    const newTask = await createAccount(userName, password);
    expect(newTask).toStrictEqual({
      userName: userName,
      password: passwordHashed,
      createdAt: "2021-12-03T16:38:58.158Z",
      updatedAt: "2021-12-03T16:38:58.158Z",
      __v: 0,
    });
  });

  it("check create user exist", async () => {
    try {
      const userName = "Anh Phan";
      const password = "1234567";

      await createAccount(userName, password);
    } catch (err) {
      expect(err).toThrow(TypeError);
    }
    
  });
})