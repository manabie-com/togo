const noteModel = require("../../../src/models/note.model");
const { addTask2DB } = require('../../../src/services/note.service');


describe("note: addTask", () => {


  it("test add task with userID correct reference", async () => {
    const content = "test add task";
    const id = "61a87e156737ae9d27b74ba";
    noteModel.create = jest.fn().mockResolvedValue({
      content: content,
      user: id,
      ticked: false,
      createdAt: "2021-12-03T16:38:58.158Z",
      updatedAt: "2021-12-03T16:38:58.158Z",
      __v: 0,
    });

    const newTask = await addTask2DB(content, id);
    expect(newTask).toStrictEqual({
      content: content,
      user: id,
      ticked: false,
      createdAt: "2021-12-03T16:38:58.158Z",
      updatedAt: "2021-12-03T16:38:58.158Z",
      __v: 0,
    });
  });


  // it("test add task with user")
})
