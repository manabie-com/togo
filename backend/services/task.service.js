const httpStatus = require('http-status');
const { tasks } = require('../models');
const ApiError = require('../utils/ApiError');

/**
 * Create an task
 * @param {Object} userBody
 * @returns {Promise<Task>}
 */
const createTask = async (userBody) => {
  return tasks.create({ ...userBody });
};

/**
 * Get task by pk
 */
const getTaskByPk = async (id) => {
  return tasks.findByPk(id);
};

/**
 * Get tasks by userId
 */
const getTaskByUserId = async (userId) => {
  return tasks.findAll({ where: { user_id: userId } });
};

/**
 * Update task by pk
 * @param {number} taskId
 * @param {Object} updateBody
 * @returns {Promise<task>}
 */
const updateTaskByPk = async (taskId, updateBody) => {
  const task = await getTaskByPk(taskId);
  if (!task) {
    throw new ApiError(httpStatus.NOT_FOUND, 'Task not found');
  }
  Object.assign(task, updateBody);
  await task.save();
  return task;
};

/**
 * Delete task by Pk
 */
const deleteTaskByPk = async (taskId) => {
  const task = await getTaskByPk(taskId);
  if (!task) {
    throw new ApiError(httpStatus.NOT_FOUND, 'User not found');
  }
  await task.destroy();
  return task;
};

module.exports = {
  getTaskByPk,
  updateTaskByPk,
  deleteTaskByPk,
  createTask,
  getTaskByUserId,
};
