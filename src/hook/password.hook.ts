import { NextFunction, Request, Response } from "express";
import { inject, injectable } from "inversify";
import { BaseMiddleware, next, request, response } from "inversify-express-utils";
import { REPOS, UserRepository } from "../repository";
import { ApiResponse } from "../core";
import { MessageConfig } from "../config";
import { PasswordHelper } from "../shared/helper/password.helper";

@injectable()
export class PasswordHook extends BaseMiddleware {

    @inject(REPOS.UserRepository) private userRepository!: UserRepository;

    async handler(@request() req: Request, @response() res: Response, @next() next: NextFunction): Promise<void> {
        try {
            const userId = String(req.query.user_id) || null;
            const password = String(req.query.password) || null;

            if (!userId) {
                return ApiResponse.create(res).notOk('USERID_REQUIRED').build();
            }

            if (!password) {
                return ApiResponse.create(res).notOk(MessageConfig.PASSWORD_REQUIRED).build();
            }

            const user: any = await this.userRepository.findOne({id: userId});
            if (!user) {
                return ApiResponse.create(res).notFound('USER_NOT_FOUND').build();
            }

            if (user.password === password) {
                return next();
            }

            if (!PasswordHelper.verify(password, user.password)) {
                return ApiResponse.create(res).notOk('PASSWORD_INVALID').build();
            }

            next();
        } catch (error) {
            return ApiResponse.create(res).error(error).build();
        }
    }
}
