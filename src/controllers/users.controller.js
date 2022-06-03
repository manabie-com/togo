const usersServices = require("../services/users.service");

async function get(req, res, next) {
  try {
    res.json(await usersServices.getAll());
  } catch (err) {
    console.error("Error while getting users", err.message);
    next(err);
  }
}

async function create(req, res, next) {
  try {
    res.json(await usersServices.create(req.body));
  } catch (err) {
    console.error("Error while creating user", err.message);
    next(err);
  }
}

async function update(req, res, next) {
  try {
    res.json(await usersServices.update(req.params.id, req.body));
  } catch (err) {
    console.error("Error while updating user", err.message);
    next(err);
  }
}

async function remove(req, res, next) {
  try {
    res.json(await usersServices.remove(req.params.id));
  } catch (err) {
    console.error("Error while deleting user", err.message);
    next(err);
  }
}

module.exports = {
  get,
  create,
  update,
  remove
};