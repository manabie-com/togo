const express = require('express');
const auth = require('../../middlewares/auth');
const validate = require('../../middlewares/validate');
const userValidation = require('../../validations/user.validation');
const userController = require('../../controllers/user.controller');

const router = express.Router();

router
  .route('/:userId')
  .patch(auth(), validate(userValidation.updateLimitDailyTask), userController.updateUser)
  .get(auth('getInfor'), validate(userValidation.myProfile), userController.getMyProfile);

module.exports = router;

/**
 * @swagger
 * tags:
 *   name: Users
 *   description: User management and retrieval
 */

/**
 * @swagger
 * /user:
 *   get:
 *     summary: Get user infor
 *     description:  user get your profile
 *     tags: [Users]
 *     security:
 *       - bearerAuth: []
 *     responses:
 *       "200":
 *         description: OK
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 status:
 *                  type: integer
 *                  example: 200
 *                 message:
 *                  type: string
 *                  example: Get profile success
 *                 data:
 *                  type: object
 *                  $ref: '#/components/schemas/user'
 *       "401":
 *         $ref: '#/components/responses/Unauthorized'
 *       "404":
 *         $ref: '#/components/responses/NotFound'
 *
 *   patch:
 *     summary: Update user
 *     description: User update information.
 *     tags: [Users]
 *     security:
 *       - bearerAuth: []
 *     requestBody:
 *       required: true
 *       description: Update any field of user
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             properties:
 *              limit_daily_task:
 *                type: integer,
 *                example: 6
 *     responses:
 *       "204":
 *         description: OK
 *         content:
 *           application/json:
 *             schema:
 *                $ref: '#/components/schemas/user'
 *       "400":
 *         $ref: '#/components/responses/LimitTaskNotAvailable'
 *       "401":
 *         $ref: '#/components/responses/Unauthorized'
 *       "404":
 *         $ref: '#/components/responses/NotFound'
 */
