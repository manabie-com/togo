const express = require("express");
const router = express.Router();
const usesController = require("../controllers/users.controller");

router.get("/", usesController.get);

router.post("/register", usesController.create);

router.put("/:id", usesController.update);

router.delete("/:id", usesController.remove);

module.exports = router;