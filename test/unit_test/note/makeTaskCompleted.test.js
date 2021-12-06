const noteModel = require("../../../src/models/note.model");
const { makeTaskCompleted } = require("../../../src/services/note.service");

describe("note: makeTaskCompleted", () => {
  it("test tick task with userID and task ID correct reference", async () => {
    const userId = "61a87e156737ae9d27b74baa";
    const idTask = "61a9ea4d9d00d91aef41c542";
    const contentTask = "test task 01";
    noteModel.findOneAndUpdate = jest.fn().mockResolvedValue({
      content: contentTask,
      user: userId,
      ticked: true,
      createdAt: "2021-12-03T16:38:58.158Z",
      updatedAt: "2021-12-03T16:38:58.158Z",
      __v: 0,
    });

    const updateTask = await makeTaskCompleted(idTask, userId);
    expect(updateTask).toStrictEqual({
      content: contentTask,
      user: userId,
      ticked: true,
      createdAt: "2021-12-03T16:38:58.158Z",
      updatedAt: "2021-12-03T16:38:58.158Z",
      __v: 0,
    });
  });

  it("test tick task with userID and task ID incorrect reference", async () => {
    try {
      const userId = "61a87e156737ae9d27b";
      const idTask = "61a9ea4d9d00d91aef4";

      await makeTaskCompleted(idTask, userId);
    } catch (err) {
      expect(err).toThrow("Opps, something went wrong");
    }
  });
});