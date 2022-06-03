const mongoose = require("mongoose");
const User = require("../models/users.model");
const { encrypt } = require("../utils/helper.util");

const getAll = async () => {
  return User.find().select("-password").lean();
};

const getUserByName = async (username) => {
  return User.findOne({ username }).select("-password").lean();
};

const create = async (body) => {
  body.password = encrypt(body.password);
  const newUser = await new User(body).save();
  return User.findById(newUser._id).select("-password").lean();
};

const update = async (id, body) => {
  return User.findOneAndUpdate({ _id: id }, body, { new: true }).select("-password").lean();
};

const remove = async (id) => {
  return User.findByIdAndRemove(id).select("-password").lean();
};

module.exports = {
  getAll,
  getUserByName,
  create,
  update,
  remove
};