import express, { Request,Response } from 'express'
import { requestValidate } from '../../middlewares/validate-request'
import { signinValidator } from '../../validator/auth/signinValidator'
import { currentUser } from '../../middlewares/current-user'
import { signupdaValidator } from '../../validator/auth/signupValidator'
import { authController } from './controller'

const routes = express.Router()

routes.get('/api/users/currentuser', currentUser,  async (req: Request, res: Response)=> await authController.getCurrentUser(req, res))

routes.post('/api/users/signin',signinValidator, requestValidate, async (req: Request, res: Response)=> await authController.signIn(req, res))

routes.post('/api/users/signup',signupdaValidator, requestValidate, async (req: Request, res: Response)=> await authController.signUp(req, res))

routes.post('/api/users/signout', async (req: Request, res: Response)=> await authController.logout(req, res))

export { routes as authRoutes }