var express = require('express');
var router = express.Router();
const userController = require('../controller/user.controller');

router.get('/list', userController.getUsers);

module.exports = router;
