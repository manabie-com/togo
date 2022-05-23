const { taskModel } = require('../../../src/models/model.task');
const taskService = require('../../../src/services/services.task');
const templateTask = {
  "_id": "61db05da8b605b8dc2524392",
  "userId": "61db05da8b605b8dc2524392",
  "title": "Job 001",
  "description": "Description for job 001",
  "status": 0,
  "createdAt": "2022-05-23T14:52:08.322Z",
  "updatedAt": "2022-05-23T14:52:08.322Z",
  "__v": 0
};

taskModel.findOneAndUpdate = jest.fn().mockResolvedValue(templateTask);
taskModel.findById = jest.fn().mockResolvedValue(templateTask);

describe("[UNIT TEST]: DELETE TASK TEST.", () => {
  // =============== CASE 01 ================
  it('Delete task success.', async () => {
    const res = await taskService.deleteTask(templateTask._id);
    expect(res.success).toBe(true);
    expect(res.data).toStrictEqual(templateTask);
  });

  // =============== CASE 02 ================
  it('Delete task with null id.', async () => {
    const res = await taskService.deleteTask(null);
    expect(res.success).toBe(false);
    expect(res.code).toStrictEqual(0);
  });

  // =============== CASE 03 ================
  it('Delete task with id not exists.', async () => {
    taskModel.findById = jest.fn().mockResolvedValue(null);
    const res = await taskService.deleteTask('61db05da8b605b8dc2524391');
    expect(res.success).toBe(false);
    expect(res.code).toStrictEqual(404);
  });
});