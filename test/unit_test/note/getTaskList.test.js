const noteModel = require("../../../src/models/note.model");
const{ getTaskList } = require("../../../src/services/note.service")
const mongoose = require('mongoose')

describe("note: getTaskList", () => {
  it("get task list in day with correct userId reference", async () => {
    const userId = mongoose.Types.ObjectId("61a87e156737ae9d27b74baa");
    const day = "2021-12-03";

    noteModel.find = jest.fn().mockResolvedValue([
      {
        content: "test task 1",
        user: userId,
        ticked: true,
        deleted: false,
        createdAt: "2021-12-03T14:38:58.158Z",
        updatedAt: "2021-12-03T16:38:58.158Z",
        __v: 0,
      },
      {
        content: "test task 2",
        user: userId,
        ticked: true,
        deleted: false,
        createdAt: "2021-12-03T16:38:58.158Z",
        updatedAt: "2021-12-03T16:38:58.158Z",
        __v: 0,
      },
    ]);

    const getTask = await getTaskList(userId, day);
    expect(getTask).toStrictEqual([
      {
        content: "test task 1",
        user: userId,
        ticked: true,
        deleted: false,
        createdAt: "2021-12-03T14:38:58.158Z",
        updatedAt: "2021-12-03T16:38:58.158Z",
        __v: 0,
      },
      {
        content: "test task 2",
        user: userId,
        ticked: true,
        deleted: false,
        createdAt: "2021-12-03T16:38:58.158Z",
        updatedAt: "2021-12-03T16:38:58.158Z",
        __v: 0,
      },
    ]);
  });

  it("get task list in day with incorrect userId reference", async () => {
    try{
      const userId = "61a87e156737ae9d27b7";
      const day = "2021-12-03";

      await getTaskList(userId, day);
    } catch (err) {
      expect(err).toThrow(TypeError);
    }
    
  });
});
