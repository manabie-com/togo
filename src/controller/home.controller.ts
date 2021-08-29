import { all, BaseHttpController, controller, httpGet, interfaces, request, response } from "inversify-express-utils";
import { ApiResponse } from "../core";
import { Request, Response } from "express";
import { MessageConfig } from "../config";

@controller("/")
export class HomeController extends BaseHttpController implements interfaces.Controller {

    @httpGet("/")
    public async index(@request() req: Request, @response() res: Response) {
        try {
            const env: string = String(process.env.SERVER_ENV) || 'develop';
            let message: string = MessageConfig.WELCOME.replace('{{SERVER_ENV}}', env);
            return ApiResponse.create(res).ok(message).build();
        } catch (error) {
            return ApiResponse.create(res).error(error).build();
        }
    }

    @all('*')
    public all(@request() req: Request, @response() res: Response) {
        return ApiResponse.create(res).notFound(MessageConfig.URL_INVALID).build();
    }
}
