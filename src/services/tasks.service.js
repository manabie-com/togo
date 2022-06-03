const mongoose = require("mongoose");
const moment = require("moment");
const Task = require("../models/tasks.model");

const getAll = async () => {
  return Task.find().lean();
};

const getAllTasksTodayByUser = async (id) => {
  return Task.find({ createdAt: { $gte: moment().startOf("day") }, user: id });
};

const create = async (body) => {
  const task = await new Task(body).save();
  return Task.findById(task._id).lean();
};

const update = async (id, body) => {
  return Task.findOneAndUpdate({ _id: id }, body, { new: true }).lean();
};

const remove = async (id) => {
  return Task.findByIdAndRemove(id).lean();
};

module.exports = {
  getAll,
  getAllTasksTodayByUser,
  create,
  update,
  remove
};