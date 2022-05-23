const { taskModel } = require('../../../src/models/model.task');
const taskService = require('../../../src/services/services.task');

const templateTask = {
  "_id": "61db05da8b605b8dc2524391",
  "userId": "61db05da8b605b8dc2524392",
  "title": "Job 001",
  "description": "Description for job 001",
  "status": 0,
  "createdAt": "2022-05-23T14:52:08.322Z",
  "updatedAt": "2022-05-23T14:52:08.322Z",
  "__v": 0
};
taskModel.findById = jest.fn().mockResolvedValue(templateTask);

describe("[UNIT TEST]: FIND TASK TEST.", () => {
  // =============== CASE 01 ================
  it('Read user task.', async () => {
    const res = await taskService.readTask('61db05da8b605b8dc2524391');
    expect(res).toStrictEqual(templateTask);
  });

  // =============== CASE 02 ================
  it('Read task with id not exists.', async () => {
    taskModel.findById = jest.fn().mockResolvedValue(null);

    const res = await taskService.readTask('61db05da8b605b8dc2524392');
    expect(res).toStrictEqual(null);
  });

  // =============== CASE 03 ================
  it('Read task with id = null.', async () => {
    taskModel.findById = jest.fn().mockResolvedValue(null);

    const res = await taskService.readTask(null);
    expect(res).toStrictEqual(null);
  });
});