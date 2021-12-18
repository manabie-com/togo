const express = require('express');
const router = express.Router();
const users = require('../controllers/users');
const { validateUser } = require('../middleware/users');

router.route('/')
    .get(users.getAllUsers)
    .post(validateUser, users.createUser);

router.route('/:id')
    .patch(users.updateUser)
    .delete(users.deleteUser);

module.exports = router;