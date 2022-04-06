express        = require 'express'
router         = express.Router()
Controllers    = require '../controllers'
UserController = Controllers['UserController']

router.post '/sign-up', UserController.signUp
router.get '/get-all', UserController.getAllUsers

module.exports = router
