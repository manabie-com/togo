const noteModel = require("../../../src/models/note.model");
const { deleteTask } = require("../../../src/services/note.service");

describe("note: deleteTask", () => {
  it("test delete task with taskID and userID correct reference", async () => {
    const content = "test add task";
    const id = "61a87e156737ae9d27b74ba";
    const userId = "as7d87a87da67w6d72h";
    const ticked = true;

    noteModel.findOneAndUpdate = jest.fn().mockResolvedValue({
      _id: id,
      content: content,
      user: userId,
      ticked: ticked,
      deleted: true,
      createdAt: "2021-12-03T16:38:58.158Z",
      updatedAt: "2021-12-03T16:38:58.158Z",
      __v: 0,
    });

    const taskDeleted = await deleteTask(id, userId);
    expect(taskDeleted).toStrictEqual({
      _id: id,
      content: content,
      user: userId,
      ticked: ticked,
      deleted: true,
      createdAt: "2021-12-03T16:38:58.158Z",
      updatedAt: "2021-12-03T16:38:58.158Z",
      __v: 0,
    });
  });

  it("test delete task with taskID and userID incorrect reference", async () => {
    try {

      const id = "61a87e156737ae9d27b74ba";
      const userId = "as7d87a87da67w6d72h";

      await deleteTask(id, userId);
    } catch (err) {
      expect(err).toThrow(TypeError);
    }
  });

  it("test delete task with task was deleted", async () => {
    const id = "61a87e156737ae9d27b74ba";
    const userId = "as7d87a87da67w6d72h";

    noteModel.findOneAndUpdate = jest.fn().mockResolvedValue();
    const taskUpdated = await deleteTask(id, userId);
    expect(taskUpdated).toStrictEqual();
  });
});
