const express = require('express');
const auth = require('../../middlewares/auth');
const taskController = require('../../controllers/task.controller');

const router = express.Router();

router.route('/').post(auth(), taskController.createTask).get(auth(), taskController.getTasks);

router
  .route('/:taskId')
  .patch(auth(), taskController.updateTask)
  .get(auth(), taskController.getTask)
  .delete(auth(), taskController.deleteTask);

module.exports = router;

/**
 * @swagger
 * tags:
 *   name: Tasks
 *   description: Tasks management and retrieval
 */

/**
 * @swagger
 * /todo:
 *   get:
 *     summary: Get Tasks
 *     description:  Get all tasks of user
 *     tags: [Tasks]
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
 *                  example: Get task success
 *                 data:
 *                  type: object
 *                  $ref: '#/components/schemas/tasks'
 *       "401":
 *         $ref: '#/components/responses/Unauthorized'
 *       "404":
 *         $ref: '#/components/responses/NotFound'
 *
 *   patch:
 *     summary: Update task
 *     description: Task update information.
 *     tags: [Tasks]
 *     security:
 *       - bearerAuth: []
 *     requestBody:
 *       required: true
 *       description: Update any field of task
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             properties:
 *              task_name:
 *                type: string,
 *                example: Coding
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
 *
 *   create task:
 *     summary: create task
 *     description: Create New Task.
 *     tags: [Tasks]
 *     security:
 *       - bearerAuth: []
 *     requestBody:
 *       required: true
 *       description: Create new task
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             properties:
 *              task_name:
 *                type: string,
 *                example: Coding
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
