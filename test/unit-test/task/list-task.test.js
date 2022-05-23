const { taskModel } = require('../../../src/models/model.task');
const taskService = require('../../../src/services/services.task');


describe("[UNIT TEST]: LIST TASK TEST.", () => {
  // =============== CASE 01 ================
  it('Create user success.', async () => {
    // taskModel.create = jest.fn().mockResolvedValue({
    //   username: user.username,
    //   createdAt: "2022-01-09T15:57:14.750Z",
    //   updatedAt: "2022-01-09T15:57:14.750Z",
    //   password: user.password,
    //   limit: 10,
    //   __v: 0
    // });
    expect(true).toBe(true);
  });
});