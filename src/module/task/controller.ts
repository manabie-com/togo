import express, { Request,Response } from 'express'
import { BadRequestError } from '../../errors/bad-request-error'
import { taskService } from './service'

class TaskController {
    async createTask(req: Request, res: Response) {
        const payload = req.body.map((data: any) => {
            data.userId = req?.currentUser?.id
            return data
        })
        const result = await taskService.createTask(payload)
        return res.status(201).send(result)
    }

    async getTaskByCurrentUser(req: Request, res: Response) {
        const result = await taskService.getTaskByCurrentUser(req.query, req!.currentUser!.id)
        return res.send(result)
    }

}

export const taskController = new TaskController()