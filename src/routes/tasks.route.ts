import { Router } from "express";

import { createNewTask } from "@/controllers/tasks.controller";
import handleAPI from "@/middlewares/handleAPI.middleware";

const router = Router();

/**
 * POST /api/task/{userID}
 * @summary Create new task
 *
 * @typedef {object} createNewTask
 * @property {string} name - Name of task
 *
 * @param {createNewTask} request.body.required - Payload to create task
 * @param {string} userID.path.required - Create new task by userID
 * @return {Status201} - 201 - Success response - application/json
 * @return {Status400} - 400 - Bad request - application/json
 * @return {Status500} - 500 - Internal error response - application/json
 * @tags task
 */
router.post("/:userID?", handleAPI(createNewTask));

export default router;
