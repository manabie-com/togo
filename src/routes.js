'use strict';


const Router = require('koa-router');
const router = new Router();
const home = require('./controllers/home');
const auth = require('./controllers/controllers.auth');
const task = require('./controllers/controllers.task');

router.get('/api/public/home', home.home);
router.post('/api/public/auth/sign-up', auth.signUp);
router.post('/api/public/auth/sign-in', auth.signIn);

router.get('/api/me/task', task.listTask);
router.post('/api/me/task', task.createTask);
router.put('/api/me/task', task.updateTask);
router.delete('/api/me/task', task.deleteTask);

module.exports = router;
