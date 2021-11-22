const express = require("express");
const router = express.Router();

router.get("/test", (req, res) => {
  res.status(200).json({
    success: "true",
  });
});

router.use('/login', require('../services/auth'));
router.use('/tasks', require('../services/todos/list'));

module.exports = router;
