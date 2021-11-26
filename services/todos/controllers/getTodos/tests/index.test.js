const { getTodos } = require("..");
const todos = require("../../../../../models/todos/todos.json");

describe("Test get todos", () => {
  it("should get todos successful with true url", () => {
    const modelUrl = "models/todos/todos.json";
    const actual = getTodos(modelUrl);
    expect(actual.length).toEqual(todos.length);
    expect(actual).toEqual(todos);
  });

  it("should get todos failed with false url", () => {
    const modelUrl = "model/todos/todos.json";
    const actual = getTodos(modelUrl);
    expect(actual.code).toEqual("ENOENT");
    expect(actual.errno).toEqual(-4058);
  });
});
