/**
 * @author Nguyen Minh Tam / ngmitamit@gmail.com
 */

const constants = require("./constants");

const { ensureAuthenticated } = require("./middlewares/auth");

const taskRoutes = require("./routes/task");

module.exports = (app) => {
  app.use(ensureAuthenticated);
  app.use(`${constants.baseApi}`, taskRoutes);
};
