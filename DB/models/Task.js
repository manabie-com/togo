/**
 * @author Nguyen Minh Tam / ngmitamit@gmail.com
 */

// mock Task data
const TaskMetaData = {
  A: {
    task1: true,
  },
  B: {
    task2: true,
  },
};
const Task = {
  task1: {
    id: "task1",
    content: "task 1",
    userId: "A",
    createdAt: "18-06-2022",
  },
  task2: {
    id: "task2",
    content: "task 1",
    userId: "B",
    createdAt: "18-06-2022",
  },
};

module.exports = {
  TaskMetaData,
  Task,
};
