const { taskModel } = require('../../../src/models/model.task');
const taskService = require('../../../src/services/services.task');

const templateTask = {
  "_id": "61db05da8b605b8dc2524392",
  "userId": "61db05da8b605b8dc2524392",
  "title": "New title",
  "description": "Description for job 001",
  "status": 0,
  "createdAt": "2022-05-23T14:52:08.322Z",
  "updatedAt": "2022-05-23T14:52:08.322Z",
  "__v": 0
};
taskModel.findByIdAndUpdate = jest.fn().mockResolvedValue(templateTask);

describe("[UNIT TEST]: UPDATE TASK TEST.", () => {
  // =============== CASE 01 ================
  it('Update task success.', async () => {
    const res = await taskService.updateTask(templateTask._id, {
      "userId": "61db05da8b605b8dc2524392",
      "title": "New title",
      "description": "Description for job 001",
      "status": 0
    });
    expect(res.success).toBe(true);
  });

  // =============== CASE 02 ================
  it('Update task not exists.', async () => {
    taskModel.findByIdAndUpdate = jest.fn().mockResolvedValue(null);

    const res = await taskService.updateTask('61db05da8b605b8dc2524391', {
      "userId": "61db05da8b605b8dc2524392",
      "title": "New title",
      "description": "Description for job 001",
      "status": 0
    });
    expect(res.data).toBe(null);
  });

  // =============== CASE 02 ================
  it('Update task with empty userId.', async () => {
    const res = await taskService.updateTask(templateTask._id, {
      "userId": "",
      "title": "New title",
      "description": "Description for job 001",
      "status": 0
    });
    expect(res.code).toBe(400);
    expect(res.message).toBe('"userId" is not allowed to be empty');
  });

  // =============== CASE 03 ================
  it('Update task with empty description.', async () => {
    const res = await taskService.updateTask(templateTask._id, {
      "userId": "61db05da8b605b8dc2524392",
      "title": "New title",
      "description": "",
      "status": 0
    });
    expect(res.code).toBe(400);
    expect(res.message).toBe('"description" is not allowed to be empty');
  });

  // =============== CASE 04 ================
  it('Update task with status is not a number.', async () => {
    const res = await taskService.updateTask(templateTask._id, {
      "userId": "61db05da8b605b8dc2524392",
      "title": "New title",
      "description": "new description",
      "status": "acbd"
    });
    expect(res.code).toBe(400);
    expect(res.message).toBe('"status" must be a number');
  });
});