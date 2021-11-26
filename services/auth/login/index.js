const express = require("express");
const { login } = require("../controllers/login");
const router = express.Router();

router.post("/", async (req, res) => {
  const { username, password } = req.body;
  const isLogin = login(username, password);
  if (!isLogin) {
    return res.status(401).json({ error: "Unauthorized" });
  }
  return res.status(200).json({ data: "testabc.xyz.ahk" });
});

module.exports = router;
