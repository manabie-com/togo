const express = require('express');
const userRoutes = require('./users');
const taskRoutes = require('./tasks');
const authRoutes = require('./auth');

const router = express.Router();

router.use(authRoutes);
router.use(userRoutes);
router.use(taskRoutes);

module.exports = router;
