const express = require('express');
const auth = require('../../middlewares/auth');
const validate = require('../../middlewares/validate');
const taskValidation = require('../../validations/task.validation');
const taskController = require('../../controllers/task.controller');

const router = express.Router();

router
  .route('/')
  .post(auth(), validate(taskValidation.createTask), taskController.createTask)
  .get(auth(), taskController.getTasks);

router
  .route('/:taskId')
  .patch(auth(), validate(taskValidation.updateTask), taskController.updateTask)
  .get(auth(), validate(taskValidation.deleteTask), taskController.getTask)
  .delete(auth(), validate(taskValidation.deleteTask), taskController.deleteTask);

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
 *                  $ref: '#/components/schemas/Tasks'
 *       "401":
 *         $ref: '#/components/responses/Unauthorized'
 *       "404":
 *         $ref: '#/components/responses/NotFound'
 *
 *   post:
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
 *              user_id:
 *                type: number,
 *                example: 1
 *              task_name:
 *                type: string,
 *                example: Coding
 *              task_priority:
 *                type: number,
 *                example: 1
 *     responses:
 *       "201":
 *         description: OK
 *         content:
 *           application/json:
 *             schema:
 *                $ref: '#/components/schemas/Tasks'
 *       "401":
 *         $ref: '#/components/responses/Unauthorized'
 *       "404":
 *         $ref: '#/components/responses/NotFound'
 */

/**
 * @swagger
 * /todo/:taskId:
 *   patch:
 *     summary: Update task
 *     description: Task update information.
 *     tags: [Tasks]
 *     security:
 *       - bearerAuth: []
 *     parameters:
 *       - in: query
 *         name: taskId
 *         schema:
 *            type: number
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
 *                $ref: '#/components/schemas/Tasks'
 *       "401":
 *         $ref: '#/components/responses/Unauthorized'
 *       "404":
 *         $ref: '#/components/responses/NotFound'
 *
 *   get:
 *     summary: get a task
 *     tags: [Tasks]
 *     security:
 *       - bearerAuth: []
 *     parameters:
 *       - in: query
 *         name:  taskId
 *         schema:
 *            type: number
 *     responses:
 *       "200":
 *         description: OK
 *         content:
 *           application/json:
 *             schema:
 *                $ref: '#/components/schemas/Tasks'
 *       "401":
 *         $ref: '#/components/responses/Unauthorized'
 *       "404":
 *         $ref: '#/components/responses/NotFound'
 *   delete:
 *     summary: delete a task
 *     tags: [Tasks]
 *     security:
 *       - bearerAuth: []
 *     parameters:
 *       - in: query
 *         name:  taskId
 *         schema:
 *            type: number
 *     responses:
 *       "200":
 *         description: OK
 *         content:
 *           application/json:
 *             schema:
 *                $ref: '#/components/schemas/Tasks'
 *       "401":
 *         $ref: '#/components/responses/Unauthorized'
 *       "404":
 *         $ref: '#/components/responses/NotFound'
 */
