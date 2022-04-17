const express = require("express");
const router = express.Router();

const { addTask } = require("../controller/taskController");
const mongoRepo = require("../data/mongoRepo");

require("dotenv").config();

let currentDate = new Date();

// POST method for adding task; uses dependency injection
router.post("/", addTask(mongoRepo, currentDate.toISOString()));

module.exports = router;
