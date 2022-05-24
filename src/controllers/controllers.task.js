'use strict!';
const moment = require('moment-timezone');
const taskService = require('../services/services.task');

// GET: /api/me/task?from=YYYY-MM-DD&to=YYYY-MM-DD&status=-1
module.exports.listTask = async (ctx) => {
  const userId = ctx.user.sub || '';
  const pageSize = ctx.query.pageSize || '10';
  const currentPage = ctx.query.currentPage || '0';
  const from = moment(ctx.query.from || '1970-01-01').startOf('day').toString();
  const to = moment(ctx.query.to || '1970-01-01').endOf('day').toString();

  if (from === 'Invalid date' || to === 'Invalid date') {
    return ctx.showError('Invalid date');
  }

  // Date query: YYYY-MM-DD
  const filterContext = [{
    userId: userId
  }, {
    createdAt: {
      $gte: from,
      $lte: to
    }
  }];

  const status = Number(ctx.query.status);

  if (!isNaN(status) && typeof (status) === 'number') {
    filterContext.push({
      status: status
    })
  }

  const list = await taskService.listTask({
    $and: filterContext
  },
    pageSize,
    currentPage,
    {
      createdAt: -1
    }
  );

  return ctx.showResult(list);
}

// POST: /api/me/task
module.exports.createTask = async (ctx) => {
  const userId = ctx.user?.sub || '';
  const payload = ctx.request.body || {};
  payload.userId = userId;
  const availability = await taskService.count(userId, moment().format('YYYY-MM-DD').toString());

  if (!availability) {
    return ctx.showError('Maximum limit daily task!');
  }

  const result = await taskService.createTask(payload);
  return result.success ? ctx.showResult(result.data, 201) : ctx.showError(result.message, 400);
}

// PUT: /api/me/task?id=acbd
module.exports.updateTask = async (ctx) => {
  const userId = ctx.user?.sub || '';
  const id = ctx.query.id || '';
  const target = await taskService.readTask(id);

  if (!target || target.userId.toString() !== userId) {
    return ctx.showError('Not found task id!', 404);
  }

  const payload = ctx.request.body || {};
  payload.userId = userId;

  const result = await taskService.updateTask(id, payload);
  return result.success ? ctx.showResult(result.data) : ctx.showError(result.message);
}

// DELETE: /api/me/task?id=...
module.exports.deleteTask = async (ctx) => {
  const userId = ctx.user?.sub || '';
  const id = ctx.query.id || '';
  const target = await taskService.readTask(id);

  // if (!target || target.userId.toString() !== userId || target.status === -1) {
  //   return ctx.showError('Task not found!', 404);
  // }

  const result = await taskService.deleteTask(id);

  return result.success ? ctx.showResult(result.data) : ctx.showError(result.message);
}