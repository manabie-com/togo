module.exports = app => {
  const tasks = require('../controllers/task')
  const { authenticateJWT } = require('../common/middlewares/authenticate')
  const router = require('express').Router()

  // user add new task
  router.post('/', authenticateJWT, tasks.add)

  app.use('/api/tasks', router)
}
