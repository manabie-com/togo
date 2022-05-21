import { model, Schema } from "mongoose";

import { ITasksModel } from "@/interfaces/tasks.interface";

const tasks = new Schema<ITasksModel>(
   {
      name: {
         require: true,
         type: String,
      },
      user: { ref: "users", type: Schema.Types.ObjectId },
   },
   { timestamps: true }
);

// Disable autoIndex to improve significant performance impact
tasks.set("autoIndex", false);

tasks.index({ username: 1 });

export default model<ITasksModel>("tasks", tasks, "tasks");
