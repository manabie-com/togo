const masterdata = require('../../utils/masterdata');

module.exports = {
  users: [
    ...masterdata.users,
    {
      id: 2,
      max_todo: 1,
      username: 'user2',
      password: 'pass',
    }
  ],
  tasks: [
    {
      id: 1,
      content: 'Task 1',
      user_id: 1,
    },
    {
      id: 2,
      content: 'Task 2',
      user_id: 1,
    },
    {
      id: 3,
      content: 'Task 3',
      user_id: 2,
    },
  ],
};
