const joi = require('joi');
const bcrypt = require('bcrypt');
const { userModel } = require('../models/model.user');
const { validateObject } = require('../utils');

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
      }
      result.message = err.message;
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
    result.message = 'Invalid payload!';
    return result;
  }

  const user = await userModel.find({
    username: username
  }).then(res => res[0]);

  if (!user) {
    result.code = 400;
    result.message = 'Invalid username or password!';
    return result;
  }

  const valid = await comparePassword(rawPassword, user.password);
  // RETURN JWT TOKEN HERE

  return user;
}

module.exports = {
  createUser,
  login
}
