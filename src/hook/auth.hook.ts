import { NextFunction, Request, Response } from "express";
import { inject, injectable } from "inversify";
import { REPOS, UserRepository } from "../repository";
import { ApiResponse } from "../core";
import { BaseMiddleware, next, request, response } from "inversify-express-utils";
import jwt from "jsonwebtoken";

@injectable()
export class AuthHook extends BaseMiddleware {

    @inject(REPOS.UserRepository) private userRepository!: UserRepository;

    async handler(@request() req: Request, @response() res: Response, @next() next: NextFunction): Promise<void> {
        try {
            const token: string = req.headers['authorization'] || '';
            if (!token) {
                return ApiResponse.create(res).unauthorized().build();
            }

            const secretKey: string = String(process.env.JWT_AUTH_SECRET_KEY);
            const payload: any = await jwt.verify(token, secretKey);
            const loggedUser: any = await this.userRepository.findOne({id: payload.user_id});
            if (!loggedUser) {
                return ApiResponse.create(res).unauthorized('ACCOUNT_NOT_FOUND').build();
            }

            res.locals.loggedUser = loggedUser;
            res.locals.user_id = loggedUser.id;
            next();
        } catch (error) {
            return ApiResponse.create(res).unauthorized().error(error).build();
        }
    }
}
