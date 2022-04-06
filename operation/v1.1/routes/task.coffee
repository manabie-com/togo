express        = require 'express'
router         = express.Router()
Controllers    = require '../controllers'
TaskController = Controllers['TaskController']

router.get '/', TaskController.getAllTasks
router.get '/by-user', TaskController.getUserTasks
router.post '/assign', TaskController.assignTask

module.exports = router
