const noteModel = require("../../../src/models/note.model");
const { updateTask } = require('../../../src/services/note.service');


describe("note: updateTask", () => {

  it("test update task with taskID and userID correct reference", async () => {
    const content = "test add task";
    const id = "61a87e156737ae9d27b74ba";
    const userId = "as7d87a87da67w6d72h";
    const ticked = true;

    noteModel.findOneAndUpdate = jest.fn().mockResolvedValue({
      _id: id,
      content: content,
      user: userId,
      ticked: ticked,
      createdAt: "2021-12-03T16:38:58.158Z",
      updatedAt: "2021-12-03T16:38:58.158Z",
      __v: 0,
    });

    const taskUpdated = await updateTask(id, userId, content, ticked);
    expect(taskUpdated).toStrictEqual({
      _id: id,
      content: content,
      user: userId,
      ticked: ticked,
      createdAt: "2021-12-03T16:38:58.158Z",
      updatedAt: "2021-12-03T16:38:58.158Z",
      __v: 0,
    });
  })

  it("test update task with taskID and userID incorrect reference", async () => {
    try {
      const content = "test add task";
      const id = "61a87e156737ae9d27b74ba";
      const userId = "as7d87a87da67w6d72h";
      const ticked = true;

      await updateTask(id, userId, content, ticked);

    } catch (err) {
      expect(err).toThrow(TypeError);
    }
  });

  it("test update task with task was deleted", async () => {
    const content = "test add task";
    const id = "61a87e156737ae9d27b74ba";
    const userId = "as7d87a87da67w6d72h";
    const ticked = true;

    noteModel.findOneAndUpdate = jest.fn().mockResolvedValue();

    const taskUpdated = await updateTask(id, userId, content, ticked);
    expect(taskUpdated).toStrictEqual();
  });
})
