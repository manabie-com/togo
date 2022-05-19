const express = require("express");

const { taskValidation } = require("../../validations");
const { taskController } = require("../../controllers");

const auth = require("../../../middlewares/auth");
const validate = require("../../../middlewares/validate");

const router = express.Router();
router.use(auth);

router.post(
  "/",
  validate(taskValidation.taskSchema),
  taskController.createTask
);

router.get("/", taskController.getTasks);

router.get(
  "/:id",
  validate(taskValidation.objectIdSchema),
  taskController.getTaskById
);

router.patch(
  "/:id",
  validate(taskValidation.updateTaskSchema),
  taskController.updateTask
);

router.delete(
  "/:id",
  validate(taskValidation.objectIdSchema),
  taskController.deleteTask
);

module.exports = router;

/**
 * @swagger
 * tags:
 *   name: Tasks
 *   description: Task management
 */

/**
 * @swagger
 * /v1/tasks:
 *   post:
 *     summary: Create a new task
 *     tags: [Tasks]
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             required:
 *               - title
 *               - description
 *               - priority
 *               - completed
 *             properties:
 *               title:
 *                 type: string
 *               description:
 *                 type: string
 *               priority:
 *                 type: string
 *                 enum: [high, medium, low]
 *               completed:
 *                 type: boolean
 *             example:
 *               title: A difficult task
 *               description: This task is really hard
 *               priority: medium
 *               completed: false
 *     responses:
 *       "201":
 *         description: Created
 *         content:
 *           application/json:
 *             schema:
 *               $ref: '#/components/schemas/Task'
 *       "400":
 *         $ref: '#/components/responses/BadRequest'
 *       "401":
 *         $ref: '#/components/responses/Unauthorized'
 *     security:
 *       - bearerAuth: []
 *   get:
 *     summary: Get tasks
 *     tags: [Tasks]
 *     responses:
 *       "200":
 *         description: Get tasks successfully
 *         content:
 *           application/json:
 *             schema:
 *               type: array
 *               items:
 *                 $ref: '#/components/schemas/Task'
 *       "401":
 *         $ref: '#/components/responses/Unauthorized'
 *     security:
 *       - bearerAuth: []
 */

/**
 * @swagger
 * /v1/tasks/{id}:
 *   get:
 *     summary: Get a task by id
 *     tags: [Tasks]
 *     parameters:
 *       - in: path
 *         name: id
 *         schema:
 *           type: string
 *         required: true
 *         description: ID of the task to get
 *     responses:
 *       "200":
 *         description: Get task successfully
 *         content:
 *           application/json:
 *             schema:
 *               $ref: '#/components/schemas/Task'
 *       "400":
 *         $ref: '#/components/responses/BadRequest'
 *       "401":
 *         $ref: '#/components/responses/Unauthorized'
 *       "404":
 *         $ref: '#/components/responses/NotFound'
 *     security:
 *       - bearerAuth: []
 *   patch:
 *     summary: Update a task
 *     tags: [Tasks]
 *     parameters:
 *       - in: path
 *         name: id
 *         schema:
 *           type: string
 *         required: true
 *         description: ID of the task to update
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             properties:
 *               title:
 *                 type: string
 *               description:
 *                 type: string
 *               priority:
 *                 type: string
 *                 enum: [high, medium, low]
 *               completed:
 *                 type: boolean
 *             example:
 *               title: An difficult task updated
 *               description: This task is really hard and it updated
 *               priority: low
 *               completed: true
 *     responses:
 *       "200":
 *         description: Update task successfully
 *         content:
 *           application/json:
 *             schema:
 *               $ref: '#/components/schemas/Task'
 *       "400":
 *         $ref: '#/components/responses/BadRequest'
 *       "401":
 *         $ref: '#/components/responses/Unauthorized'
 *       "404":
 *         $ref: '#/components/responses/NotFound'
 *     security:
 *       - bearerAuth: []
 *   delete:
 *     summary: Delete a task by id
 *     tags: [Tasks]
 *     parameters:
 *       - in: path
 *         name: id
 *         schema:
 *           type: string
 *         required: true
 *         description: ID of the task to delete
 *     responses:
 *       "204":
 *         description: Delete task successfully
 *       "400":
 *         $ref: '#/components/responses/BadRequest'
 *       "401":
 *         $ref: '#/components/responses/Unauthorized'
 *       "404":
 *         $ref: '#/components/responses/NotFound'
 *     security:
 *       - bearerAuth: []
 */
