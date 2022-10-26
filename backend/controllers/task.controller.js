const httpStatus = require('http-status');
const ApiError = require('../utils/ApiError');
const catchAsync = require('../utils/catchAsync');
const { taskService } = require('../services');
const response = require('../utils/responseTemp');

/**
 * Create task
 */
const createTask = catchAsync(async (req, res) => {
  const task = await taskService.createTask(req.body);
  res.send(response(httpStatus.OK, 'Create Task success', task));
});

/**
 * Get task infor
 */
const getTask = catchAsync(async (req, res) => {
  const task = await taskService.getTaskByPk(req.task.id);
  if (!task) {
    throw new ApiError(httpStatus.NOT_FOUND, 'Task not found');
  }
  res.send(response(httpStatus.OK, 'Get Task success', task));
});

/**
 * Get tasks
 */
const getTasks = catchAsync(async (req, res) => {
  const tasks = await taskService.getTaskByUserId(req.user.id);
  res.send(response(httpStatus.OK, 'Get Task success', tasks));
});

/**
 * Update task
 */
const updateTask = catchAsync(async (req, res) => {
  const user = await taskService.updateTaskByPk(req.params.taskId, req.body);
  res.send(response(httpStatus.OK, 'Update task success', user));
});

/**
 * Delete an user
 */
const deleteTask = catchAsync(async (req, res) => {
  await taskService.deleteTaskByPk(req.params.taskId);
  res.send(response(httpStatus.OK, 'Delete task success'));
});

module.exports = {
  updateTask,
  deleteTask,
  getTask,
  getTasks,
  createTask,
};
