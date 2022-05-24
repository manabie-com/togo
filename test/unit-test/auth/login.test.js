const { login } = require('../../../src/services/services.user');
const { userModel } = require('../../../src/models/model.user');

userModel.find = jest.fn().mockResolvedValue([{
  username: 'user_002',
  password: '$2b$10$yYkJCQztrnxaraL/HZ2FFeCZpNwzm4WcghMawKQ7dblUOECknDHQu', // my_password
  limit: 10,
  createdAt: "2022-01-09T15:57:14.750Z",
  updatedAt: "2022-01-09T15:57:14.750Z",
  limit: 10,
  __v: 0
}]);

describe("[UNIT TEST]: LOGIN TEST.", () => {
  // =================== CASE 01 ===================
  it('Login with correct username and password', async () => {
    const res = await login('user_002', 'my_password');
    expect(res.success).toBe(true);
    expect(typeof (res.data) === 'string').toBe(true);
  });

  // =================== CASE 02 ===================
  it('Login with wrong username', async () => {
    const res = await login('user_002_wrong', 'my_password');
    expect(res.success).toBe(false);
    expect(typeof (res.data) === 'string').toBe(false);
  });

  // =================== CASE 03 ===================
  it('Login with wrong password', async () => {
    const res = await login('user_002', 'wrong_password');
    expect(res.success).toBe(false);
    expect(typeof (res.data) === 'string').toBe(false);
  });

});