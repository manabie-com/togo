import { Task } from '../../models/task'
import { User } from '../../models/user'
import { BadRequestError } from '../../errors/bad-request-error'
import { Request } from 'express'

interface bodyInterface {
    description: string,
    title: string,
    currentUserId: string
}

interface createTaskAttrs {
    success: boolean,
}

type Match = {
    complete?: boolean;
};

class TaskService {
    
    async createTask(body: bodyInterface) : Promise<createTaskAttrs> {
        const task = await Task.insertMany(body)
        return { success: true  }
    }

    async getTaskByCurrentUser(query: { 
        complete?: string, 
        sortBy?: string, 
        limit?: string, 
        skip?: string 
    } , id: string) {
        const match: Match = {}
        const sort: any = {}
    
        if(query.complete) {
            match.complete = query.complete === 'true'
        }
    
        if(query.sortBy) {
            const parts = (query.sortBy as string).split(':')
            sort[parts[0]] = parts[1] === 'desc' ? -1 : 1
        }
    
        const task = await Task.find({ userId: id })

        return task
    }

}

export const taskService = new TaskService()