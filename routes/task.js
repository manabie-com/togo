const express = require('express');
const router = express.Router();
const userModel = require('../models/user_model');
const taskUserModel = require('../models/task_user_model');

// middleware
router.use((req, res, next) => {
  if (typeof req.body.user_id === 'undefined' || typeof req.body.task_name === 'undefined') {
    return res.status(401).send({ error: 'invalid data' });
  }

  const {user_id, task_name} = req.body;

  const findUser = userModel.findById(user_id);
  if (!findUser) {
    return res.status(404).send({ error: 'User Not Found' });
  }

  const today = new Date();
  const date = today.getFullYear()+'-'+(today.getMonth()+1)+'-'+today.getDate();
  req.requestTime = date;

  const finded = taskUserModel.find(user_id, date);
  if (finded && finded.length >= findUser.metadata.max_tasks) {
    return res.status(403).send({ error: 'You can not add task now' });
  }
  
  next()
})

router.post('/', function (req, res) {
  const {user_id, task_name} = req.body;
  const requestTime = req.requestTime;

  const result = taskUserModel.save({user_id, task_name, date: requestTime})

  return res.status(200).send(result);
})

module.exports = router;