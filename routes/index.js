const express = require("express");
const router = express.Router();
const cors = require("cors");
const validAuthorized = require("../helpers/authorize");

router.use((req, res, next) => {
  res.header("Access-Control-Allow-Origin", req.get("Origin") || "*");
  res.header(
    "Access-Control-Allow-Headers",
    "Origin, X-Requested-With, Content-Type, Accept"
  );
  next();
});
router.use(cors());

router.use(
  express.urlencoded({
    extended: true,
  })
);
router.use(express.json());

router.use(async (req, res, next) => {
  if (req.path === "/login") {
    return next();
  }
  const { authorization } = req.headers;
  if (!validAuthorized(authorization)) {
    return res.status(401).json({ error: "Unauthorized" });
  }
  return next();
});

router.use("/login", require("../services/auth/login"));
router.use("/tasks", require("../services/todos/list"));

module.exports = router;
