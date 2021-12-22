const TaskModel = require('../../../src/models/task')
const {
    deleteTaskService
} = require('../../../src/services/task')

describe("Task: Delete task", () => {
  it("Delete task with taskID correct reference", async () => {
    const id= "61bef6594815ad48a1afe950";

    TaskModel.deleteOne = jest.fn().mockResolvedValue({
        deletedCount: 1
    });

    const taskDeleted = await deleteTaskService(id);
    expect(taskDeleted).toStrictEqual(true);
  });

  it("Delete task with taskID and userID incorrect reference", async () => {
    try {

      const id = "61bef6594815ad48a1afe950";
      await deleteTaskService(id);
    } catch (err) {
      expect(err).toThrow(TypeError);
    }
  });
});
