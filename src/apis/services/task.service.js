const httpStatus = require("http-status");

const ApiError = require("../../utils/api-error");
const { Task } = require("../models");
const { getStartOfDay, objectIdFromDate } = require("../../utils/common");

/**
 * Create a task
 * @param {Object} taskBody
 * @param {Object} user
 * @returns {Promise<Task>}
 */
const createTask = async (taskBody, user) => {
  const startOfDay = getStartOfDay();
  const taskInDay = await countTasks(objectIdFromDate(startOfDay), user._id);
  if (taskInDay >= user.maxTask) {
    throw new ApiError(httpStatus.BAD_REQUEST, "Reach task limit in day");
  }

  return Task.create({ ...taskBody, createdBy: user._id });
};

/**
 * Get tasks by created user
 * @param {string} createdBy
 * @returns {Promise<Task>}
 */
const getTasks = async (createdBy) => {
  return Task.find({ createdBy }).sort({ _id: -1 }).lean();
};

/**
 * Count tasks
 * @param {string} from
 * @param {string} createdBy
 * @returns {Promise<number>}
 */
const countTasks = async (from, createdBy) => {
  return Task.count({ _id: { $gte: from }, createdBy });
};

/**
 * Get a task by id
 * @param {string} id
 * @param {string} createdBy
 * @returns {Promise<Task>}
 */
const getTask = async (id, createdBy) => {
  const task = await Task.findById(id);
  if (!task) {
    throw new ApiError(httpStatus.NOT_FOUND, `No task found with ID: ${id}`);
  }

  if (task.createdBy.toString() !== createdBy) {
    throw new ApiError(
      httpStatus.FORBIDDEN,
      "You don't have permission to access this resource"
    );
  }

  return task;
};

/**
 * Update a task
 * @param {string} id
 * @param {Object} taskBody
 * @param {Object} createdBy
 * @returns {Promise<Task>}
 */
const updateTask = async (id, taskBody, createdBy) => {
  const task = await getTask(id, createdBy);

  Object.keys(taskBody).forEach((key) => {
    task[key] = taskBody[key];
  });
  await task.save();

  return task;
};

/**
 * Get a task by id
 * @param {string} id
 * @param {string} createdBy
 * @returns {Promise<any>}
 */
const deleteTask = async (id, createdBy) => {
  const task = await getTask(id, createdBy);

  return task.remove();
};

module.exports = {
  createTask,
  getTasks,
  countTasks,
  getTask,
  updateTask,
  deleteTask,
};
