const Task = require("../model/task");

// Controller addTask  method that receives the repository (can be MongoDB or other DB) and current date; uses dependency injection
const addTask = (repository, currentDate) => async (req, res, next) => {
  const { user_name, task_name } = req.body;
  let task = new Task(task_name, user_name);

  // Calls the addTask method of the injected repository and waits for response
  const result = await repository.addTask(task, currentDate);

  // Result contains the error object; if null, the task is successfully created; otherwise, return the error object
  if (!result) {
    res.status(201).send("Task created");
  } else {
    return next(result);
  }
};

module.exports = { addTask };
