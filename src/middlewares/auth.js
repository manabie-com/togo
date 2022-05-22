const { UNAUTHORIZED } = require("http-status");
const jwt = require("jsonwebtoken");

const { authService } = require("../apis/services");
const env = require("../configs/env");

module.exports = function (req, res, next) {
  try {
    const token = authService.getTokenFromHeaders(req.headers);

    const decoded = jwt.verify(token, env.jwt.secret);
    req.user = decoded;
    return next();
  } catch (ex) {
    res
      .status(ex.statusCode ? ex.statusCode : UNAUTHORIZED)
      .send("Please authenticate");
  }
};
