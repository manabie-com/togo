const httpStatus = require("http-status");
const { get } = require("lodash");

const catchAsync = require("../../utils/catch-async");
const { taskService } = require("../services");

const createTask = catchAsync(async (req, res) => {
  const task = await taskService.createTask(req.body, req.user);
  res.status(httpStatus.CREATED).send(task);
});

const getTasks = catchAsync(async (req, res) => {
  const tasks = await taskService.getTasks(req.user._id);
  res.status(httpStatus.OK).send(tasks);
});

const getTaskById = catchAsync(async (req, res) => {
  const createdBy = get(req.user, "_id", "");
  const task = await taskService.getTask(req.params.id, createdBy);
  res.status(httpStatus.OK).send(task);
});

const updateTask = catchAsync(async (req, res) => {
  const createdBy = get(req.user, "_id", "");

  const task = await taskService.updateTask(req.params.id, req.body, createdBy);
  res.status(httpStatus.OK).send(task);
});

const deleteTask = catchAsync(async (req, res) => {
  const createdBy = get(req.user, "_id", "");

  await taskService.deleteTask(req.params.id, createdBy);
  res.status(httpStatus.NO_CONTENT).send();
});

module.exports = {
  createTask,
  getTasks,
  getTaskById,
  updateTask,
  deleteTask,
};
