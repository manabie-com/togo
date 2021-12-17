const express = require('express');
const router = express.Router();
const users = require('../controllers/users');

router.route('/')
    .get(users.getAllUsers)
    .post(users.createUser);

router.route('/:id')
    .patch(users.updateUser)
    .delete(users.deleteUser);

module.exports = router;