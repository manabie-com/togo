const express = require("express");
const todoController = require("../controllers/todoController");

const router = express.Router();
router.get("/", todoController.getTodos);
router.post("/", todoController.createTodo);
router.get("/:id", todoController.getTodoById);

router.delete("/:id", todoController.deleteTodo);

router.patch("/:id", todoController.updateTodo);
module.exports = router;