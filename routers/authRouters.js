const express = require("express");
const authController = require("../controllers/authController");

const router = express.Router();
router.post("/login", authController.postUserLogin);
router.get("/logout", authController.userLogout);

module.exports = router;