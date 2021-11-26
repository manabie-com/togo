const express = require("express");
const validAuthorized = require("../../../helpers/authorize");
const { getTodos } = require("../controllers/getTodos");
const router = express.Router();
const TODO_MODEL_URL = "models/todos/todos.json";

router.get("/", (req, res) => {
  const todos = getTodos(TODO_MODEL_URL);
  if (!todos) {
    return res.status(400).json({ error: todos });
  }
  return res.status(200).json({ data: todos });
});

module.exports = router;
