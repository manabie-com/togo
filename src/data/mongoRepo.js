const TaskMongo = require("../model/taskMongo");
const UserMongo = require("../model/userMongo");
const httpError = require("../model/http-error");

module.exports = {
  // The addTask method of the MongoDB repo, which handles all DB transactions
  async addTask(task, currentDate) {
    let error;
    const newTask = new TaskMongo({
      task_name: task.task_name,
      user_name: task.user_name,
      date_created: currentDate,
    });

    try {
      // Retrieve the user object for the daily limit assigned to it
      const user = await UserMongo.findOne({
        user_name: task.user_name,
      });

      // If user is null, it does not exist
      if (!user) {
        error = new httpError("User does not exist", 404);
      } else {
        // Get the date by separating the timestamp
        const dateOnly = currentDate.split("T")[0];
        // Find all the tasks created on the current date
        const tasks = await TaskMongo.find({
          user_name: task.user_name,
          date_created: { $regex: ".*" + dateOnly + ".*" },
        });

        // If tasks object is null, there is an error
        if (!tasks) {
          throw new Error("Error getting tasks");
        }

        // If the tasks array size is the same as the user's daily limit, return an error
        if (tasks.length === user.task_daily_limit) {
          error = new httpError("Daily limit reached.", 403);
        } else {
          // Save the task if the current task count for the day is less than the daily limit
          await newTask.save();
        }
      }
    } catch (err) {
      console.log(err);
      error = new httpError(err.message, 500);
    }

    return error;
  },
};
