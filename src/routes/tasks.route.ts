import { Router } from "express";

import { createNewTask, getListTask } from "@/controllers/tasks.controller";
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
 * @param {string} userID.path.required -  userID
 * @return {Status201} - 201 - Success response - application/json
 * @return {Status400} - 400 - Bad request - application/json
 * @return {Status500} - 500 - Internal error response - application/json
 * @tags task
 */
router.post("/:userID?", handleAPI(createNewTask));

/**
 * GET /api/task/{userID}
 * @summary Get list task of user
 *
 * @param {string} userID.path.required - userID
 * @return {Status200} - 200 - Success response - application/json
 * @return {Status400} - 400 - Bad request - application/json
 * @return {Status500} - 500 - Internal error response - application/json
 * @tags task
 */
router.get("/:userID?", handleAPI(getListTask));

export default router;
