const express = require("express");
const noteRoute = express.Router();
const {
  create,
  makeTaskCompleted,
  getList,
  updateTask,
  deleteTask
} = require("../controllers/note.controller");
const checkLogin = require('../middlewares/checkLogin.middleware')
const addTaskValidator = require('./validators/addTask.validator')

noteRoute.route("/create").post(checkLogin, addTaskValidator, create);

noteRoute.route("/task/tick/:id").put(checkLogin, makeTaskCompleted);

noteRoute.route("/tasks").get(checkLogin, getList);

noteRoute.route("/task/update/:id").put(checkLogin, updateTask)

noteRoute.route("/task/delete/:id").delete(checkLogin, deleteTask);

module.exports = noteRoute;
