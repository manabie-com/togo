import { inject } from "inversify";
import { Request, Response } from 'express';
import { ApiResponse } from "../core";
import { REPOS, UserRepository } from "../repository";
import { BaseHttpController, controller, httpGet, interfaces, request, response } from "inversify-express-utils";
import { HOOKS } from "../hook";
import { SERVICES, UserService } from "../service";

@controller('/login')
export class AuthController extends BaseHttpController implements interfaces.Controller {

    @inject(REPOS.UserRepository) private userRepository!: UserRepository;
    @inject(SERVICES.UserService) private userService!: UserService;

    @httpGet('/', HOOKS.PasswordHook)
    async login(@request() req: Request, @response() res: Response): Promise<void> {
        try {
            const userId = String(req.query.user_id);
            const user: any = await this.userRepository.findOne({id: userId});
            if (!user) {
                return ApiResponse.create(res).unauthorized('USER_NOT_FOUND').build();
            }

            const token: string = await this.userService.createToken(user.id);
            return ApiResponse.create(res).ok('LOGIN_SUCCESS').data({data: token}).build();
        } catch (error) {
            return ApiResponse.create(res).error(error).build();
        }
    }
}
