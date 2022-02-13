const router = require('express').Router();
const auth = require('../middlewares/auth')
const access = require('../middlewares/limiter')
const UserController = require('../controllers/userController');

router.post('/login', UserController.login);
router.post('/register', UserController.register);
router.post('/', auth.verify, access.limiter, UserController.addTask);
router.get('/', auth.verify, UserController.getTask);

module.exports = router;