const express = require('express');
const httpStatus = require('http-status');
const response = require('../../utils/responseTemp');

const router = express.Router();

router.get('/', (req, res) => {
  res.send(response(httpStatus.OK, 'Todos-App', null));
});

module.exports = router;

/**
 * @swagger
 * tags:
 *   name: Home
 *   description: Home api
 */

/**
 * @swagger
 * /:
 *   get:
 *     summary: Home test
 *     description:  For test api success
 *     tags: [Home]
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
 *                  example: Todos-App
 *                 data: null
 */
