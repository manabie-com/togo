const { Router } = require('express');
const auth = require('../middlewares/authenticate');
const TaskModule = require('../modules/tasks');
const router = Router();

router
  .route('/tasks')
  .get(auth('authorize'), TaskModule.getTasks);

router
  .route('/tasks')
  .post(auth('authorize'), TaskModule.createTask);

module.exports = router;
