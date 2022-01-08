'use strict';


const Router = require('koa-router');
const router = new Router();
const home = require('./controllers/home');


router.get('/api/public/home', home.home);

module.exports = router;
