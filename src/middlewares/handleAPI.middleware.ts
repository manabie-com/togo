import { NextFunction, Request, Response } from "express";

import { IAPIFunction } from "@/interfaces/api.interface";
import { CODE_200, CODE_500 } from "@/utils/responseStatus.util";
/**
 * Caller function for global error handling
 * route all call this for try and handle errors
 */
export default (callback: IAPIFunction) =>
   (req: Request, res: Response, next: NextFunction): Promise<void | Response> =>
      Promise.resolve(callback(req, res, next))
         .then((data) => {
            const { status } = data;
            if (status) return res.status(status).send(data);

            return res.status(200).send(CODE_200({ ...data }));
         })
         .catch((error) => {
            console.error(`${error.message}`);
            res.status(500).send(CODE_500({ error: error.message }));
            next(error);
         });
