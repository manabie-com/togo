import { Request } from "express";

import { IAPIResponse } from "@/interfaces/api.interface";
import UsersModel from "@/models/users.model";
import { CODE_201, CODE_400, CODE_404 } from "@/utils/responseStatus.util";

export const createNewUser = async (req: Request): Promise<IAPIResponse> => {
   const { body } = req;
   if (!body.username) return CODE_400();

   let user = await UsersModel.findOne({ username: body.username }).exec();
   if (user) return CODE_400({ message: "Username is invalid or already taken" });

   user = new UsersModel();
   user.username = body.username;
   if (body.limit) user.limit = body.limit;

   await user.save();

   return CODE_201();
};

export const getProfile = async (req: Request): Promise<IAPIResponse> => {
   const { username } = req.params;

   if (!username) return CODE_400();

   const user = await UsersModel.findOne({ username }, { _id: 0, limit: 1, username: 1 })
      .lean()
      .exec();

   if (!user) return CODE_404();

   return { data: user };
};
