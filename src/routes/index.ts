import { Router } from "express";

import PingRoute from "@/routes/ping.route";
import UsersRoute from "@/routes/users.route";

const router = Router();
// #region Swagger typedef global
/**
 * A response of 200 status
 * @typedef {object} Status200
 * @property {integer} status - 200
 * @property {string} message - "OK: Success"
 * @property {object} data - object
 */
/**
 * A response of 201 status
 * @typedef {object} Status201
 * @property {integer} status - 201
 * @property {string} message - "OK: Success"
 * @property {object} data - object
 */
/**
 * A response of 400 status
 * @typedef {object} Status400
 * @property {integer} status - 400
 * @property {string} message - "Bad Request: Please check parameter",
 */
/**
 * A response of 404 status
 * @typedef {object} Status404
 * @property {integer} status - 404
 * @property {string} message - "Not found: There is no resource behind the URI",
 */
/**
 * A response of 500 status
 * @typedef {object} Status500
 * @property {integer} status - 500
 * @property {string} message - "Internal Server Error: API developers should avoid this error",
 * @property {string} error - "Cannot do inclusion on field url in exclusion projection"
 */
// #endregion
router.use("/ping/", PingRoute);
router.use("/user/", UsersRoute);

export default router;
