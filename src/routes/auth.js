const { Router } = require('express');
const AuthModule = require('../modules/authenticate');

const router = Router();

router.route('/register').post(AuthModule.register);

router.route('/login').post(AuthModule.login);

module.exports = router;
