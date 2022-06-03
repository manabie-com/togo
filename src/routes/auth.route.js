const express = require("express");
const router = new express.Router();
const authenController = require("../controllers/authen.controller");

const publicEndpoint = [
  {
    method: "POST",
    url: "\/users\/register$",
  },
];

const publicResources = (req) => {
  for (let i = 0; i < publicEndpoint.length; i++) {
    if ((new RegExp(publicEndpoint[i].url)).test(req.originalUrl) &&
      (publicEndpoint[i].method == "ANY" || publicEndpoint[i].method == req.method)) {
      return true;
    }
  }
  return false;
};

router.post("/login", (req, res, next) => {
  authenController.getToken(req.body.username, req.body.password)
    .then((data) => res.status(200).json(data))
    .catch(next);
});

router.use("/*", (req, res, next) => {
  if (publicResources(req)) {
    next();
  } else {
    authenController.isValidToken(req.headers)
      .then(() => next())
      .catch(next);
  }
});



module.exports = router;
