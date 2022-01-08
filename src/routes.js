'use strict';


const Router = require('koa-router');
const router = new Router();
const home = require('./controllers/home');
const auth = require('./controllers/auth');
const task = require('./controllers/task');

router.get('/api/public/home', home.home);
router.post('/api/public/auth/sign-up', auth.signUp);
router.post('/api/public/auth/sign-in', auth.signIn);

module.exports = router;
