import express, { Request,Response } from 'express'
import { currentUser } from '../../middlewares/current-user'
import { requireAuth } from '../../middlewares/require-auth'
import { taskController } from './controller'
import { requestValidate } from '../../middlewares/validate-request'
import { createTaskValidator } from '../../validator/task/createTaskValdator'

const routes = express.Router()

routes.post(
    '/api/tasks', 
    currentUser, 
    requireAuth, 
    createTaskValidator, 
    requestValidate,
    async (req: Request, res: Response)=> await taskController.createTask(req, res)
)

routes.get(
    '/api/tasks', 
    currentUser, 
    requireAuth, 
    async (req: Request, res: Response)=> await taskController.getTaskByCurrentUser(req, res)
)



export { routes as taskRoutes }