const jwt = require("jsonwebtoken");

module.exports = (req, res, next) => {
  let authorization = req.header("Authorization");
  if (!authorization) {
    return res.status(401).json({ message: "Token has expired" });
  }

  let token = authorization.split(" ")[1];
  if (!token) {
    return res.status(401).json({ message: "Invalid token" });
  }

  const JWT_SECRET = process.env.JWT_SECRET || "adwkji8ad7w65wa3d1s3adw";
  jwt.verify(token, JWT_SECRET, (err, data) => {
    if (err) {
      return res.status(401).json({ message: "Token has expired or invalid" });
    }
    req.user = data;
    next();
  });
};
