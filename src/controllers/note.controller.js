const {
  addTask2DB,
  makeTaskCompleted,
  getTaskList,
  checkLimitTask,
  updateTask,
  deleteTask,
} = require("../services/note.service");


module.exports.create = async (req, res) => {
  try {
    const { content } = req.body;
    if(!content) throw new Error("Missing content")
    const userId = req.user.id;
    const limitTask = await checkLimitTask(userId);
    if(limitTask) throw new Error("Limit task can add in day");
    let data = await addTask2DB(content, userId);
    return res.status(200).json({ message: "Successful", data: data});
  } catch (err) {
    return res.status(400).json({ message: err.message});
  }
}

module.exports.makeTaskCompleted = async (req, res) => {
  try {
    const { id } = req.params;
    const userId = req.user.id;
    if (!id) throw new Error("Missing id");
    const data = await makeTaskCompleted(id, userId);
    if(!data) throw new Error("Id not found");
    return res.status(200).json({ message: "Successful", data: data });
  } catch (err) {
    return res.status(400).json({ message: err.message });
  }
}; 

module.exports.getList = async (req, res) => {
  try {
    const { day } = req.query;
    const userId = req.user.id
    const data = await getTaskList(userId, day);
    return res.status(200).json({ message: "Successful", data: data})
  } catch (err) {
    return res.status(400).json({ message: err.message });
  }
};

module.exports.updateTask = async (req, res) => {
  try {
    const { id } = req.params;
    const { content, ticked } = req.body;
    if(!content) throw new Error("Missing content")
    const userId= req.user.id;
    const taskUpdated = await updateTask(id, userId, content, ticked);
    if (!taskUpdated) throw new Error("Task not found")
    return res.status(200).json({ message: "Successful", data: taskUpdated });
  } catch (err) {
    return res.status(400).json({ message: err.message });
  }
}

module.exports.deleteTask = async (req, res) => {
  try {
    const { id } = req.params
    const userId = req.user.id;
    const taskDeleted = await deleteTask(id, userId);
    if (!taskDeleted) throw new Error("Task not found");
    return res.status(200).json({ message: "Successful"});
  } catch (err) {
    return res.status(400).json({ message: err.message });
  }
}

