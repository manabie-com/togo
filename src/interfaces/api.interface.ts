import { NextFunction, Request, Response } from "express";

import { IUsersModel } from "./users.interface";
/**
 * Interface of calling api(express)
 * req: `Request`, res: `Response`, next: `NextFunction`
 */
export interface IAPIFunction {
   // eslint-disable-next-line @typescript-eslint/no-explicit-any
   (req: Request, res: Response, next: NextFunction): Promise<any>;
}

export interface IAPIResponse {
   status?: number;
   data?: IUsersModel | unknown | null;
   total?: number;
}
