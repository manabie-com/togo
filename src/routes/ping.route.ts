import { Router } from "express";

import { pingResponse } from "@/controllers/ping.controller";
import handleAPI from "@/middlewares/handleAPI.middleware";

const router = Router();

// #region Swagger typedef
/**
 * A ping type
 * @typedef {object} Ping
 * @property {integer} status - 200
 * @property {string} message - "OK: Success"
 * @property {string} url - "/api/ping"
 * @property {string} version - "1.0.0"
 * @property {string} date - "19-05-2022"
 */
// #region

/**
 * GET /api/ping
 * @summary Return the version of the application
 * @return {Ping} 200 - Success response - application/json
 */
router.get("/", handleAPI(pingResponse));

export default router;
