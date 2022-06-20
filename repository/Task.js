/**
 * @author Nguyen Minh Tam / ngmitamit@gmail.com
 */

const {TaskMetaData, Task} = require("../DB/models/Task");

const { v4: uuidv4 } = require("uuid");

const { getTodayString } = require("../utils/index");

/**
 *
 * @param {string} userId
 * @returns list of task user who is the owner of this userId
 */
const getTaskListByUserId = (userId) => {
  const userTask = TaskMetaData[userId];

  if (!userTask) return [];

  const taskList = [];

  for (const taskId in userTask) {
    taskList.push(Task[taskId]);
  }

  return taskList;
};

module.exports = {
  /**
   *
   * @param {string} id
   * @returns task's data or undefined if id is invalid
   */
  getTaskById: (id) => {
    return Task[id];
  },

  getTaskListByUserId,

  /**
   *
   * @param {string} userId
   * @param {string} day - dd-mm-yyyy
   * @returns
   */
  getNumberOfTaskByUserIdAndDay: (userId, day) => {
    const taskList = getTaskListByUserId(userId) || [];

    const result = taskList.reduce((total, task) => {
      return total + (task.createdAt === day);
    }, 0);

    return result;
  },

  /**
   *
   * @param {string} userId
   * @param {object} task
   * @returns task id
   */
  insertTaskByUserId: (userId, task) => {
    const content = "" + task.content;
    const id = uuidv4();

    Task[id] = {
      id,
      content,
      userId,
      createdAt: getTodayString(),
    };

    if (!TaskMetaData[userId]) TaskMetaData[userId] = {};

    TaskMetaData[userId][id] = true;

    return id;
  },
};
