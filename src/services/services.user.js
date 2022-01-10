const joi = require('joi');
const bcrypt = require('bcrypt');
const moment = require('moment-timezone');
const jwt = require('jsonwebtoken');
const { userModel } = require('../models/model.user');
const { validateObject } = require('../utils');
const jwtSecret = process.env.JWT_SECRET || '&^%D&%&*^SD%^&777sd%S%D$%SD';

const userSchema = joi.object({
  username: joi.string().min(6).max(32).required(),
  password: joi.string().min(6).required()
});

async function hashPassword(raw) {
  return await new Promise((resolve) => {
    bcrypt.genSalt(10, (error, salt) => {
      bcrypt.hash(raw, salt, (err, hash) => {
        resolve(hash);
      });
    });
  })
}

async function comparePassword(raw, hash) {
  return await bcrypt.compare(raw, hash);
}

async function signJWT(payload = {}, exp) {
  return await jwt.sign({
    ...payload,
    exp: moment().unix() + exp
  }, jwtSecret);
}

async function verifyToken(token = '') {
  let verify = false;

  try {
    verify = jwt.verify(token, jwtSecret)
  } catch (error) { }

  return verify;
}

async function createUser(user = {}) {
  const result = {
    success: false,
    code: 0,
    data: {},
    message: ''
  }

  // Check data payload
  const valid = validateObject(userSchema, user);
  if (!valid.valid) {
    result.code = 400;
    result.message = valid.message;
    return result;
  }

  user.password = await hashPassword(user.password);
  await userModel.create(user)
    .then(res => {
      result.success = true;
      result.code = 201;
      result.data = res;
    })
    .catch(err => {
      if (err.message.includes('duplicate')) {
        result.code = 409;
        result.message = `username already exists!`;
      } else {
        result.message = err.message;
      }
    });

  return result;
}

async function login(username, rawPassword) {
  const result = {
    success: false,
    code: 0,
    data: {},
    message: ''
  }

  if (!username ||
    !rawPassword ||
    typeof (username) !== 'string' ||
    typeof (rawPassword) !== 'string') {

    result.code = 400;
    result.message = `Invalid input username or password!`;
    return result;
  }

  const user = await userModel.find({
    username: username
  }).then(res => res[0]);

  if (!user) {
    result.code = 404;
    result.message = 'Not found user!';
    return result;
  }

  const valid = await comparePassword(rawPassword, user.password);

  if (!valid) {
    result.code = 400;
    result.message = 'Invalid username or password!';
    return result;
  }

  const token = await signJWT({
    sub: user.id
  }, 86400);

  result.success = true;
  result.code = 200;
  result.data = token;

  return result;
}

module.exports = {
  createUser,
  login,
  verifyToken,
  hashPassword
}
