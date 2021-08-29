import { Request, Response } from 'express';
import { ApiResponse } from "../core";
import { BaseHttpController, controller, httpGet, httpPost, interfaces, request, response } from "inversify-express-utils";
import { HOOKS } from "../hook";
import { inject } from "inversify";
import { REPOS, TaskRepository } from "../repository";
import { Task } from "../model";
import { StringHelper } from "../shared";
import moment from "moment";

@controller('/tasks')
export class TaskController extends BaseHttpController implements interfaces.Controller {

    @inject(REPOS.TaskRepository) private taskRepository!: TaskRepository;

    @httpGet('/', HOOKS.AuthHook)
    async list(@request() req: Request, @response() res: Response): Promise<void> {
        try {
            const tasks: Task[] = await this.taskRepository.find(req.query);
            return ApiResponse.create(res).ok().data({data: tasks}).build();
        } catch (error) {
            return ApiResponse.create(res).error(error).build();
        }
    }

    @httpPost('/', HOOKS.AuthHook, HOOKS.TaskLimitHook)
    async create(@request() req: Request, @response() res: Response): Promise<void> {
        try {
            const body: any = {
                id: StringHelper.generateUUID(),
                content: req.body.content,
                user_id: res.locals.user_id,
                created_date: moment().format('YYYY-MM-DD')
            };
            const task: Task = await this.taskRepository.create(body);
            return ApiResponse.create(res).ok().data({data: task}).build();
        } catch (error) {
            return ApiResponse.create(res).error(error).build();
        }
    }
}
