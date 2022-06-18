/**
 * @author Nguyen Minh Tam / ngmitamit@gmail.com
 */

const express = require("express");

const router = express.Router();

const taskController = require("../controller/task");

router.get("/tasks", taskController.getTasks);

module.exports = router;
