// mock Task data
const TaskMetaData = {
  A: {
    task1: true,
  },
};
const Task = {
  task1: {
    id: "task1",
    content: "task 1",
    userId: "A",
    createdAt: "18-06-2022",
  },
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

  /**
   *
   * @param {string} userId
   * @returns list of task user who is the owner of this userId
   */
  getTaskListByUserId: (userId) => {
    const userTask = TaskMetaData[userId];

    if (!userTask) return [];

    const taskList = [];

    for (const taskId in userTask) {
      taskList.push(Task[taskId]);
    }

    return taskList;
  },
};
