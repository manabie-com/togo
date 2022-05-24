const { userModel } = require('../../../src/models/model.user');
const userService = require('../../../src/services/services.user');


describe("[UNIT TEST]: CREATE USER TEST.", () => {
  // =============== CASE 01 ================
  it('Create user success.', async () => {
    const user = {
      username: 'tiennm',
      password: 'hashed_password'
    }

    userModel.create = jest.fn().mockResolvedValue({
      username: user.username,
      createdAt: "2022-01-09T15:57:14.750Z",
      updatedAt: "2022-01-09T15:57:14.750Z",
      password: user.password,
      limit: 10,
      __v: 0
    });

    const result = await userService.createUser(user);
    delete result?.data?.password; // Not knowing the hash value in advance

    expect(result.data).toStrictEqual({
      username: user.username,
      createdAt: "2022-01-09T15:57:14.750Z",
      updatedAt: "2022-01-09T15:57:14.750Z",
      limit: 10,
      __v: 0
    });
  });

  // =============== CASE 02 ================
  it('Create with username is a number.', async () => {
    const user = {
      username: 123456,
      password: 'hashed_password'
    }
    const result = await userService.createUser(user);
    delete result?.data?.password; // Not knowing the hash value in advance
    expect(result.message).toBe('"username" must be a string');
  });

  // =============== CASE 03 ================
  it('Create with password is a number.', async () => {
    const user = {
      username: 'username',
      password: 123456
    }
    const result = await userService.createUser(user);
    expect(result.message).toBe('"password" must be a string');
  });

  // =============== CASE 04 ================
  it('Create with username is blank.', async () => {
    const user = {
      username: '',
      password: '123456'
    }
    const result = await userService.createUser(user);
    expect(result.message).toBe('"username" is not allowed to be empty');
  });

  // =============== CASE 05 ================
  it('Create with password is blank.', async () => {
    const user = {
      username: 'username',
      password: ''
    }
    const result = await userService.createUser(user);
    expect(result.message).toBe('"password" is not allowed to be empty');
  });
});