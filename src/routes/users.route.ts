import { Router } from "express";

import { createNewUser, getProfile } from "@/controllers/users.controller";
import handleAPI from "@/middlewares/handleAPI.middleware";

const router = Router();

/**
 * POST /api/user/
 * @summary Create new user
 *
 * @typedef {object} createNewUser
 * @property {string} username - Name of user to login
 * @property {number} limit - Maximum daily limit
 *
 * @param {createNewUser} request.body.required - Username to create user
 * @return {Status201} - 201 - Success response - application/json
 * @return {Status400} - 400 - Bad request - application/json
 * @return {Status500} - 500 - Internal error response - application/json
 * @tags user
 */
router.post("/", handleAPI(createNewUser));

/**
 * GET /api/user/{username}
 *
 * @summary get user profile
 * @param {string} username.path.required - Get profile by username
 * @return {Status200} - 200 - Success response - application/json
 * @return {Status400} - 400 - Bad request - application/json
 * @return {Status500} - 500 - Internal error response - application/json
 * @tags user
 */
router.get("/:username?", handleAPI(getProfile));

export default router;
