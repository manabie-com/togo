import { NextFunction, Request, Response } from "express";
import dot from "dot-object";
import camelcaseKeys from "camelcase-keys";
import * as rfs from "rotating-file-stream";
import * as path from "path";
import morgan from "morgan";
import cors from "cors";
import _ from "lodash";

export type ExpressFunction = (req: Request, res: Response, next: NextFunction) => void;

export class BaseMiddleware {

    public static get cors(): ExpressFunction {
        return cors({
            allowedHeaders: [
                "Origin",
                "X-Requested-With",
                "Content-Type",
                "Accept-Language",
                "Content-Language",
                "X-Access-Token",
                "Authorization",
                "Authorization-Old",
                "X-Api-Key",
                "X-Skip-Interceptor",
                "X-Time-Zone",
                "X-Timezone-Offset",
                "Access-Control-Allow-Origin",
                "Access-Control-Allow-Methods",
                "Access-Control-Allow-Credentials",
            ],
            credentials: true,
            methods: "GET,POST,PUT,PATCH,DELETE",
            origin: "*",
            preflightContinue: false
        });
    }

    public static get accessLog(): ExpressFunction {
        const accessLogStream = rfs.createStream('access.log', {interval: '1d', path: path.resolve('logs/access')});
        return morgan('combined', {stream: accessLogStream});
    }

    public static get dotObject(): ExpressFunction {
        return function (req, res, next) {
            req.query = dot.object(req.query || {}) as any;
            next();
        }
    }

    public static get camelcaseKey(): ExpressFunction {
        return function (req, res, next) {
            const exclude = ['_id', /(.*)\.(.*)/];
            if (_.isObject(req.body)) {
                req.body = camelcaseKeys(req.body || {}, {deep: true, exclude});
            }
            req.params = camelcaseKeys(req.params || {}, {deep: true, exclude});
            req.query = camelcaseKeys(req.query || {}, {deep: true, exclude});
            next();
        }
    }
}
