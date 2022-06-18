// mock User data
const Users = {
  A: {
    id: "A",
    taskLimitPerDay: 5,
  },
  B: {
    id: "B",
    taskLimitPerDay: 10,
  },
  C: {
    id: "C",
    taskLimitPerDay: 0,
  },
};

module.exports = {
  /**
   *
   * @param {string} id
   * @returns user's data or undefined if id is invalid
   */
  getUserById: (id) => {
    return Users[id];
  },
};
