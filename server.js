global.BASE_DIR = __dirname;

const express = require("express");
const mongoose = require("mongoose");
const serverConfig = require("./src/configs/server.config");
const dbConfig = require("./src/configs/db.config");
const authRouter = require("./src/routes/auth.route");
const tasksRouter = require("./src/routes/tasks.route");
const usesRouter = require("./src/routes/users.route");

const app = express();

// Set up mongoose connection
mongoose.connect(process.env.NODE_ENV === "dev" ? dbConfig.db.url : dbConfig.db.urlTest, { useNewUrlParser: true });
mongoose.connection.on("error", console.error.bind(console, "MongoDB connection error:"));

app.use(express.json());
app.use(express.urlencoded({ extended: false }));

app.use("/api", authRouter);
app.use("/api/tasks", tasksRouter);
app.use("/api/users", usesRouter);

//Handle 404
app.use((req, res) => {
  res.status(404).json({ message: `CANNOT ${req.method} API ${req.originalUrl}` });
});

//Handle Error
app.use((err, req, res, _next) => {
  res.status(err.statusCode || 500).json({ message: err.message || String(err) });
});

const server = app.listen(serverConfig.port.http, "0.0.0.0", (err) => {
  if (err) {
  } else {
    console.log(`Togo app listening at http://localhost:${serverConfig.port.http}`);
  }
});

module.exports = server;
