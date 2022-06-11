const express = require('express');
const cors = require('cors');
const mongoose = require('mongoose');
const TaskController = require('./controller/TaskController');
const UserController = require('./controller/UserController');

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
//mongodb connection
// const uri = process.env.ATLAS_URI
// mongoose.connect(uri,{useNewUrlParser: true, useUnifiedTopology: true })
// const connection = mongoose.connection;
// connection.once('open', () => {
//     console.log("MongoDB database connection is established.")
// })

//port
app.listen(port, () => {
    console.log(`Server is running in port : ${port}`);
})