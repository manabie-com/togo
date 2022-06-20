/**
 * @author Nguyen Minh Tam / ngmitamit@gmail.com
 */

const { Users } = require("../DB/models/User");

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
