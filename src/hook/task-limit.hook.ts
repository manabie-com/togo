import { NextFunction, Request, Response } from "express";
import { inject, injectable } from "inversify";
import { BaseMiddleware, next, request, response } from "inversify-express-utils";
import { REPOS, TaskRepository, UserRepository } from "../repository";
import { ApiResponse } from "../core";
import moment from "moment";

@injectable()
export class TaskLimitHook extends BaseMiddleware {

    @inject(REPOS.UserRepository) private userRepository!: UserRepository;
    @inject(REPOS.TaskRepository) private taskRepository!: TaskRepository;

    async handler(@request() req: Request, @response() res: Response, @next() next: NextFunction): Promise<void> {
        try {
            const userId = String(res.locals.user_id);
            const user: any = await this.userRepository.findOne({id: userId});
            if (!user) {
                return ApiResponse.create(res).notFound('USER_NOT_FOUND').build();
            }

            const max_todo: number = Number(user.max_todo) || 5;
            const today: string = moment().format('YYYY-MM-DD');
            const tasks: any[] = await this.taskRepository.find({user_id: userId, created_date: today});
            if (tasks.length >= max_todo) {
                return ApiResponse.create(res).notOk('LIMIT_TODO_REACHED').build();
            }

            next();
        } catch (error) {
            return ApiResponse.create(res).error(error).build();
        }
    }
}
