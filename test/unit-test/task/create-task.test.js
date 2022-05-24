const { taskModel } = require('../../../src/models/model.task');
const taskService = require('../../../src/services/services.task');
const templateTask = {
  "userId": "61db05da8b605b8dc2524392",
  "title": "Job 001",
  "description": "Description for job 001",
  "status": 0,
  "createdAt": "2022-05-23T14:52:08.322Z",
  "updatedAt": "2022-05-23T14:52:08.322Z",
  "__v": 0
};
taskModel.create = jest.fn().mockResolvedValue(templateTask);

describe("[UNIT TEST]: CREATE TASK TEST.", () => {
  // =============== CASE 01 ================
  it('Create task success.', async () => {
    const res = await taskService.createTask({
      "userId": "61db05da8b605b8dc2524392",
      "title": "Job 001",
      "description": "Description for job 001",
      "status": 0
    })
    expect(res.success).toBe(true);
    expect(res.data).toStrictEqual(templateTask);
  });

  // =============== CASE 02 ================
  it('Create task with title is a number.', async () => {
    const res = await taskService.createTask({
      "userId": "61db05da8b605b8dc2524392",
      "title": 123456,
      "description": "Description for job 001",
      "status": 0
    })
    expect(res.success).toBe(false);
    expect(res.message).toBe(`"title" must be a string`);
  });

  // =============== CASE 03 ================
  it('Create task with userId is a number.', async () => {
    const res = await taskService.createTask({
      "userId": 1234567,
      "title": "123456",
      "description": "Description for job 001",
      "status": 0
    })
    expect(res.success).toBe(false);
    expect(res.message).toBe(`"userId" must be a string`);
  });

  // =============== CASE 04 ================
  it('Create task with userId empty', async () => {
    const res = await taskService.createTask({
      "userId": "",
      "title": "123456",
      "description": "Description for job 001",
      "status": 0
    })
    expect(res.success).toBe(false);
    expect(res.message).toBe(`"userId" is not allowed to be empty`);
  });

  // =============== CASE 06 ================
  it('Create task with userId with length < 16', async () => {
    const res = await taskService.createTask({
      "userId": "abcd",
      "title": "123456",
      "description": "Description for job 001",
      "status": 0
    });

    expect(res.success).toBe(false);
    expect(res.message).toBe(`"userId" length must be at least 16 characters long`);
  });
});