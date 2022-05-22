const mongoose = require("mongoose");

const { taskService } = require("../../../../apis/services");
const { Task } = require("../../../../apis/models");
const { priorityTypes } = require("../../../../configs/priority");

describe("createTask", () => {
  it("should throw if the number of tasks has reached the limit for the day", () => {
    const user = {
      maxTask: 1,
    };
    taskService.countTasks = jest.fn().mockResolvedValue(1);

    expect(async () => {
      await taskService.createTask({}, user);
    }).rejects.toThrow();
  });

  it("should return created task info", () => {
    const taskBody = {
      title: "a",
    };

    const user = {
      _id: mongoose.Types.ObjectId().toHexString(),
      maxTask: 2,
    };

    taskService.countTasks = jest.fn().mockResolvedValue(1);
    Task.create = jest.fn().mockResolvedValue(taskBody);

    expect(
      Promise.resolve(taskService.createTask(taskBody, user))
    ).resolves.toMatchObject(taskBody);
  });
});

describe("getTasks", () => {
  it("should return list of task", () => {
    const tasks = [
      {
        title: "a",
        priority: priorityTypes.HIGH,
      },
      {
        title: "b",
        priority: priorityTypes.MEDIUM,
      },
    ];
    Task.find = jest.fn().mockReturnValue({
      sort: jest.fn().mockReturnValue({
        lean: jest.fn().mockResolvedValue(tasks),
      }),
    });

    expect(Promise.resolve(taskService.getTasks(""))).resolves.toMatchObject(
      tasks
    );
    expect(Task.find).toBeCalledTimes(1);
  });
});

describe("getTask", () => {
  it("should throw if can not find task by this id", () => {
    Task.findById = jest.fn().mockResolvedValue(null);
    const id = mongoose.Types.ObjectId().toHexString();
    const createdBy = mongoose.Types.ObjectId().toHexString();

    expect(async () => {
      await taskService.getTask(id, createdBy);
    }).rejects.toThrowError(/No task found/);
    expect(Task.findById).toBeCalledTimes(1);
  });

  it("should throw if user don't have permission to access this task", async () => {
    Task.findById = jest.fn().mockResolvedValue({
      title: "a",
      createdBy: mongoose.Types.ObjectId(),
    });
    const id = mongoose.Types.ObjectId().toHexString();
    const createdBy = mongoose.Types.ObjectId().toHexString();

    const result = taskService.getTask(id, createdBy);
    await expect(result).rejects.toThrowError(/don't have permission/);
    expect(Task.findById).toBeCalledTimes(1);
  });

  it("should return a task by id and createdBy", async () => {
    const createdBy = mongoose.Types.ObjectId().toHexString();
    const id = mongoose.Types.ObjectId().toHexString();
    const task = {
      id,
      title: "a",
      createdBy,
    };

    Task.findById = jest.fn().mockResolvedValue(task);

    const result = taskService.getTask(id, createdBy);

    await expect(result).resolves.toMatchObject(task);
    expect(Task.findById).toBeCalledTimes(1);
  });
});

describe("updateTask", () => {
  it("should update task", async () => {
    const id = mongoose.Types.ObjectId().toHexString();
    const createdBy = mongoose.Types.ObjectId().toHexString();

    const oldTask = new Task({
      title: "a",
      priority: priorityTypes.HIGH,
      createdBy,
    });
    const updatingTask = {
      title: "b",
      priority: priorityTypes.LOW,
    };

    Task.findById = jest.fn().mockResolvedValue(oldTask);
    taskService.getTask = jest.fn().mockResolvedValue(oldTask);
    oldTask.save = jest.fn();

    const result = taskService.updateTask(id, updatingTask, createdBy);
    await expect(result).resolves.toMatchObject(new Task(updatingTask));
  });

  it("deleteTask", async () => {
    const createdBy = mongoose.Types.ObjectId().toHexString();
    const id = mongoose.Types.ObjectId().toHexString();
    const task = {
      id,
      title: "a",
      createdBy,
    };

    Task.findById = jest.fn().mockResolvedValue(task);

    taskService.getTask = jest.fn().mockResolvedValue(new Task(task));
    task.remove = jest.fn();

    await taskService.deleteTask(id, createdBy.toString());
    expect(task.remove).toBeCalledTimes(1);
  });
});

describe("countTasks", () => {
  it("should return count result", async () => {
    const from = mongoose.Types.ObjectId().toHexString();
    const createdBy = mongoose.Types.ObjectId().toHexString();
    const numberOfTask = 1;

    Task.count = jest.fn().mockResolvedValue(numberOfTask);

    const result = taskService.countTasks(from, createdBy);

    await expect(result).resolves.toBe(numberOfTask);
    expect(Task.findById).toBeCalledTimes(1);
  });
});
