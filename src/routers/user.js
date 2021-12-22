const express = require('express')
const auth = require('../middleware/auth')
const {createUser, login, getUser, logOut, logOutAll,limitTaskUser} = require('../controller/user')
const userRouter = express.Router()
const userValidate = require('./validators/user.validator')

userRouter.post('/register',userValidate, createUser)
userRouter.post('/login', userValidate, login)
userRouter.get('/me', auth, getUser)

userRouter.post('/me/logout', auth, logOut)
userRouter.post('/me/logout-all', auth, logOutAll)
userRouter.post('/:id/limit-task', auth, limitTaskUser)
module.exports = userRouter