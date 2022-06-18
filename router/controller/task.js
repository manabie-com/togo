/**
 * @author Nguyen Minh Tam / ngmitamit@gmail.com
 */

const Tasks = require("../../DB/models/Task");
module.exports = {
  getTasks(req, res, next) {
    const taskList = Tasks.getTaskListByUserId(req?.userData?.id);

    return res.json({
      data: taskList,
    });
  },
};
