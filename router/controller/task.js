/**
 * @author Nguyen Minh Tam / ngmitamit@gmail.com
 */

const Tasks = require("../../DB/models/Task");
const CustomError = require("../middlewares/customeError");

const { getTodayString } = require("../../utils/index");

module.exports = {
  getTasks(req, res, next) {
    const taskList = Tasks.getTaskListByUserId(req?.userData?.id);

    return res.json({
      data: taskList,
    });
  },

  createTask(req, res, next) {
    const content = req.body?.content;
    if (!content) return next(new CustomError(10000));

    const taskLimitPerDay = req.userData?.taskLimitPerDay || 0;
    const userId = req.userData?.id;

    const numberOfTaskOfUserToday = Tasks.getNumberOfTaskByUserIdAndDay(
      userId,
      getTodayString()
    );

    if (numberOfTaskOfUserToday >= taskLimitPerDay)
      return next(new CustomError(10001));

    const taskId = Tasks.insertTaskByUserId(userId, { content });

    return res.json({
      data: { taskId },
    });
  },
};
