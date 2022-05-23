const { taskModel } = require('../../../src/models/model.task');
const { userModel } = require('../../../src/models/model.user');
const taskService = require('../../../src/services/services.task');
const user = {
  _id: '61db05da8b605b8dc2524391',
  username: 'tiennm1',
  password: '123456',
  limit: 5
}

taskModel.count = jest.fn().mockResolvedValue(4);
userModel.findById = jest.fn().mockResolvedValue(user);

describe("[UNIT TEST]: COUNT TASK TEST.", () => {
  // =============== CASE 01 ================
  it('Count task per day/user.', async () => {
    const res = await taskService.count(user._id);
    expect(res).toBe(true);
  });
  // =============== CASE 02 ================
  it('Count task per day/user (over task).', async () => {
    taskModel.count = jest.fn().mockResolvedValue(5);
    const res = await taskService.count(user._id);
    expect(res).toBe(false);
  });
});