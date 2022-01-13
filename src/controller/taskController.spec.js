const { addTask } = require("./taskController");

// Tests the addTask method using mocked data and methods
describe("addTask", () => {
  it("should add one task", () => {
    const mockData = { users: [], tasks: [] };
    const mockReqData = { user_name: "user1", task_name: "task1" };

    const mockAddTask = jest.fn((task) => {
      mockData.tasks.push({ ...task });
      return null;
    });

    const mockRepo = {
      addTask: mockAddTask,
    };

    const addTaskHandler = addTask(mockRepo, new Date().toISOString());

    addTaskHandler(
      { body: mockReqData },
      { json: jest.fn(), status: jest.fn(), send: jest.fn() },
      jest.fn()
    );

    expect(mockAddTask).toBeCalledTimes(1);
    expect(mockData.tasks.length).toEqual(1);
    expect(mockData.tasks[0]).toMatchObject(mockReqData);
  });
});
