const express = require("express");

const { authValidation } = require("../../validations");
const { authController } = require("../../controllers");

const validate = require("../../../middlewares/validate");

const router = express.Router();

router.post(
  "/register",
  validate(authValidation.registerSchema),
  authController.register
);

router.post(
  "/login",
  validate(authValidation.loginSchema),
  authController.login
);

module.exports = router;

/**
 * @swagger
 * tags:
 *   name: Auth
 *   description: Authentication
 */

/**
 * @swagger
 * /v1/auth/register:
 *   post:
 *     summary: Register as user
 *     tags: [Auth]
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             required:
 *               - name
 *               - email
 *               - password
 *             properties:
 *               name:
 *                 type: string
 *               email:
 *                 type: string
 *                 format: email
 *                 description: must be unique
 *               password:
 *                 type: string
 *                 format: password
 *                 minLength: 6
 *                 description: At least one number and one letter
 *               maxTask:
 *                 type: integer
 *                 minimum: 1
 *                 description: must be integer, 1 is minimum
 *             example:
 *               name: Manabie's user
 *               email: manabie1@manabie.com
 *               password: manabie1
 *               maxTask: 10
 *     responses:
 *       "201":
 *         description: Created
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 user:
 *                   $ref: '#/components/schemas/User'
 *                 token:
 *                   type: string
 *               example:
 *                 user:
 *                   _id: 6283d4900000000000000000
 *                   name: Manabie's user
 *                   email: manabie1@manabie.com
 *                   maxTask: 10
 *                 token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjYyODUwNGFlMjA5YmFhZmViNjY5M2RhZiIsIm1heFRhc2siOjEwLCJpYXQiOjE2NTI5NTQ0Mzh9.No2n32pB8whz0M1yF8RHB14P0mrXZlqYB4QgV6uHL1s
 *       "400":
 *         $ref: '#/components/responses/DuplicateEmail'
 */

/**
 * @swagger
 * /v1/auth/login:
 *   post:
 *     summary: Login
 *     tags: [Auth]
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             required:
 *               - email
 *               - password
 *             properties:
 *               email:
 *                 type: string
 *                 format: email
 *               password:
 *                 type: string
 *                 format: password
 *                 minLength: 6
 *                 description: At least one number and one letter
 *             example:
 *               email: manabie1@manabie.com
 *               password: manabie1
 *     responses:
 *       "200":
 *         description: Login successfully
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 user:
 *                   $ref: '#/components/schemas/User'
 *                 token:
 *                   type: string
 *               example:
 *                 user:
 *                   _id: 6283d4900000000000000000
 *                   name: Manabie's user
 *                   email: manabie1@manabie.com
 *                   maxTask: 10
 *                 token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjYyODUwNGFlMjA5YmFhZmViNjY5M2RhZiIsIm1heFRhc2siOjEwLCJpYXQiOjE2NTI5NTQ0Mzh9.No2n32pB8whz0M1yF8RHB14P0mrXZlqYB4QgV6uHL1s
 *       "401":
 *         $ref: '#/components/responses/UnauthorizedLogin'
 */
