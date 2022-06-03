const tasksServices = require("../services/tasks.service");
const usersServices = require("../services/users.service");
const { Mutex } = require("async-mutex");

const mutex = new Mutex();

const get = async (req, res, next) => {
  try {
    res.json(await tasksServices.getAll());
  } catch (err) {
    console.error("Error while getting tasks", err.message);
    next(err);
  }
};

const create = async (req, res, next) => {
  const release = await mutex.acquire();
  try {
    const user = await usersServices.getUserByName(req.body.username);
    const tasksOfUserToday = await tasksServices.getAllTasksTodayByUser(user._id);
    if (tasksOfUserToday.length < user.limit) {
      req.body.user = user._id;
      const task = await tasksServices.create(req.body);
      res.json(task);
    } else throw new Error("You have reached your limit of tasks per day");
  } catch (err) {
    console.error("Error while creating task", err.message);
    next(err);
  } finally {
    release();
  }

};

const update = async (req, res, next) => {
  try {
    res.json(await tasksServices.update(req.params.id, req.body));
  } catch (err) {
    console.error("Error while updating task", err.message);
    next(err);
  }
};

const remove = async (req, res, next) => {
  try {
    res.json(await tasksServices.remove(req.params.id));
  } catch (err) {
    console.error("Error while deleting task", err.message);
    next(err);
  }
};

module.exports = {
  get,
  create,
  update,
  remove
};