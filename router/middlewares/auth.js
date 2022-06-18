/**
 * @author Nguyen Minh Tam / ngmitamit@gmail.com
 */

const Users = require("../../DB/models/User");

const user = {
  // token -> userId,
  "userA:password": "A",
  "userB:password": "B",
};

/**
 *
 * @param {string} token
 * @returns userData or undefined if token is invalid
 */
function getUserDataByToken(token) {
  const userId = user[token];

  if (!userId) return undefined;

  return Users.getUserById(userId);
}

module.exports = {
  /**
   * Authenticate user and add userData to req
   * If the user is not Authenticated return with HTTP 401 error
   */
  ensureAuthenticated(req, res, next) {
    if (!req.cookies["access_token"]) {
      return res.sendStatus(401);
    }

    const token = req.cookies["access_token"];
    const userData = getUserDataByToken(token);

    if (!userData) return res.sendStatus(401);

    req.userData = userData;
    next();
  },
};
