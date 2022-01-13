const express = require("express");
const cors = require("cors");
const mongoose = require("mongoose");

// The main router that handles the request
const taskRoute = require("./routes/taskRoute");

const app = express();

require("dotenv").config();

app.use(cors());
app.use(express.json());

// The endpoint that receives task information
app.use("/task", taskRoute);

// Returns an error if the endpoint or HTTP method is not supported
app.use((req, res, next) => {
  return next(new Error("Invalid route"));
});

// Generic error handler; the error is returned from the router
app.use((error, req, res, next) => {
  console.log(error);
  res.status(error.code || 500).send(error.message);
});

module.exports = app;
