const jwt = require("jsonwebtoken");
const mongoose = require("mongoose");
const authenConfig = require("../configs/authen.config");
const { matches, mandatory } = require("../utils/helper.util");

const tokenConfig = authenConfig.jwt;

const isValidToken = async (headers = mandatory("headers")) => {
  const readAuthorization = async (authorization) => {
    if (!authorization || !authorization.trim()) {
      throw new Error("Missing Authorization");
    }
    const parts = authorization.split("-");
    if (String(parts[0]).toUpperCase() === "JWT" && authorization.substring(4)) {
      return authorization.substring(4).trim();
    }
    throw new Error("Wrong Authorization");
  };

  const decodeToken = (token) => {
    return new Promise((resolve, reject) => {
      jwt.verify(token, tokenConfig.secretKey, { algorithm: tokenConfig.algorithm }, (err, payload) => {
        if (err) return reject(err);
        return resolve(payload);
      });
    });
  };

  const token = await readAuthorization(headers.authorization);
  const payload = await decodeToken(token);
  return payload;
};

const getToken = async (username = mandatory("username"), password = mandatory("password")) => {
  const User = mongoose.model("users");
  const user = await User.findOne({ username }).lean();
  if (user && matches(password, user.password)) {
    delete user.password;
    const payload = {
      username: user.username,
    };
    const token = jwt.sign(payload, tokenConfig.secretKey, { algorithm: tokenConfig.algorithm, expiresIn: tokenConfig.expiresInMinutes * 60 });
    return { user, token: `JWT-${token}` };
  }
  throw new Error("Invalid username or password");
};

module.exports = { isValidToken, getToken };

