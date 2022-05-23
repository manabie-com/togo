const { taskModel } = require('../../../src/models/model.task');
const taskService = require('../../../src/services/services.task');
const templateTask = {
  "userId": "61db05da8b605b8dc2524392",
  "title": "Job 001",
  "description": "Description for job 001",
  "status": 0,
  "createdAt": "2022-01-11T14:52:08.322Z",
  "updatedAt": "2022-01-11T14:52:08.322Z",
  "__v": 0
};
taskModel.findByIdAndUpdate = jest.fn().mockResolvedValue(templateTask);

describe("[UNIT TEST]: DELETE TASK TEST.", () => {
  // =============== CASE 01 ================
  it('DELETE task success.', async () => {
    
  });
});