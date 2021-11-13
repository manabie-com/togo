const { Router } = require('express');
const auth = require('../middlewares/authenticate');
const UserModule = require('../modules/users');
const router = Router();

router
  .route('/users')
  .get(auth('authorize'), UserModule.getUserInfo);

router
  .route('/users')
  .put(auth('authorize'), UserModule.updateUserInfo);

module.exports = router;
