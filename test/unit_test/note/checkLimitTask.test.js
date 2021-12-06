const noteModel = require("../../../src/models/note.model");
const { checkLimitTask } = require("../../../src/services/note.service");

describe("note: checkLimitTask", () => {
  it("test check task less than limit task numner ", async () => {
    const content = "test add task";
    const id = "61a87e156737ae9d27b74ba";
    noteModel.count = jest.fn().mockResolvedValue(3);
    const checkTask = await checkLimitTask(content, id);
    expect(checkTask).toStrictEqual(false);
  });

  it("test check task more than limit task numner ", async () => {
    const content = "test add task";
    const id = "61a87e156737ae9d27b74ba";
    noteModel.count = jest.fn().mockResolvedValue(6);
    const checkTask = await checkLimitTask(content, id);
    expect(checkTask).toStrictEqual(true);
  });

  // it("test add task with user")
});
