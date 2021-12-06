const authRoute = require("./auth.route");
const noteRoute = require('./note.route');
const route = (app) => {
  app.use("/api/auth", authRoute);
  app.use("/api/todo", noteRoute);
};

module.exports = route;
