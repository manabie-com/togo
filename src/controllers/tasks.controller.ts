import startOfDay from "date-fns/startOfDay";
import { Request } from "express";

import { IAPIResponse } from "@/interfaces/api.interface";
import TasksModel from "@/models/tasks.model";
import UsersModel from "@/models/users.model";
import { CODE_201, CODE_400, CODE_404 } from "@/utils/responseStatus.util";

export const createNewTask = async (req: Request): Promise<IAPIResponse> => {
   const { body, params } = req;
   if (!body.name || !params.userID) return CODE_400();

   const beginningOfDay = startOfDay(new Date());

   const getListTaskToday = TasksModel.find({ createdAt: { $gte: beginningOfDay } })
      .populate({ match: { id: params.userID }, path: "user", select: "_id" })
      .lean()
      .exec();
   const getInfoUser = UsersModel.findOne({ id: params.userID }, { _id: 0, limit: 1 })
      .lean()
      .exec();

   const [listTaskToday, infoUser] = await Promise.all([getListTaskToday, getInfoUser]);
   if (!infoUser) return CODE_404();

   if (listTaskToday && listTaskToday.length >= infoUser.limit)
      return CODE_400({
         message: "Bad request: This user is out of the limit in order to create a new task. ",
      });

   await TasksModel.create({ name: body.name, userId: params.userID });

   return CODE_201();
};

export const getListTask = async (req: Request): Promise<IAPIResponse> => {
   const { params } = req;
   if (!params.userID) return CODE_400();

   const listTask = await TasksModel.find({ userId: params.userID }, { _v: 0 }).lean().exec();

   return { data: listTask };
};
