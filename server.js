const express = require('express');
const cors = require('cors');
const mongoose = require('mongoose');
const TaskController = require('./controller/TaskController');
const UserController = require('./controller/UserController');
const SchedulerController = require('./controller/ScheduledController');

//accessing env file
require('dotenv').config();

const app = express();
const port = process.env.PORT || 5000;

//middleware
app.use(cors());
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

//route
app.post("/api/task/addTask", TaskController.addTask);
app.get("/api/task/getTask", TaskController.getTask);
app.post("/api/user/register", UserController.registration);
app.post("/api/user/login", UserController.login);

//Reset Limitt Every Day
SchedulerController.scheduled();
//port
app.listen(port, () => {
    console.log(`Server is running in port : ${port}`);
})