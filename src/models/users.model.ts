import { model, Schema } from "mongoose";

import { IUsersModel } from "@/interfaces/users.interface";

const users = new Schema<IUsersModel>(
   {
      limit: { default: 10, type: Number },
      username: {
         require: true,
         type: String,
         unique: true,
      },
   },
   { timestamps: true }
);

// Disable autoIndex to improve significant performance impact
users.set("autoIndex", false);

users.index({ username: 1 });

export default model<IUsersModel>("users", users, "users");
