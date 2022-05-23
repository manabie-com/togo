const joi = require('joi');
const moment = require('moment-timezone');
const { validateObject } = require('../utils');
const { taskModel } = require('../models/model.task');
const { userModel } = require('../models/model.user');


const taskSchema = joi.object({
  userId: joi.string().min(16).max(128).required(),
  title: joi.string().required().min(1),
  description: joi.string().required(),
  status: joi.number().required().min(0).max(2) // 0: TODO, 1: In progress, 2: Done
});

async function count(userId) {
  const today = moment();
  const limit = await userModel.findById(userId)
    .then(res => {
      return res?.limit || -1;
    }).catch(err => {
      console.log(err);
      return -1;
    });

  if (limit < 0) return false;
  const dailyTask = await taskModel.count({
    $and: [{
      userId: userId
    }, {
      createdAt: {
        $gte: today.startOf('day').toString(),
        $lte: today.endOf('day').toString()
      }
    }]
  }).then(res => {
    return res;
  }).catch(err => {
    console.log(err);
    return Infinity;
  })

  return limit > dailyTask;
}

async function readTask(id) {
  return await taskModel.findById(id)
    .catch(err => {
      console.log(err);
      return null;
    });
}

async function createTask(task = {}) {
  const result = {
    success: false,
    code: 0,
    data: {},
    message: ''
  }

  const valid = validateObject(taskSchema, task);
  if (!valid.valid) {
    result.code = 400;
    result.message = valid.message || 'Invalid payload!';
    return result;
  }
  // Check quantity task in today of user

  await taskModel.create(task)
    .then(res => {
      result.code = 201;
      result.success = true;
      result.data = res;
      return result;
    })
    .catch(err => {
      result.code = 400;
      result.message = err.message;
      return result;
    });

  return result;
}

async function updateTask(id = '', task = {}) {
  const result = {
    success: false,
    code: 0,
    data: {},
    message: ''
  }

  const valid = validateObject(taskSchema, task);
  if (!valid.valid) {
    result.code = 400;
    result.message = valid.message || 'Invalid payload!';
    return result;
  }
  // Check quantity task in today of user

  await taskModel.findByIdAndUpdate(id, task)
    .then(res => {
      result.code = 200;
      result.success = true;
      result.data = res;
      return result;
    })
    .catch(err => {
      result.code = 400;
      result.message = err.message;
      return result;
    });

  return result;
}

async function deleteTask(id) {
  // Status: 0: TODO, 1: Inprogess: 2: Done. -1: Deleted
  const result = {
    success: false,
    code: 0,
    data: {},
    message: ''
  }

  const target = await taskModel.findById(id);
  if (!target || target.status < 0) {
    result.code = 404;
    result.message = `Not found task id ${id}`;
    return result;
  }

  await taskModel.findByIdAndUpdate(id, {
    status: -1
  }).then(res => {
    result.success = true;
    result.code = 200;
    result.data = res;
  }).catch(err => {
    result.success = false;
    result.code = 400;
    result.message = err.message;
  })


  return result;
}

async function listTask(filter = {}, pageSize = 10, currentPage = 0, sort = {}) {
  const list = await taskModel
    .find(filter)
    .limit(pageSize)
    .skip(currentPage * pageSize)
    .sort(sort)
    .then(res => res)
    .catch(err => {
      console.log(err);
      return [];
    });
  return list;
}

module.exports = {
  count,
  readTask,
  createTask,
  updateTask,
  deleteTask,
  listTask
}