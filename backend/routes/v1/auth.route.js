const express = require('express');
const validate = require('../../middlewares/validate');
const authValidation = require('../../validations/auth.validation');
const authController = require('../../controllers/auth.controller');
const authorMinRole = require('../../middlewares/auth');

const router = express.Router();

router.post('/register', validate(authValidation.register), authController.register);
router.post('/login', validate(authValidation.login), authController.login);
router.post('/logout', validate(authValidation.logout), authController.logout);
router.post('/refresh-tokens', validate(authValidation.refreshTokens), authController.refreshTokens);
router.post('/forgot-password', validate(authValidation.forgotPassword), authController.forgotPassword);
router.post('/send-verification-email', authorMinRole(), authController.sendVerificationEmail);

module.exports = router;

/**
 * @swagger
 * tags:
 *   name: Auth
 *   description: Authentication
 */

/**
 * @swagger
 * /auth/register:
 *   post:
 *     summary: Register an user
 *     description: input date format mm/dd/yyyy
 *     tags: [Auth]
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             required:
 *               - accountId
 *               - email
 *               - password
 *               - dob
 *               - age
 *               - fullName
 *               - phoneNumber
 *               - fullDepartmentName
 *               - mainSkill
 *               - level
 *               - status
 *               - site
 *             properties:
 *               accountId:
 *                 type: string
 *                 description: id of account from FPT
 *               email:
 *                 type: string
 *                 format: email
 *                 description: email FPT, must be unique
 *               password:
 *                 type: string
 *               dob:
 *                 type: string
 *                 format: date
 *               age:
 *                 type: number
 *               fullName:
 *                 type: string
 *               phoneNumber:
 *                 type: string
 *                 description: should be "+84xxxxxxxxx" or "0xxxxxxxxx"
 *               fullDepartmentName:
 *                 type: string
 *               mainSkill:
 *                 type: string
 *               level:
 *                 type: number
 *               status:
 *                 type: string
 *               type:
 *                 type: string
 *               site:
 *                 type: string
 *               onboardDay:
 *                 type: string
 *                 format: date
 *               mentor:
 *                 type: string
 *               mentorFullName:
 *                 typeL string
 *             example:
 *                accountId: "quangnh69"
 *                email: quangnh69@gmail.com
 *                password: pass
 *                dob: 03-04-1999
 *                age: 23
 *                fullName: Nguyễn Hồng Quang
 *                phoneNumber: "+84855513506"
 *                fullDepartmentName: FWA.EC
 *                mainSkill: Java
 *                level: 10
 *                status: Onboard
 *                type: Internal
 *                site: hà nội
 *                onboardDay: 03-01-2022
 *                mentor: TuanDV32
 *                mentorFullName: Đoàn Văn Tuấn
 *     responses:
 *       "201":
 *         description: Register success
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 status:
 *                  type: integer
 *                  example: 201
 *                 message:
 *                  type: string
 *                  example: Created user success
 *                 data:
 *                  type: object
 *                  $ref: '#/components/schemas/user'
 *       "400":
 *          $ref: '#/components/responses/DuplicateEmail'
 *       400:
 *          $ref: '#/components/responses/DuplicateNumberPhone'
 */

/**
 * @swagger
 * /auth/login:
 *   post:
 *     summary: Login
 *     tags: [Auth]
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             required:
 *               - email
 *               - password
 *             properties:
 *               email:
 *                 type: string
 *                 format: email
 *               password:
 *                 type: string
 *             example:
 *               email: fake@gmail.com
 *               password: pass
 *     responses:
 *       "200":
 *         description: Login success
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
 *                  example: Login success
 *                 data:
 *                  type: object
 *                  properties:
 *                    user:
 *                      $ref: '#/components/schemas/user'
 *                    tokens:
 *                      $ref: '#/components/schemas/AuthTokens'
 *       "400":
 *          $ref: '#/components/responses/EmailNotFound'
 *       "401":
 *          $ref: '#/components/responses/IncorrectEmailOrPassword'
 */

/**
 * @swagger
 * /auth/logout:
 *   post:
 *     summary: Logout
 *     tags: [Auth]
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             required:
 *               - refreshToken
 *             properties:
 *               refreshToken:
 *                 type: string
 *             example:
 *               refreshToken: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI1ZWJhYzUzNDk1NGI1NDEzOTgwNmMxMTIiLCJpYXQiOjE1ODkyOTg0ODQsImV4cCI6MTU4OTMwMDI4NH0.m1U63blB0MLej_WfB7yC2FTMnCziif9X8yzwDEfJXAg
 *     responses:
 *       "200":
 *         description: Logout success
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
 *                  example: Logout success
 *                 data:
 *                  type: string
 *                  example: null
 *       "404":
 *         $ref: '#/components/responses/NotFound'
 */

/**
 * @swagger
 * /auth/refresh-tokens:
 *   post:
 *     summary: Refresh auth tokens
 *     tags: [Auth]
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             required:
 *               - refreshToken
 *             properties:
 *               refreshToken:
 *                 type: string
 *             example:
 *               refreshToken: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI1ZWJhYzUzNDk1NGI1NDEzOTgwNmMxMTIiLCJpYXQiOjE1ODkyOTg0ODQsImV4cCI6MTU4OTMwMDI4NH0.m1U63blB0MLej_WfB7yC2FTMnCziif9X8yzwDEfJXAg
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
 *                  example: Refresh Token Success
 *                 data:
 *                  type: object
 *                  properties:
 *                    access:
 *                      $ref: '#/components/schemas/Token'
 *                    refresh:
 *                      $ref: '#/components/schemas/Token'
 *       "404":
 *         $ref: '#/components/responses/NotFound'
 *       "401":
 *         $ref: '#/components/responses/Unauthorized'
 */
